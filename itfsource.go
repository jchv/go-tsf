package tsf

import "unsafe"

// ITfSource COM interface.
type ITfSource struct {
	IUnknown
}

// ITfSourceVtbl COM interface vtable.
type ITfSourceVtbl struct {
	IUnknownVtbl
	AdviseSink   uintptr
	UnadviseSink uintptr
}

// ITfSource returns the ITfSource vtable.
func (obj *ITfSource) ITfSource() *ITfSourceVtbl {
	return (*ITfSourceVtbl)(unsafe.Pointer(obj.Vtbl))
}

// AdviseSink installs an advise sink.
func (obj *ITfSource) AdviseSink(riid *GUID, sink *COMInstance, cookie *uint32) error {
	return result(call3(
		obj.ITfSource().AdviseSink,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(sink)),
		uintptr(unsafe.Pointer(cookie)),
	))
}
