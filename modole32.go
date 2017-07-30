package tsf

import (
	"syscall"
	"unsafe"
)

// ClsCtx specifies the mode of activation for a COM class.
type ClsCtx uint32

// Possible class contexts
const (
	ClsCtxInProcServer        ClsCtx = 0x1
	ClsCtxInProcHandler              = 0x2
	ClsCtxLocalServer                = 0x4
	ClsCtxInProcServer16             = 0x8
	ClsCtxRemoteServer               = 0x10
	ClsCtxInProcHandler16            = 0x20
	ClsCtxReserved1                  = 0x40
	ClsCtxReserved2                  = 0x80
	ClsCtxReserved3                  = 0x100
	ClsCtxReserved4                  = 0x200
	ClsCtxNoCodeDownload             = 0x400
	ClsCtxReserved5                  = 0x800
	ClsCtxNoCustomMarshal            = 0x1000
	ClsCtxEnableCodeDownload         = 0x2000
	ClsCtxNoFailureLOG               = 0x4000
	ClsCtxDisableAAA                 = 0x8000
	ClsCtxEnableAAA                  = 0x10000
	ClsCtxFromDefaultContext         = 0x20000
	ClsCtxActivate32BitServer        = 0x40000
	ClsCtxActivate64BitServer        = 0x80000
	ClsCtxEnableCloaking             = 0x100000
	ClsCtxAppContainer               = 0x400000
	ClsCtxActivateAAAAsIU            = 0x800000
	ClsCtxPSDLL                      = 0x80000000
)

var (
	modole32 = syscall.NewLazyDLL("ole32.dll")

	procCoInitialize     = modole32.NewProc("CoInitialize")
	procCoCreateInstance = modole32.NewProc("CoCreateInstance")
	procCoTaskMemAlloc   = modole32.NewProc("CoTaskMemAlloc")
	procCoTaskMemFree    = modole32.NewProc("CoTaskMemFree")
)

// CoInitialize initializes the COM library on the current thread and
// identifies the concurrency model as single-thread apartment (STA).
func CoInitialize() error {
	hr, _, _ := procCoInitialize.Call(0)
	return result(hr)
}

// CoCreateInstance creates a single uninitialized object of the class
// associated with a specified CLSID.
func CoCreateInstance(clsid *GUID, outer uintptr, ctx ClsCtx, iid *GUID, out uintptr) (err error) {
	hr, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(clsid)),
		uintptr(outer),
		uintptr(ctx),
		uintptr(unsafe.Pointer(iid)),
		uintptr(out))
	return result(hr)
}

// CoTaskMemAlloc allocates a block of task memory in the same way that
// IMalloc::Alloc does.
func CoTaskMemAlloc(size uintptr) unsafe.Pointer {
	ret, _, _ := procCoTaskMemAlloc.Call(size)
	return unsafe.Pointer(ret)
}

// CoTaskMemFree frees a block of task memory previously allocated through a
// call to the CoTaskMemAlloc or CoTaskMemRealloc function.
func CoTaskMemFree(mem unsafe.Pointer) {
	procCoTaskMemFree.Call(uintptr(mem))
}
