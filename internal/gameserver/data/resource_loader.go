package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ResourceLoader is the Go counterpart of emu.lunarcore.data.ResourceLoader.
// It loads JSON resources from the local resources directory.
type ResourceLoader struct {
	paths Paths
	root  string // absolute

	mu    sync.RWMutex
	cache map[string][]byte // absolute file path -> bytes
}

func NewResourceLoader(resourceRoot string) *ResourceLoader {
	root := strings.TrimSpace(resourceRoot)
	if root == "" {
		root = "."
	}
	if abs, err := filepath.Abs(root); err == nil {
		root = abs
	}
	return &ResourceLoader{
		paths: Paths{Root: root},
		root:  root,
		cache: map[string][]byte{},
	}
}

func (l *ResourceLoader) Paths() Paths { return l.paths }

func (l *ResourceLoader) ReadFile(relOrAbsPath string) ([]byte, error) {
	p := strings.TrimSpace(relOrAbsPath)
	if p == "" {
		return nil, fmt.Errorf("empty path")
	}
	if !filepath.IsAbs(p) {
		p = filepath.Join(l.root, filepath.FromSlash(p))
	}
	p = filepath.Clean(p)

	l.mu.RLock()
	if b, ok := l.cache[p]; ok {
		l.mu.RUnlock()
		return b, nil
	}
	l.mu.RUnlock()

	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	l.mu.Lock()
	l.cache[p] = b
	l.mu.Unlock()

	return b, nil
}

func (l *ResourceLoader) UnmarshalJSON(relOrAbsPath string, out any) error {
	b, err := l.ReadFile(relOrAbsPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func (l *ResourceLoader) LoadJSON(relOrAbsPath string, out any) error {
	if err := l.UnmarshalJSON(relOrAbsPath, out); err != nil {
		return fmt.Errorf("load json %q: %w", relOrAbsPath, err)
	}
	return nil
}

func (l *ResourceLoader) ExcelJSONPath(name string) string {
	name = strings.TrimSpace(name)
	if !strings.HasSuffix(strings.ToLower(name), ".json") {
		name += ".json"
	}
	return filepath.Join("ExcelOutput", name)
}

func (l *ResourceLoader) LoadExcelJSON(name string, out any) error {
	path := l.ExcelJSONPath(name)
	if err := l.UnmarshalJSON(path, out); err != nil {
		return fmt.Errorf("load excel %q: %w", name, err)
	}
	return nil
}

// LoadConfigJSON loads a JSON file under resources/Config.
// Examples:
//
//	LoadConfigJSON("LevelOutput/RuntimeFloor/P10000_F10000003.json", &out)
//	LoadConfigJSON("Config/LevelOutput/RuntimeFloor/P10000_F10000003.json", &out)
func (l *ResourceLoader) LoadConfigJSON(pathUnderConfig string, out any) error {
	p := filepath.ToSlash(strings.TrimSpace(pathUnderConfig))
	p = strings.TrimPrefix(p, "./")
	p = strings.TrimPrefix(p, "Config/")
	full := filepath.Join("Config", filepath.FromSlash(p))
	if err := l.UnmarshalJSON(full, out); err != nil {
		return fmt.Errorf("load config %q: %w", pathUnderConfig, err)
	}
	return nil
}

func (l *ResourceLoader) ClearCache() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.cache = map[string][]byte{}
}
