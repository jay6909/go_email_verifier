package main

/*
#include <stdlib.h>
*/
import "C"
import "github.com/jay6909/go_email_verifier/emailverifier"

//export Verify
func Verify(cEmail *C.char) C.int {
	// 1. Convert C string to Go string
	goEmail := C.GoString(cEmail)

	// 2. Call your existing package logic
	res, err := emailverifier.Verify(goEmail)

	// 3. Return a simple C-compatible type (1 for true, 0 for false)
	if err == nil && res.IsValid {
		return 1
	}
	return 0
}

// Required for c-shared build mode
func main() {}
