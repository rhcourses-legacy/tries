package simplezip

import (
	"fmt"

	"github.com/rhcourses/tries/tries"
)

// ExampleDecompress_flat shows how to decompress a string using a "flat" dictionary.
// "flat" means that the words in the dictionary are not prefixes of other words.
func ExampleDecompress_flat() {

	// build dictionary
	dict := tries.NewTrie()
	dict.Insert("1", "Fischer's Fritz")
	dict.Insert("2", "fischt")
	dict.Insert("3", "frische Fische")

	// define compressed string
	compressed := "1 2 3, 3 2 1."

	// decompress and print result
	decompressed := Decompress(compressed, dict)
	fmt.Println(decompressed)

	// Output:
	// Fischer's Fritz fischt frische Fische, frische Fische fischt Fischer's Fritz.
}

// ExampleDecompress_prefix shows how to decompress a string using a "prefix" dictionary.
// "prefix" means that the words in the dictionary are prefixes of other words.
func ExampleDecompress_prefix() {

	// build dictionary
	dict := tries.NewTrie()
	dict.Insert("1", "Fischer's Fritz.")
	dict.Insert("11", "Fischer's Fritz fischt frische Fische, ")
	dict.Insert("12", "frische Fische ")
	dict.Insert("2", "fischt ")

	// define compressed string
	compressed := "111221"

	// decompress and print result
	decompressed := Decompress(compressed, dict)
	fmt.Println(decompressed)

	// Output:
	// Fischer's Fritz fischt frische Fische, frische Fische fischt Fischer's Fritz.
}
