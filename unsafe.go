package generic

import (
	"unsafe"
)

func UnsafeRawBytes[T any](v *T) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(v)), unsafe.Sizeof(*v))
}

func UnsafeRawBytesOffset[T, S any](v *T, memberPtr *S) []byte {
	offset := uintptr(unsafe.Pointer(memberPtr))
	base := uintptr(unsafe.Pointer(v))
	size := uintptr(unsafe.Sizeof(*v))
	diff := offset - base
	if size > offset {
		panic("Invalid offset")
	}
	sliceBase := unsafe.Pointer(memberPtr)
	sliceLen := int(size - diff)
	return unsafe.Slice((*byte)(sliceBase), sliceLen)
}

func UnsafeSliceBytes[T any](v []T) []byte {
	if len(v) == 0 {
		return nil
	}
	size := int(unsafe.Sizeof(v[0])) * len(v)
	data := (*byte)(unsafe.Pointer(&v[0]))
	return unsafe.Slice(data, uintptr(size))
}

func UnsafeStringBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func UnsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
