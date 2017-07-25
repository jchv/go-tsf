package tsf

import "unsafe"

// ITfCandidateListUIElementBehavior COM interface.
type ITfCandidateListUIElementBehavior struct {
	IUnknown
}

// ITfCandidateListUIElementBehaviorVtbl COM interface vtable.
type ITfCandidateListUIElementBehaviorVtbl struct {
	IUnknownVtbl
}

// ITfCandidateListUIElementBehavior returns the ITfCandidateListUIElementBehavior vtable.
func (obj *ITfCandidateListUIElementBehavior) ITfCandidateListUIElementBehavior() *ITfCandidateListUIElementBehaviorVtbl {
	return (*ITfCandidateListUIElementBehaviorVtbl)(unsafe.Pointer(obj.Vtbl))
}
