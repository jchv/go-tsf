package tsf

import "unsafe"

// ITfRangeBackup COM interface.
type ITfRangeBackup struct {
	IUnknown
}

// ITfRangeBackupVtbl COM interface vtable.
type ITfRangeBackupVtbl struct {
	IUnknownVtbl
}

// ITfRangeBackup returns the ITfRangeBackup vtable.
func (obj *ITfRangeBackup) ITfRangeBackup() *ITfRangeBackupVtbl {
	return (*ITfRangeBackupVtbl)(unsafe.Pointer(obj.Vtbl))
}
