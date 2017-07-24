package tsf

// ITextStoreACP represents the ITextStoreACP COM interface.
type ITextStoreACP struct {
	IUnknown
}

// ITextStoreACPVtbl COM interface vtable.
type ITextStoreACPVtbl struct {
	IUnknownVtbl
	AdviseSink                          uintptr
	UnadviseSink                        uintptr
	RequestLock                         uintptr
	GetStatus                           uintptr
	QueryInsert                         uintptr
	GetSelection                        uintptr
	SetSelection                        uintptr
	GetText                             uintptr
	SetText                             uintptr
	GetFormattedText                    uintptr
	GetEmbedded                         uintptr
	QueryInsertEmbedded                 uintptr
	InsertEmbedded                      uintptr
	InsertTextAtSelection               uintptr
	InsertEmbeddedAtSelection           uintptr
	RequestSupportedAttrs               uintptr
	RequestAttrsAtPosition              uintptr
	RequestAttrsTransitioningAtPosition uintptr
	FindNextAttrTransition              uintptr
	RetrieveRequestedAttrs              uintptr
	GetEndACP                           uintptr
	GetActiveView                       uintptr
	GetACPFromPoint                     uintptr
	GetTextExt                          uintptr
	GetScreenExt                        uintptr
	GetWnd                              uintptr
}
