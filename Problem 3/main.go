package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstruct([]string{"the", "quick", "brown", "fox"}, "thequickbrownfox"))
	fmt.Println(reconstruct([]string{"bed", "bath", "and", "beyond", "bedbath"}, "bedbathandbeyond"))
}

// To solve this problem we first sort the list of words by length, from largest to smallest.
// This lets us check the largest possible word first, so we don't have a case where
// the largest word something like "Photography" gets cut in half by a smaller word "Photo",
// leaving us with the word "graphy" that will likely not be in the given word list.
// If we find a matching word, we add it to the output, and remove it from the sentence, and
// continue checking the new first set of characters to see if it matches any of our words.
// This will continue until we have all the words in the sentence, or if there is a point in
// the sentence where no words are found, we will instead return an empty string because there
// is no possible reconstruction.
func reconstruct(words []string, sentence string) []string {
	output := []string{}

	//first sort words by length (largest to smallest)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	//loop until we have removed the contents of sentence
	for sentence != "" {
		found := false
		for _, w := range words {
			if len(w) <= len(sentence) && w == sentence[:len(w)] {
				output = append(output, w)
				sentence = sentence[len(w):]
				found = true
				break
			}
		}

		// if we hit this block, we've looped through the entire
		// word list without finding a word that fits in the
		// remainder of the sentence. Thus there is no possible
		// reconstruction
		if !found {
			return []string{}
		}
	}
	return output
}
