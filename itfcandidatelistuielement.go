package tsf

import "unsafe"

// ITfCandidateListUIElement COM interface.
type ITfCandidateListUIElement struct {
	IUnknown
}

// ITfCandidateListUIElementVtbl COM interface vtable.
type ITfCandidateListUIElementVtbl struct {
	IUnknownVtbl
}

// ITfCandidateListUIElement returns the ITfCandidateListUIElement vtable.
func (obj *ITfCandidateListUIElement) ITfCandidateListUIElement() *ITfCandidateListUIElementVtbl {
	return (*ITfCandidateListUIElementVtbl)(unsafe.Pointer(obj.Vtbl))
}
