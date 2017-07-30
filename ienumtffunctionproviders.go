package tsf

import "unsafe"

// IEnumTfFunctionProviders represents the IEnumTfFunctionProviders COM interface.
type IEnumTfFunctionProviders struct {
	IUnknown
}

// IEnumTfFunctionProvidersVtbl COM interface vtable.
type IEnumTfFunctionProvidersVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfFunctionProviders returns the IEnumTfFunctionProviders vtable.
func (obj *IEnumTfFunctionProviders) IEnumTfFunctionProviders() *IEnumTfFunctionProvidersVtbl {
	return (*IEnumTfFunctionProvidersVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfFunctionProviders) Clone(enum **IEnumTfFunctionProviders) error {
	return result(call1(
		uintptr(obj.IEnumTfFunctionProviders().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfFunctionProviders) Next(count uint32, items **ITfFunctionProvider, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfFunctionProviders().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfFunctionProviders) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfFunctionProviders().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfFunctionProviders) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfFunctionProviders().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
