package tsf

import "unsafe"

// ITfCandidateString COM interface.
type ITfCandidateString struct {
	IUnknown
}

// ITfCandidateStringVtbl COM interface vtable.
type ITfCandidateStringVtbl struct {
	IUnknownVtbl
}

// ITfCandidateString returns the ITfCandidateString vtable.
func (obj *ITfCandidateString) ITfCandidateString() *ITfCandidateStringVtbl {
	return (*ITfCandidateStringVtbl)(unsafe.Pointer(obj.Vtbl))
}
