package simplezip

import (
	"testing"

	"github.com/rhcourses/tries/tries"
)

// TestUnzipper_New_nonempty_done tests the NewUnzipper function.
// It creates a new unzipper with a non-empty string and checks whether it is done.
func TestUnzipper_New_nonempty_done(t *testing.T) {
	uz := NewUnzipper("abc", tries.NewTrie())
	if uz.Done() {
		t.Error("New unzipper should not be done")
	}
}

// TestUnzipper_New_empty_done tests the NewUnzipper function.
// It creates a new unzipper with an empty string and checks whether it is done.
func TestUnzipper_New_empty_done(t *testing.T) {
	uz := NewUnzipper("", tries.NewTrie())
	if !uz.Done() {
		t.Error("New unzipper should be done")
	}
}

// TestUnzipper_New_result tests the NewUnzipper function.
// It creates a new unzipper with a non-empty string and checks whether the result is empty.
func TestUnzipper_New_result(t *testing.T) {
	uz := NewUnzipper("abc", tries.NewTrie())
	if uz.Result() != "" {
		t.Error("New unzipper should have empty result")
	}
}

// TestUnzipper_AdvanceSingleChar tests the AdvanceSingleChar function.
// It creates a new unzipper with a non-empty string, advances it by one character,
// and checks whether the result is correct.
func TestUnzipper_AdvanceSingleChar(t *testing.T) {
	uz := NewUnzipper("abc", tries.NewTrie())
	uz.AdvanceSingleChar()
	if uz.Result() != "a" {
		t.Error("Result should be \"a\"")
	}
}

// TestUnzipper_InsertCurrentData tests the InsertCurrentData function.
// It creates a new unzipper with a non-empty string and a trie with data.
// It advances the unzipper to the data node and calls InsertCurrentData.
// It checks whether the result is correct.
func TestUnzipper_InsertCurrentData(t *testing.T) {
	trie := tries.NewTrie()
	trie.Insert("1", "abc")
	uz := NewUnzipper("1", trie)
	uz.tw.Step('1')
	uz.InsertCurrentData()
	if uz.Result() != "abc" {
		t.Errorf("Result should be \"%v\", but is \"%v\".", "abc", uz.Result())
	}
}

// TestUnzipper_Run_clean tests the Run function.
// It creates a new unzipper with a string, calls Run, and checks whether the result is correct.
// In this test case, every single character is the code for a string in the dictionary.
func TestUnzipper_Run_clean(t *testing.T) {
	original := "aaabbbcccaaabbbcccaaabbbccc"
	compressed := "123123123"

	trie := tries.NewTrie()
	trie.Insert("1", "aaa")
	trie.Insert("2", "bbb")
	trie.Insert("3", "ccc")

	uz := NewUnzipper(compressed, trie)
	uz.Run()

	if uz.Result() != original {
		t.Errorf("Decompressed string should be \"%s\", but is \"%s\"", original, uz.Result())
	}
}

// TestUnzipper_Run_dirty tests the Run function.
// It creates a new unzipper with a string, calls Run, and checks whether the result is correct.
// In this test case, the codes are mixed with other characters that do not have to be unpacked.
func TestUnzipper_Run_dirty(t *testing.T) {
	original := "aaabbbcccaaabbbcccaaabbbccc"
	compressed := "12ccc12ccc12ccc"

	trie := tries.NewTrie()
	trie.Insert("1", "aaa")
	trie.Insert("2", "bbb")

	uz := NewUnzipper(compressed, trie)
	uz.Run()

	if uz.Result() != original {
		t.Errorf("Decompressed string should be \"%s\", but is \"%s\"", original, uz.Result())
	}
}

// TestUnzipper_Run_prefix tests the Run function.
// It creates a new unzipper with a string, calls Run, and checks whether the result is correct.
// In this test case, the compressed string contains a prefix of a string in the dictionary.
func TestUnzipper_Run_prefix(t *testing.T) {
	original := "aaa#bbb#ccc"
	compressed := "#1##2##3"

	trie := tries.NewTrie()
	trie.Insert("#1", "aaa")
	trie.Insert("#2", "bbb")
	trie.Insert("#3", "ccc")

	uz := NewUnzipper(compressed, trie)
	uz.Run()

	if uz.Result() != original {
		t.Errorf("Decompressed string should be \"%s\", but is \"%s\"", original, uz.Result())
	}
}
