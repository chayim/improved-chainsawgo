package smurf

// #include "resources/libcommands.h"
// #cgo LDFLAGS: -lresources/linux-x86/libcommands
import "C"

func Chicken(x, y int32) int {
	_x := C.int(x)
	_y := C.int(y)
	i := C.numbers(_x, _y)
	return int(i)
}
