A numeronym is a word where numbers are used to form an abbreviation. One way to do this is by replacing the letters between the first and the last letter with a number representing the number of letters omittedd, such as in the numeronym i18n for internationalization.

Write a function that takes in a word and a numeronym and returns a boolean value true if the numeronym is a valid abbreviation for the word given, and false if it is not.

Example: `isValid("internationalization", "i18n")` and `isValid("internationalization", "i2e15n")` should return true, while `isValid("internationalization", "i4n15n")` should return false.