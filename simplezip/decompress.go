package simplezip

import (
	"github.com/rhcourses/tries/tries"
)

// Decompress expects a compressed string and returns the decompressed string.
// It also expects a Trie that is used as a dictionary.
func Decompress(compressedString string, trie *tries.Trie) string {
	uz := NewUnzipper(compressedString, trie)
	uz.Run()

	return uz.Result()
}
