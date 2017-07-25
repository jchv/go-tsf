package tsf

import "unsafe"

// ITfInputProcessorProfileSubstituteLayout COM interface.
type ITfInputProcessorProfileSubstituteLayout struct {
	IUnknown
}

// ITfInputProcessorProfileSubstituteLayoutVtbl COM interface vtable.
type ITfInputProcessorProfileSubstituteLayoutVtbl struct {
	IUnknownVtbl
}

// ITfInputProcessorProfileSubstituteLayout returns the ITfInputProcessorProfileSubstituteLayout vtable.
func (obj *ITfInputProcessorProfileSubstituteLayout) ITfInputProcessorProfileSubstituteLayout() *ITfInputProcessorProfileSubstituteLayoutVtbl {
	return (*ITfInputProcessorProfileSubstituteLayoutVtbl)(unsafe.Pointer(obj.Vtbl))
}
