package tsf

import "unsafe"

// ITfConfigureSystemKeystrokeFeed COM interface.
type ITfConfigureSystemKeystrokeFeed struct {
	IUnknown
}

// ITfConfigureSystemKeystrokeFeedVtbl COM interface vtable.
type ITfConfigureSystemKeystrokeFeedVtbl struct {
	IUnknownVtbl
}

// ITfConfigureSystemKeystrokeFeed returns the ITfConfigureSystemKeystrokeFeed vtable.
func (obj *ITfConfigureSystemKeystrokeFeed) ITfConfigureSystemKeystrokeFeed() *ITfConfigureSystemKeystrokeFeedVtbl {
	return (*ITfConfigureSystemKeystrokeFeedVtbl)(unsafe.Pointer(obj.Vtbl))
}
