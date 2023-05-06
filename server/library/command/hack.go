package command

import (
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// no copy to change slice to string
// use your own risk
func String(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// no copy to change string to slice
// use your own risk
func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func IsSqlSep(r rune) bool {
	return r == ' ' || r == ',' ||
		r == '\t' || r == '/' ||
		r == '\n' || r == '\r'
}

func ArrayToString(array []int) string {
	if len(array) == 0 {
		return ""
	}
	var strArray []string
	for _, v := range array {
		strArray = append(strArray, strconv.FormatInt(int64(v), 10))
	}

	return strings.Join(strArray, ", ")
}
