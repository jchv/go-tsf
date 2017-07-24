package tsf

import "fmt"

// COMError represents an error in a COM call.
type COMError struct{ hresult uint32 }

func (e COMError) Error() string {
	return fmt.Sprintf("error 0x%08x", e.hresult)
}
