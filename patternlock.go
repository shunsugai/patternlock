package patternlock

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"github.com/kokardy/listing"
	"runtime"
)

const max int = 9
const min int = 4

func Calc(hash []byte) []byte {
	list := ByteReplacer([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})
	buf := runtime.NumCPU()

	for i := min; i <= max; i++ {
		for p := range listing.Permutations(list, i, false, buf) {
			pattern := ByteReplacer(p.(ByteReplacer))
			tmp := sha1.Sum(pattern)
			if bytes.Equal(hash, tmp[:]) {
				return pattern
			}
		}
	}
	return []byte{}
}

func PrintResult(res []byte) {
	var ary [9]int
	for idx, val := range res {
		ary[int(val)] = idx + 1
	}
	fmt.Println("Pattern  : ", res)
	fmt.Println(" --- --- ---")
	fmt.Println("|", ary[0], "|", ary[1], "|", ary[2], "|")
	fmt.Println(" ---+---+---")
	fmt.Println("|", ary[3], "|", ary[4], "|", ary[5], "|")
	fmt.Println(" ---+---+---")
	fmt.Println("|", ary[6], "|", ary[7], "|", ary[8], "|")
	fmt.Println(" --- --- ---")
}
