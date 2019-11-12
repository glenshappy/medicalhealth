package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//生成随机的字符串
func randomPalindrome(rng *rand.Rand)  string{
	n:=rand.Intn(26)
	fmt.Println("随机数值：",n)
	runes:=make([]rune,n)
	for i:=0;i<(n+1)/2;i++ {
		r:=rune(rng.Intn(0x1000)) //最大为\u999
		runes[i] = r
		runes[n-i-1] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	fmt.Println("max num:",0x1000)
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	//t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 10; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

//func TestIsPalindrome(t *testing.T)  {
//	var tests = []struct{
//		input string
//		want bool
//	}{
//		{"", true},
//		{"a", true},
//		{"aa", true},
//		{"abc1", true},
//		{"kayak", true},
//		{"detartrated", true},
//		{"A man, a plan, a canal: Panama", true},
//		{"Evil I did dwell; lewd did I live.", true},
//		{"Able was I ere I saw Elba", true},
//		{"été", true},
//		{"Et se resservir, ivresse reste.", true},
//		{"palindrome", false}, // non-palindrome
//		{"desserts", false}, // semi-palindrome
//	}
//	for _,w:=range tests {
//		rs:=IsPalindrome(w.input)
//		if rs!=w.want {
//			t.Errorf("IsPalindrome(%q) = %v",w.input,rs)
//		}
//	}
//}