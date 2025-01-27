package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import (
	"time"
	"unsafe"
)

type TGeogPointInst struct {
	_inner *C.Temporal
}

func NewTGeogPointInst(tgmpi_in string) *TGeogPointInst {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := TGeogPointInst{_inner: c_tgmpi}
	return &g_tgmpi
}

func (tgmpi *TGeogPointInst) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tgmpi *TGeogPointInst) Inner() *C.Temporal {
	return tgmpi._inner
}

func (tgmpi *TGeogPointInst) Init(c_temp *C.Temporal) {
	tgmpi._inner = c_temp
}

func (tgmpi *TGeogPointInst) IsTGeogPoint() bool {
	return true
}

func (tgmpi *TGeogPointInst) IsTPoint() bool {
	return true
}

func (tgmpi *TGeogPointInst) String() string {
	return tgmpi.TPointOut(10)
}

func (tgmpi *TGeogPointInst) Type() string {
	return "TGeogPointInst"
}

func (tgmpi *TGeogPointInst) IsTInstant() bool {
	return true
}

func (tgmpi *TGeomPointInst) Timestamptz() time.Time {
	c_inst := C.cast_temporal_to_tinstant(tgmpi._inner)
	return TimestamptzToDatetime(c_inst.t)
}

func (tgmpi *TGeomPointInst) TimestampOut() string {
	c_inst := C.cast_temporal_to_tinstant(tgmpi._inner)
	return C.GoString(C.pg_timestamptz_out(c_inst.t))
}

type TGeogPointSeq struct {
	_inner *C.Temporal
}

func NewTGeogPointSeq(tgmpi_in string) TGeogPointSeq {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := TGeogPointSeq{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointSeq) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tgmpi *TGeogPointSeq) Inner() *C.Temporal {
	return tgmpi._inner
}

func (tgmpi *TGeogPointSeq) Init(c_temp *C.Temporal) {
	tgmpi._inner = c_temp
}

func (tgmpi *TGeogPointSeq) IsTGeogPoint() bool {
	return true
}

func (tgmpi *TGeogPointSeq) IsTPoint() bool {
	return true
}

func (tgmpi *TGeogPointSeq) String() string {
	return tgmpi.TPointOut(10)
}

func (tgmpi *TGeogPointSeq) Type() string {
	return "TGeogPointSeq"
}

func (tgmpi *TGeogPointSeq) IsTSequence() bool {
	return true
}

type TGeogPointSeqSet struct {
	_inner *C.Temporal
}

func NewTGeogPointSeqSet(tgmpi_in string) *TGeogPointSeqSet {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := &TGeogPointSeqSet{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointSeqSet) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tgmpi *TGeogPointSeqSet) Inner() *C.Temporal {
	return tgmpi._inner
}

func (tgmpi *TGeogPointSeqSet) Init(c_temp *C.Temporal) {
	tgmpi._inner = c_temp
}

func (tgmpi *TGeogPointSeqSet) IsTGeogPoint() bool {
	return true
}

func (tgmpi *TGeogPointSeqSet) IsTPoint() bool {
	return true
}

func (tgmpi *TGeogPointSeqSet) String() string {
	return tgmpi.TPointOut(10)
}

func (tgmpi *TGeogPointSeqSet) Type() string {
	return "TGeogPointSeqSet"
}

func TGeogPointIn[TG TGeogPoint](input string, output TG) TG {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_geogpoint := C.tgeogpoint_in(c_str)
	output.Init(c_geogpoint)
	return output
}

func TGeogPointFromMFJSON[TG TGeogPoint](input string, output TG) TG {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_geogpoint := C.tgeogpoint_from_mfjson(c_str)
	output.Init(c_geogpoint)
	return output
}
