package simplezip

import (
	"fmt"

	"github.com/rhcourses/tries/tries"
)

// Decompress expects a compressed string and returns the decompressed string.
// It also expects a Trie that is used as a dictionary.
func Decompress(compressedString string, trie *tries.Trie) string {
	tw := tries.NewTrieWalker(trie)
	result := ""

	for len(compressedString) > 0 {
		consumed := tw.Walk(compressedString)
		data := tw.Data()
		if len(data) != 0 {
			result += fmt.Sprintf("%v", data[0])
		}
		tw.Reset()
		compressedString = compressedString[consumed:]
	}

	return result
}
