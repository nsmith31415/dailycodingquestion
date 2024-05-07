package problem7

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(isValid("internationalization", "i18n"))
	fmt.Println(isValid("internationalization", "i2e15n"))
	fmt.Println(isValid("internationalization", "i4n15n"))
}

func isValid(word, numeronym string) bool {
	isnumber, _ := regexp.Compile("^[0123456789]+")
	wordIndex := 0
	i := 0
	for i < len(numeronym) {
		if isnumber.MatchString(numeronym[i:]) { // if current character is a number, find the number value
			number := isnumber.FindString((numeronym[i:]))
			value, err := strconv.Atoi(number)
			if err != nil {
				return false
			}
			//increment indices
			wordIndex += value
			i += len(number)
		} else if word[wordIndex] != numeronym[i] { // if current character is not a number, make sure it matches word value
			return false
		}
		// if current character is not a number and it successfully matched up with word, increment index
		i++
	}
	return true
}
