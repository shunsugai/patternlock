package patternlock

import (
	"github.com/kokardy/listing"
)

type ByteReplacer []byte

func (br ByteReplacer) Len() int {
	return len(br)
}

func (br ByteReplacer) Replace(indices []int) listing.Replacer {
	result := make(ByteReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = br[idx]
	}
	return result
}
