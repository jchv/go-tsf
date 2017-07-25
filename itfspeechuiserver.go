package tsf

import "unsafe"

// ITfSpeechUIServer COM interface.
type ITfSpeechUIServer struct {
	IUnknown
}

// ITfSpeechUIServerVtbl COM interface vtable.
type ITfSpeechUIServerVtbl struct {
	IUnknownVtbl
}

// ITfSpeechUIServer returns the ITfSpeechUIServer vtable.
func (obj *ITfSpeechUIServer) ITfSpeechUIServer() *ITfSpeechUIServerVtbl {
	return (*ITfSpeechUIServerVtbl)(unsafe.Pointer(obj.Vtbl))
}
