package simplezip

import "testing"

// TestUnzipper_New_nonempty_done tests the NewUnzipper function.
// It creates a new unzipper with a non-empty string and checks whether it is done.
func TestUnzipper_New_nonempty_done(t *testing.T) {
	uz := NewUnzipper("abc")
	if uz.Done() {
		t.Error("New unzipper should not be done")
	}
}

// TestUnzipper_New_empty_done tests the NewUnzipper function.
// It creates a new unzipper with an empty string and checks whether it is done.
func TestUnzipper_New_empty_done(t *testing.T) {
	uz := NewUnzipper("")
	if !uz.Done() {
		t.Error("New unzipper should be done")
	}
}

// TestUnzipper_New_result tests the NewUnzipper function.
// It creates a new unzipper with a non-empty string and checks whether the result is empty.
func TestUnzipper_New_result(t *testing.T) {
	uz := NewUnzipper("abc")
	if uz.Result() != "" {
		t.Error("New unzipper should have empty result")
	}
}
