package word

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool{

	fmt.Println(unicode.IsLetter('L'),unicode.ToLower('你'))

	var letters []rune

	for _,r:=range s {
		if unicode.IsLetter(r) {
			letters = append(letters,unicode.ToLower(r))
		}
	}

	for i:=range letters {
		if letters[i]!=letters[len(letters)-i-1]{
			return false
		}
	}
	return true
}


