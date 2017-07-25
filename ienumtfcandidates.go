package tsf

import "unsafe"

// IEnumTfCandidates represents the IEnumTfCandidates COM interface.
type IEnumTfCandidates struct {
	IUnknown
}

// IEnumTfCandidatesVtbl COM interface vtable.
type IEnumTfCandidatesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfCandidates returns the IEnumTfCandidates vtable.
func (obj *IEnumTfCandidates) IEnumTfCandidates() *IEnumTfCandidatesVtbl {
	return (*IEnumTfCandidatesVtbl)(unsafe.Pointer(obj.Vtbl))
}
