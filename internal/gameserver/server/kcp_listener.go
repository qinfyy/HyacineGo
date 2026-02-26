package server

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	kcp "github.com/xtaci/kcp-go/v5"
)

// Star Rail uses a KCP variant where each segment header includes an extra uint32 "token"
// after conv: <conv:uint32><token:uint32><cmd:uint8><frg:uint8><wnd:uint16><ts:uint32><sn:uint32><una:uint32><len:uint32>
// The upstream kcp-go implementation follows the standard header (<conv> then cmd...). We adapt by
// stripping/inserting the token when feeding segments in/out of the KCP state machine.

const (
	srKcpHeaderOverhead = 32 // conv+token+cmd+frg+wnd+ts+sn+una+len
)

type srKcpKey struct {
	Conv  uint32
	Token uint32
	Addr  string
}

type srKcpConn struct {
	udp    *net.UDPConn
	remote *net.UDPAddr
	token  uint32

	kcp *kcp.KCP

	mu           sync.Mutex
	closed       bool
	readDeadline time.Time
	writeDead    time.Time

	// stash holds leftover bytes when the caller's Read buffer is smaller than a received message.
	stash []byte

	maxRecv int

	onClose func()
}

func newSrKcpConn(udp *net.UDPConn, remote *net.UDPAddr, conv, token uint32, maxRecv int, onClose func()) *srKcpConn {
	c := &srKcpConn{
		udp:     udp,
		remote:  remote,
		token:   token,
		maxRecv: maxRecv,
		onClose: onClose,
	}

	output := func(buf []byte, size int) {
		if size <= 0 {
			return
		}
		seg := buf[:size]
		if len(seg) < 4 {
			return
		}

		// Insert token right after conv.
		out := make([]byte, len(seg)+4)
		copy(out[0:4], seg[0:4])
		binary.LittleEndian.PutUint32(out[4:8], token)
		copy(out[8:], seg[4:])
		_, _ = udp.WriteToUDP(out, remote)
	}

	k := kcp.NewKCP(conv, output)
	// Match typical SR settings (as seen in python/LC implementations).
	_ = k.SetMtu(1400)
	_ = k.WndSize(32, 128)
	_ = k.NoDelay(1, 100, 2, 1)
	c.kcp = k
	return c
}

func (c *srKcpConn) feedSegmentWithToken(pkt []byte) {
	// pkt is a single UDP datagram containing 1..n KCP segments with SR header format.
	// Convert by stripping the token field in each segment header.
	if len(pkt) < srKcpHeaderOverhead {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return
	}

	converted := make([]byte, 0, len(pkt))
	i := 0
	for {
		if len(pkt[i:]) < srKcpHeaderOverhead {
			break
		}
		// Validate token (offset 4..8).
		if binary.LittleEndian.Uint32(pkt[i+4:i+8]) != c.token {
			return
		}
		segLen := int(binary.LittleEndian.Uint32(pkt[i+28 : i+32]))
		total := srKcpHeaderOverhead + segLen
		if total <= 0 || len(pkt[i:]) < total {
			return
		}

		// Standard KCP expects the header without token.
		converted = append(converted, pkt[i:i+4]...)
		converted = append(converted, pkt[i+8:i+total]...)
		i += total
		if i >= len(pkt) {
			break
		}
	}

	if len(converted) == 0 {
		return
	}
	_ = c.kcp.Input(converted, kcp.IKCP_PACKET_REGULAR, true)
}

func (c *srKcpConn) tick() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return
	}
	c.kcp.Update()
}

func (c *srKcpConn) Read(p []byte) (int, error) {
	for {
		c.mu.Lock()
		if c.closed {
			c.mu.Unlock()
			return 0, io.EOF
		}

		if len(c.stash) > 0 {
			n := copy(p, c.stash)
			c.stash = c.stash[n:]
			c.mu.Unlock()
			return n, nil
		}

		if !c.readDeadline.IsZero() && time.Now().After(c.readDeadline) {
			c.mu.Unlock()
			return 0, timeoutError{}
		}

		peek := c.kcp.PeekSize()
		if peek > 0 {
			if peek > c.maxRecv {
				// Drop oversized message to avoid unbounded alloc.
				buf := make([]byte, peek)
				_ = c.kcp.Recv(buf)
				c.mu.Unlock()
				return 0, fmt.Errorf("kcp message too large: %d", peek)
			}

			buf := make([]byte, peek)
			n := c.kcp.Recv(buf)
			if n < 0 {
				c.mu.Unlock()
			} else {
				buf = buf[:n]
				// Serve partial read if caller buffer is smaller.
				if len(buf) > len(p) {
					nn := copy(p, buf)
					c.stash = append(c.stash, buf[nn:]...)
					c.mu.Unlock()
					return nn, nil
				}
				copy(p, buf)
				c.mu.Unlock()
				return len(buf), nil
			}
		} else {
			c.mu.Unlock()
		}

		// No data available yet; wait a bit. kcp.Update() is called externally by a ticker.
		time.Sleep(5 * time.Millisecond)
	}
}

func (c *srKcpConn) Write(p []byte) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return 0, net.ErrClosed
	}
	if !c.writeDead.IsZero() && time.Now().After(c.writeDead) {
		return 0, timeoutError{}
	}
	if ret := c.kcp.Send(p); ret < 0 {
		return 0, fmt.Errorf("kcp send failed: %d", ret)
	}
	// Flush immediately to reduce client-side timeouts.
	c.kcp.Update()
	return len(p), nil
}

func (c *srKcpConn) Close() error {
	c.mu.Lock()
	already := c.closed
	c.closed = true
	c.mu.Unlock()
	if already {
		return nil
	}
	if c.onClose != nil {
		c.onClose()
	}
	return nil
}

func (c *srKcpConn) LocalAddr() net.Addr  { return c.udp.LocalAddr() }
func (c *srKcpConn) RemoteAddr() net.Addr { return c.remote }
func (c *srKcpConn) SetDeadline(t time.Time) error {
	_ = c.SetReadDeadline(t)
	_ = c.SetWriteDeadline(t)
	return nil
}
func (c *srKcpConn) SetReadDeadline(t time.Time) error {
	c.mu.Lock()
	c.readDeadline = t
	c.mu.Unlock()
	return nil
}
func (c *srKcpConn) SetWriteDeadline(t time.Time) error {
	c.mu.Lock()
	c.writeDead = t
	c.mu.Unlock()
	return nil
}

type timeoutError struct{}

func (timeoutError) Error() string   { return "i/o timeout" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }

func (s *Server) runKCP(ctx context.Context) error {
	udpAddr, err := net.ResolveUDPAddr("udp", s.addr)
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	s.udp = udpConn

	_ = udpConn.SetReadBuffer(8 * 1024 * 1024)
	_ = udpConn.SetWriteBuffer(8 * 1024 * 1024)

	s.mu.RLock()
	maxFrame := s.opts.MaxFrameBytes
	s.mu.RUnlock()
	if maxFrame <= 0 {
		maxFrame = 1024 * 1024
	}

	sessions := make(map[srKcpKey]*srKcpConn)
	var sessionsMu sync.Mutex

	removeSession := func(k srKcpKey) func() {
		return func() {
			sessionsMu.Lock()
			delete(sessions, k)
			sessionsMu.Unlock()
		}
	}

	recvBuf := make([]byte, 64*1024)

	readErrCh := make(chan error, 1)
	go func() {
		defer close(readErrCh)
		for {
			_ = udpConn.SetReadDeadline(time.Now().Add(1 * time.Second))
			n, addr, err := udpConn.ReadFromUDP(recvBuf)
			if err != nil {
				var ne net.Error
				if errors.As(err, &ne) && ne.Timeout() {
					select {
					case <-ctx.Done():
						return
					default:
						continue
					}
				}
				readErrCh <- err
				return
			}
			if n < srKcpHeaderOverhead {
				continue
			}

			pkt := append([]byte(nil), recvBuf[:n]...)
			conv := binary.LittleEndian.Uint32(pkt[0:4])
			token := binary.LittleEndian.Uint32(pkt[4:8])
			key := srKcpKey{Conv: conv, Token: token, Addr: addr.String()}

			sessionsMu.Lock()
			c := sessions[key]
			if c == nil {
				c = newSrKcpConn(udpConn, addr, conv, token, maxFrame+64, removeSession(key))
				sessions[key] = c
				go s.handleConn(ctx, c)
			}
			sessionsMu.Unlock()

			c.feedSegmentWithToken(pkt)
		}
	}()

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			_ = udpConn.Close()
			return nil
		case err := <-readErrCh:
			if err == nil {
				return nil
			}
			return err
		case <-ticker.C:
			sessionsMu.Lock()
			for _, c := range sessions {
				c.tick()
			}
			sessionsMu.Unlock()
		}
	}
}
