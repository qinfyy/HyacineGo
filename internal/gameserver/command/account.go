package command

import (
	"fmt"
	"strconv"
	"strings"
)

func RegisterAccount(r *Registry) {
	r.Register("account", handleAccount)
}

func handleAccount(ctx Context, args []string) error {
	if ctx.DataBase == nil {
		return fmt.Errorf("database not configured")
	}
	if len(args) == 0 {
		return fmt.Errorf("usage: /account create <email> [reservedPlayerUid]")
	}

	switch strings.ToLower(args[0]) {
	case "create":
		return accountCreate(ctx, args[1:])
	default:
		return fmt.Errorf("usage: /account create <email> [reservedPlayerUid]")
	}
}

func accountCreate(ctx Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: /account create <email> [reservedPlayerUid]")
	}
	email := strings.TrimSpace(args[0])
	if email == "" {
		return fmt.Errorf("empty email")
	}

	var reserved int
	if len(args) >= 2 {
		n, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid reservedPlayerUid: %v", err)
		}
		reserved = n
	}

	acc, created, err := ctx.DataBase.GetOrCreateAccount(email)
	if err != nil {
		return err
	}
	if reserved != 0 {
		acc.ReservedPlayerUID = reserved
		acc.Touch()
	}
	if acc.DispatchToken == "" {
		acc.GenerateDispatchToken()
	}
	if acc.ComboToken == "" {
		acc.GenerateComboToken()
	}
	if err := ctx.DataBase.SaveAccounts(); err != nil {
		return err
	}

	if created {
		fmt.Fprintf(ctx.Out, "created: uid=%s email=%s reservedPlayerUid=%d\n", acc.UID, acc.Username, acc.ReservedPlayerUID)
	} else {
		fmt.Fprintf(ctx.Out, "updated: uid=%s email=%s reservedPlayerUid=%d\n", acc.UID, acc.Username, acc.ReservedPlayerUID)
	}
	fmt.Fprintf(ctx.Out, "dispatch_token=%s\n", acc.DispatchToken)
	fmt.Fprintf(ctx.Out, "combo_token=%s\n", acc.ComboToken)
	return nil
}
