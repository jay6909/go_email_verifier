package emailverifier

import (
	"context"
	"net"
	"strings"
	"time"
)

// Result holds the details of the verification
type Result struct {
	Email        string
	IsValid      bool
	HasMXRecords bool
	IsDisposable bool
}

// Check syntax and MX records
func Verify(email string) (*Result, error) {
	res := &Result{Email: email}

	// 1. Basic Syntax & Split Check
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		res.IsValid = false
		return res, nil // Fail fast if syntax is obviously wrong
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	domain := parts[1]
	mx, err := net.DefaultResolver.LookupMX(ctx, domain)

	if err != nil {
		res.HasMXRecords = false
		res.IsValid = false
		return res, err
	}

	if len(mx) > 0 {
		res.HasMXRecords = true
		res.IsValid = true
	}
	return res, nil
}
