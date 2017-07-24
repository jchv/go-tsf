package tsf

import (
	"unsafe"
)

// IAnchor represents the IAnchor COM interface.
type IAnchor struct {
	IUnknown
}

// IAnchorVtbl COM interface vtable.
type IAnchorVtbl struct {
	IUnknownVtbl
	SetGravity           uintptr
	GetGravity           uintptr
	IsEqual              uintptr
	Compare              uintptr
	Shift                uintptr
	ShiftTo              uintptr
	ShiftRegion          uintptr
	SetChangeHistoryMask uintptr
	GetChangeHistory     uintptr
	ClearChangeHistory   uintptr
	Clone                uintptr
}

// IAnchor returns the IAnchor vtable.
func (obj *IAnchor) IAnchor() *IAnchorVtbl {
	return (*IAnchorVtbl)(unsafe.Pointer(obj.Vtbl))
}

// SetGravity sets the gravity of the anchor.
func (obj *IAnchor) SetGravity(gravity TSGravity) error {
	return result(call1(
		uintptr(obj.IAnchor().SetGravity),
		uintptr(unsafe.Pointer(obj)),
		uintptr(gravity),
	))
}

// GetGravity Retrieves the gravity of the anchor.
func (obj *IAnchor) GetGravity(gravity *TSGravity) error {
	return result(call1(
		uintptr(obj.IAnchor().GetGravity),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(gravity)),
	))
}

// IsEqual specifies the equality or inequality of the positions of two anchors.
func (obj *IAnchor) IsEqual(with *IAnchor, equal *int32) error {
	return result(call2(
		uintptr(obj.IAnchor().IsEqual),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(with)),
		uintptr(unsafe.Pointer(equal)),
	))
}

// Compare compares the relative position of two anchors within a text stream.
func (obj *IAnchor) Compare(with *IAnchor, out *int32) error {
	return result(call2(
		uintptr(obj.IAnchor().Compare),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(with)),
		uintptr(unsafe.Pointer(out)),
	))
}

// Shift shifts the anchor forward or backward.
func (obj *IAnchor) Shift(flags uint32, req int32, ch *int32, halt *IAnchor) error {
	return result(call4(
		uintptr(obj.IAnchor().Shift),
		uintptr(unsafe.Pointer(obj)),
		uintptr(flags),
		uintptr(req),
		uintptr(unsafe.Pointer(ch)),
		uintptr(unsafe.Pointer(halt)),
	))
}

// ShiftTo shifts the current anchor to the same position as another anchor.
func (obj *IAnchor) ShiftTo(site *IAnchor) error {
	return result(call1(
		uintptr(obj.IAnchor().ShiftTo),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(site)),
	))
}

// ShiftRegion shifts the anchor into an adjacent region in the text stream.
func (obj *IAnchor) ShiftRegion(flags uint32, dir TSShiftDir, noregion *int32) error {
	return result(call3(
		uintptr(obj.IAnchor().ShiftRegion),
		uintptr(unsafe.Pointer(obj)),
		uintptr(flags),
		uintptr(dir),
		uintptr(unsafe.Pointer(noregion)),
	))
}

// SetChangeHistoryMask is not implemented.
func (obj *IAnchor) SetChangeHistoryMask(mask uint32) error {
	return result(call1(
		uintptr(obj.IAnchor().SetChangeHistoryMask),
		uintptr(unsafe.Pointer(obj)),
		uintptr(mask),
	))
}

// GetChangeHistory gets the change history of deletions that have occurred
// immediately preceding or following the anchor.
func (obj *IAnchor) GetChangeHistory(history *uint32) error {
	return result(call1(
		uintptr(obj.IAnchor().SetChangeHistoryMask),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(history)),
	))
}

// ClearChangeHistory clears the change history flags.
func (obj *IAnchor) ClearChangeHistory() error {
	return result(call0(
		uintptr(obj.IAnchor().SetChangeHistoryMask),
		uintptr(unsafe.Pointer(obj)),
	))

}

// Clone produces a new anchor object positioned at the same location, and
// with the same gravity, as the current anchor.
func (obj *IAnchor) Clone(clone **IAnchor) error {
	return result(call1(
		uintptr(obj.IAnchor().SetChangeHistoryMask),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(clone)),
	))
}
