package tsf

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
