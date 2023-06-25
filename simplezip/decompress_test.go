package simplezip

import (
	"testing"

	"github.com/rhcourses/tries/tries"
)

// TestDecompress_clean tests the Decompress function.
// It unpacks a compressed string where every single character is the code for a string in the dictionary.
// It checks whether the decompressed string is correct.
func TestDecompress_clean(t *testing.T) {
	original := "aaabbbcccaaabbbcccaaabbbccc"
	compressed := "123123123"

	trie := tries.NewTrie()
	trie.Insert("1", "aaa")
	trie.Insert("2", "bbb")
	trie.Insert("3", "ccc")

	decompressed := Decompress(compressed, trie)
	if decompressed != original {
		t.Errorf("Decompressed string should be \"%s\", but is \"%s\"", original, decompressed)
	}
}

// TestDecompress_dirty tests the Decompress function.
// It unpacks a compressed string where the codes are mixed with other characters
// that do not have to be unpacked and should be copied to the decompressed string.
func TestDecompress_dirty(t *testing.T) {
	original := "aaabbbcccaaabbbcccaaabbbccc"
	compressed := "12ccc12ccc12ccc"

	trie := tries.NewTrie()
	trie.Insert("1", "aaa")
	trie.Insert("2", "bbb")

	decompressed := Decompress(compressed, trie)
	if decompressed != original {
		t.Errorf("Decompressed string should be \"%s\", but is \"%s\"", original, decompressed)
	}
}

// TestDecompress_prefix tests the Decompress function.
// It unpacks a compressed string that contains a prefix of a string in the dictionary.
// It checks whether the decompressed string is correct.
func TestDecompress_prefix(t *testing.T) {
	original := "aaa#bbb#ccc"
	compressed := "#1##2##3"

	trie := tries.NewTrie()
	trie.Insert("#1", "aaa")
	trie.Insert("#2", "bbb")
	trie.Insert("#3", "ccc")

	decompressed := Decompress(compressed, trie)
	if decompressed != original {
		t.Errorf("Decompressed string should be \"%s\", but is \"%s\"", original, decompressed)
	}
}
