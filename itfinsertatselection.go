package tsf

import "unsafe"

// ITfInsertAtSelection COM interface.
type ITfInsertAtSelection struct {
	IUnknown
}

// ITfInsertAtSelectionVtbl COM interface vtable.
type ITfInsertAtSelectionVtbl struct {
	IUnknownVtbl
}

// ITfInsertAtSelection returns the ITfInsertAtSelection vtable.
func (obj *ITfInsertAtSelection) ITfInsertAtSelection() *ITfInsertAtSelectionVtbl {
	return (*ITfInsertAtSelectionVtbl)(unsafe.Pointer(obj.Vtbl))
}
