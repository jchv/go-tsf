package tsf

import "unsafe"

// ITfEditTransactionSink COM interface.
type ITfEditTransactionSink struct {
	IUnknown
}

// ITfEditTransactionSinkVtbl COM interface vtable.
type ITfEditTransactionSinkVtbl struct {
	IUnknownVtbl
}

// ITfEditTransactionSink returns the ITfEditTransactionSink vtable.
func (obj *ITfEditTransactionSink) ITfEditTransactionSink() *ITfEditTransactionSinkVtbl {
	return (*ITfEditTransactionSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
