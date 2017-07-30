package main

import (
	"log"
	"syscall"
	"unsafe"

	"github.com/AllenDang/gform"
	"github.com/johnwchadwick/go-tsf"
)

var UIElementSinkVTbl = (*tsf.IUnknownVtbl)(unsafe.Pointer(&tsf.ITfUIElementSinkVtbl{
	IUnknownVtbl: tsf.IUnknownBaseVtbl,
	BeginUIElement: syscall.NewCallback(func(this *tsf.COMInstance, elemID uint32, show *uint32) uintptr {
		log.Printf("BeginUIElement: 0x%08x\n", elemID)
		return 0
	}),
	UpdateUIElement: syscall.NewCallback(func(this *tsf.COMInstance, elemID uint32) uintptr {
		log.Printf("UpdateUIElement: 0x%08x\n", elemID)
		return 0
	}),
	EndUIElement: syscall.NewCallback(func(this *tsf.COMInstance, elemID uint32) uintptr {
		log.Printf("EndUIElement: 0x%08x\n", elemID)
		return 0
	}),
}))

var UIElementSinkIntfs = map[tsf.GUID]*tsf.IUnknownVtbl{
	tsf.IIDTFUIElementSink: UIElementSinkVTbl,
}

func main() {
	gform.Init()

	mainWindow := gform.NewForm(nil)
	mainWindow.SetPos(300, 100)
	mainWindow.SetSize(500, 300)
	mainWindow.SetCaption("IME Demo")

	edt := gform.NewEdit(mainWindow)
	edt.SetPos(10, 10)

	mainWindow.OnClose().Bind(func(e *gform.EventArg) {
		mainWindow.Close()
		gform.Exit()
	})

	mainWindow.Show()

	var threadMgr *tsf.ITfThreadMgrEx
	tsf.CoInitialize()
	must(tsf.CoCreateInstance(
		&tsf.CLSIDTFThreadMgr,
		0,
		tsf.ClsCtxInProcServer,
		&tsf.IIDTFThreadMgrEx,
		uintptr(unsafe.Pointer(&threadMgr)),
	))

	var clientid tsf.TFClientID
	threadMgr.ActivateEx(&clientid, tsf.TFTMAEUIElementEnabledOnly)

	var source *tsf.ITfSource
	must(threadMgr.QueryInterface(
		&tsf.IIDTFSource,
		uintptr(unsafe.Pointer(&source)),
	))

	cookie := uint32(0)
	uiElemSink := tsf.NewBaseCOMInstance(UIElementSinkVTbl, nil, UIElementSinkIntfs)
	must(source.AdviseSink(&tsf.IIDTFUIElementSink, uiElemSink, &cookie))

	log.Println("UIElementSink cookie:", cookie)

	gform.RunMainLoop()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
