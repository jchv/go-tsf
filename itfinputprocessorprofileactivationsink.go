package tsf

import "unsafe"

// ITfInputProcessorProfileActivationSink COM interface.
type ITfInputProcessorProfileActivationSink struct {
	IUnknown
}

// ITfInputProcessorProfileActivationSinkVtbl COM interface vtable.
type ITfInputProcessorProfileActivationSinkVtbl struct {
	IUnknownVtbl
}

// ITfInputProcessorProfileActivationSink returns the ITfInputProcessorProfileActivationSink vtable.
func (obj *ITfInputProcessorProfileActivationSink) ITfInputProcessorProfileActivationSink() *ITfInputProcessorProfileActivationSinkVtbl {
	return (*ITfInputProcessorProfileActivationSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
