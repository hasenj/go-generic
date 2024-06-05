package generic

import (
	"unsafe"
)

// UnsafeRawBytes returns the raw bytes behind the object
func UnsafeRawBytes[T any](v *T) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(v)), unsafe.Sizeof(*v))
}

// UnsafeRawBytesOffset returns the raw bytes behind an object, skipping the
// first few fields
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

// UnsafeSliceBytes takes a slice of any type and reinterprets it as a slice of
// bytes
func UnsafeSliceBytes[T any](v []T) []byte {
	if len(v) == 0 {
		return nil
	}
	size := int(unsafe.Sizeof(v[0])) * len(v)
	data := (*byte)(unsafe.Pointer(&v[0]))
	return unsafe.Slice(data, uintptr(size))
}

// UnsafeStringBytes converts a string to a byte slice without copying
func UnsafeStringBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnsafeString converts a byte slice to a string without copying
func UnsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
