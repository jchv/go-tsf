package tsf

import "unsafe"

// ITfCandidateList COM interface.
type ITfCandidateList struct {
	IUnknown
}

// ITfCandidateListVtbl COM interface vtable.
type ITfCandidateListVtbl struct {
	IUnknownVtbl
}

// ITfCandidateList returns the ITfCandidateList vtable.
func (obj *ITfCandidateList) ITfCandidateList() *ITfCandidateListVtbl {
	return (*ITfCandidateListVtbl)(unsafe.Pointer(obj.Vtbl))
}
