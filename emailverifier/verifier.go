package emailverifier

import (
	"net"
	"strings"
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

	// 1. Basic Syntax
	if !strings.Contains(email, "@") {
		return res, nil
	}

	// 2. Lookup MX Records
	parts := strings.Split(email, "@")
	domain := parts[1]
	mx, err := net.LookupMX(domain)

	if err == nil && len(mx) > 0 {
		res.HasMXRecords = true
		res.IsValid = true
	}

	return res, nil
}
