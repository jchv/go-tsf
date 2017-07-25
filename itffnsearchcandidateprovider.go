package tsf

import "unsafe"

// ITfFnSearchCandidateProvider COM interface.
type ITfFnSearchCandidateProvider struct {
	IUnknown
}

// ITfFnSearchCandidateProviderVtbl COM interface vtable.
type ITfFnSearchCandidateProviderVtbl struct {
	IUnknownVtbl
}

// ITfFnSearchCandidateProvider returns the ITfFnSearchCandidateProvider vtable.
func (obj *ITfFnSearchCandidateProvider) ITfFnSearchCandidateProvider() *ITfFnSearchCandidateProviderVtbl {
	return (*ITfFnSearchCandidateProviderVtbl)(unsafe.Pointer(obj.Vtbl))
}
