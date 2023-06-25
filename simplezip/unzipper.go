package simplezip

import "github.com/rhcourses/tries/tries"

// Unzipper is a data structure that stores a compressed string and everything that is
// needed to track the progress while decompressing it.

type Unzipper struct {
	compressed string

	tw       *tries.TrieWalker
	position int

	result string
}

// NewUnzipper creates a new unzipper.
func NewUnzipper(compressed string) *Unzipper {
	return &Unzipper{
		tw:         tries.NewTrieWalker(tries.NewTrie()),
		compressed: compressed,
	}
}

// Done returns true if the unzipper is done.
func (uz *Unzipper) Done() bool {
	return uz.position == len(uz.compressed)
}

// Result returns the decompressed string.
func (uz *Unzipper) Result() string {
	return uz.result
}
