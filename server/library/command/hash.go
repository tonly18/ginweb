package command

import (
	"github.com/spaolacci/murmur3"
)

// HashMurmur32: MurmurHash3
func HashMurmur32(str string) uint32 {
	return murmur3.Sum32([]byte(str))
}
