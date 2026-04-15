package test_go_pointers

import (
	"testing"

	"github.com/jay6909/go-pointers/email"
)

func TestEmailVerifier(t *testing.T) {
	goodEmail := "jayadityagaikwad@gmail.com"
	badEmail := "jayadityagaikwad@bossbadmail.com"

	result, err := email.Verify(goodEmail)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// log result
	t.Logf("Result for %s: %+v", goodEmail, result)
	if !result.IsValid {
		t.Errorf("Expected valid email, got invalid")
	}
	if !result.HasMXRecords {
		t.Errorf("Expected MX records, got none")
	}

	result, err = email.Verify(badEmail)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// log result
	t.Logf("Result for %s: %+v", goodEmail, result)
	if result.IsValid {
		t.Errorf("Expected invalid email, got valid")
	}

}
