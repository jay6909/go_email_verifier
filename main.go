package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"net"

	"github.com/jay6909/go_email_verifier/emailverifier"
)

//export Verify
func Verify(cEmail *C.char) C.int {
	// 1. Convert C string to Go string
	goEmail := C.GoString(cEmail)

	res, err := emailverifier.Verify(goEmail)

	// 2. Call your existing package logic
	if err != nil {
		// Detect if it's a network/timeout error
		if netErr, ok := err.(net.Error); ok && (netErr.Timeout() || netErr.Temporary()) {
			return -1 // Network/Offline status
		}
		return 0 // Other errors treated as invalid
	}

	if res.IsValid {
		return 1
	}

	return 0
}

// Required for c-shared build mode
func main() {}
