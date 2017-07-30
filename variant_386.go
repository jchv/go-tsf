//+build 386

package tsf

type Variant struct {
	VT         VT
	wReserved1 uint16
	wReserved2 uint16
	wReserved3 uint16
	Val        int64
}
