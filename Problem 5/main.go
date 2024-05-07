package problem5

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// Case 1: expect: `["the quick", "brown fox", "jumps over", "the lazy", "dog"]`
	lines := splitSentence("the quick brown fox jumps over the lazy dog", 10)
	if reflect.DeepEqual(lines, []string{"the quick", "brown fox", "jumps over", "the lazy", "dog"}) {
		fmt.Println("Case 1 Success")
	}

	// Case 2: expect nil
	lines = splitSentence("the quickasdfasdf brown fox jumps over the lazy dog", 10)
	if lines == nil {
		fmt.Println("Case 2 Success")
	}

}

func splitSentence(s string, k int) []string {
	lines := []string{}

	// first let's split the string by all space characters to get a list of words:
	words := strings.Split(s, " ")

	// current line will hold the current list of words we're working with. We'll set it to the first word to begin, as long as that word
	// isn't longer than the limit.
	var currentLine string
	if len(words) > 0 && len(words[0]) <= k {
		currentLine = words[0]
	}

	// now we'll loop through them and put them together if they fit in the length limit
	for _, word := range words[1:] {
		// the only way this is impossible is if the length of a single word is longer than the limit. In this case, we return nil.
		if len(word) > k {
			return nil
		}

		// as long as our current word isn't longer than the limit, we can now decide to either add it to the currentLine or make a new line:
		// if the length of the current line plus the current word plus one (for the space) is longer than the limit, we should add
		// currentLine to our output list and make a new current line.
		if len(currentLine)+len(word)+1 > k {
			lines = append(lines, currentLine)
			currentLine = word
		} else { //if a space and the next word does fit, then we'll add it to currentLine and continue.
			currentLine += " " + word
		}
	}
	// if there is leftover in currentLine, add it to lines
	if currentLine != lines[len(lines)-1] {
		lines = append(lines, currentLine)
	}

	return lines
}
