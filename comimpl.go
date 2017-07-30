package tsf

import (
	"syscall"
	"unsafe"
)

// COMInstance represents a COM instance with its own refcount. There may be
// multiple COMInstances for each object, because (probably unlike C++ COM
// classes, but explicitly allowed by COM,) we return a different pointer
// for each COM interface of an object.
type COMInstance struct {
	IUnknown
	Base  *COMInstance
	Intfs map[GUID]*IUnknownVtbl
	Data  interface{}
	Ref   int
}

// AllocCOMInstance allocates an empty COM instance. The memory will be
// uninitialized, so be careful.
func AllocCOMInstance() *COMInstance {
	return (*COMInstance)(CoTaskMemAlloc(unsafe.Sizeof(COMInstance{})))
}

// FreeCOMInstance will release the memory used by a COM instance. Usually
// you should not do this. It will automatically get released if the program
// you are using is calling Release properly.
func FreeCOMInstance(inst *COMInstance) {
	CoTaskMemFree(unsafe.Pointer(inst))
}

// IUnknownBaseVtbl is a vtbl for implementing IUnknown with COMInstance.
// You can embed this into your own structures to implement COM objects using
// the COMInstance interface.
var IUnknownBaseVtbl = IUnknownVtbl{
	QueryInterface: syscall.NewCallback(func(this *COMInstance, riid *GUID, object *uintptr) uintptr {
		if object == nil {
			return EPointer
		}

		if *riid == IIDIUnknown {
			*object = uintptr(unsafe.Pointer(this.Base))
			return 0
		}

		// Try to find interface implementation.
		if vtbl, ok := this.Intfs[*riid]; ok {
			// If this is the same interface, add a ref and return ourselves.
			if this.Vtbl == vtbl {
				this.Ref++
				*object = uintptr(unsafe.Pointer(this))
				return 0
			}

			// Otherwise, return a new derived interface with its own instance.
			*object = uintptr(unsafe.Pointer(NewDerivedCOMInstance(vtbl, this)))
			return 0
		}

		return ENoInterface
	}),
	AddRef: syscall.NewCallback(func(this *COMInstance) uintptr {
		this.Ref++

		return 0
	}),
	Release: syscall.NewCallback(func(this *COMInstance) uintptr {
		this.Ref--
		if this.Ref == 0 {
			FreeCOMInstance(this)
		}

		return 0
	}),
}

// NewBaseCOMInstance creates a base COM instance with the given vtbl and
// interface vtables.
func NewBaseCOMInstance(vtbl *IUnknownVtbl, data interface{}, intfs map[GUID]*IUnknownVtbl) *COMInstance {
	impl := AllocCOMInstance()
	impl.IUnknown = IUnknown{Vtbl: vtbl}
	impl.Base = impl
	impl.Intfs = intfs
	impl.Data = data
	impl.Ref = 1
	return impl
}

// NewDerivedCOMInstance creates a derived-class COM instance with the given
// vtbl and base COM instance.
func NewDerivedCOMInstance(vtbl *IUnknownVtbl, base *COMInstance) *COMInstance {
	inst := AllocCOMInstance()
	inst.IUnknown = IUnknown{Vtbl: vtbl}
	inst.Base = base
	inst.Intfs = base.Intfs
	inst.Data = base.Data
	inst.Ref = 1
	return inst
}
