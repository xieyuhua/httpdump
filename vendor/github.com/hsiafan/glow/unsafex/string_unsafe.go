package unsafex

import (
	"reflect"
	"runtime"
	"unsafe"
)

// BytesToString unsafe convert byte array content to string
// This function do not allocate new buffer, reuse the bytes data.
func BytesToString(bytes []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	sp := (*string)(unsafe.Pointer(&stringHeader))
	runtime.KeepAlive(&bytes)
	return *sp
}

// StringToBytes unsafe convert string to byte array
// This function do not allocate new buffer, reuse the string buffer.
func StringToBytes(str string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len, // just same with len
	}
	sp := (*[]byte)(unsafe.Pointer(&sliceHeader))
	runtime.KeepAlive(&str)
	return *sp
}

// CopyString copy a string content, for reducing large string content memory usage when do substring.
// This method allocate a new string content byte array, thereby allow the larger string to be released by the garbage collector once it is no longer referenced
func CopyString(s string) string {
	return string(StringToBytes(s))
}
