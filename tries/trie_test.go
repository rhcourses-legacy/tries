package tries

import (
	"reflect"
	"testing"
)

// TestTrie_New tests the NewTrie function.
// It creates a new trie and checks whether it is empty and valid.
func TestTrie_New(t *testing.T) {
	trie := NewTrie()
	if !trie.IsEmpty() {
		t.Error("New trie should be empty")
	}
	if !trie.IsValid() {
		t.Error("New trie should be valid")
	}
}

// TestTrie_IsEmpty tests the IsEmpty function.
// It creates a new trie and checks whether it is empty.
// Then it adds a child node and checks whether the trie is not empty.
func TestTrie_IsEmpty(t *testing.T) {
	trie := NewTrie()
	if !trie.IsEmpty() {
		t.Error("New trie should be empty")
	}
	trie.root.children['a'] = NewNode()
	if trie.IsEmpty() {
		t.Error("Trie should not be empty")
	}
}

// TestTrie_Insert_nodata tests the Insert function.
// It creates a new trie and inserts a node withoud data.
// It checks whether the trie is not empty and valid.
func TestTrie_Insert_nodata(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	if trie.IsEmpty() {
		t.Error("Trie should not be empty")
	}
	if !trie.IsValid() {
		t.Error("Trie should be valid")
	}
}

// TestTrie_Insert_data tests the Insert function.
// It creates a new trie and inserts a node with data.
// Then it gets the data and checks whether it is correct.
func TestTrie_Insert_data(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc", "data1", "data2")

	expecteddata := []any{"data1", "data2"}
	actualdata := trie.GetData("abc")
	if !reflect.DeepEqual(actualdata, expecteddata) {
		t.Errorf("Data should be %v, but is %v", expecteddata, actualdata)
	}
}

// TestTrie_GetData_nonexisting tests the GetData function.
// It creates a new trie and asks for data of a non-existing node.
func TestTrie_GetData_nonexisting(t *testing.T) {
	trie := NewTrie()
	if trie.GetData("abc") != nil {
		t.Error("Data should be nil")
	}
}
