package tsf

import "unsafe"

// ITfInputProcessorProfilesEx COM interface.
type ITfInputProcessorProfilesEx struct {
	IUnknown
}

// ITfInputProcessorProfilesExVtbl COM interface vtable.
type ITfInputProcessorProfilesExVtbl struct {
	IUnknownVtbl
}

// ITfInputProcessorProfilesEx returns the ITfInputProcessorProfilesEx vtable.
func (obj *ITfInputProcessorProfilesEx) ITfInputProcessorProfilesEx() *ITfInputProcessorProfilesExVtbl {
	return (*ITfInputProcessorProfilesExVtbl)(unsafe.Pointer(obj.Vtbl))
}
