package tsf

import "unsafe"

// ITfIntegratableCandidateListUIElement COM interface.
type ITfIntegratableCandidateListUIElement struct {
	IUnknown
}

// ITfIntegratableCandidateListUIElementVtbl COM interface vtable.
type ITfIntegratableCandidateListUIElementVtbl struct {
	IUnknownVtbl
}

// ITfIntegratableCandidateListUIElement returns the ITfIntegratableCandidateListUIElement vtable.
func (obj *ITfIntegratableCandidateListUIElement) ITfIntegratableCandidateListUIElement() *ITfIntegratableCandidateListUIElementVtbl {
	return (*ITfIntegratableCandidateListUIElementVtbl)(unsafe.Pointer(obj.Vtbl))
}
