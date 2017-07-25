package tsf

import "unsafe"

// ITfInputProcessorProfiles COM interface.
type ITfInputProcessorProfiles struct {
	IUnknown
}

// ITfInputProcessorProfilesVtbl COM interface vtable.
type ITfInputProcessorProfilesVtbl struct {
	IUnknownVtbl
}

// ITfInputProcessorProfiles returns the ITfInputProcessorProfiles vtable.
func (obj *ITfInputProcessorProfiles) ITfInputProcessorProfiles() *ITfInputProcessorProfilesVtbl {
	return (*ITfInputProcessorProfilesVtbl)(unsafe.Pointer(obj.Vtbl))
}
