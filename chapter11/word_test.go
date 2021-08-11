package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated")==false`)
	}
	if !IsPalindrome("kafyak") {
		t.Fatal(`IsPalindrome("kayak")==false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`) }
}

func TestCanalPalindrome(t * testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

func TestFrenchPalindrome(t *testing.T) { if !IsPalindrome("été") {
	t.Error(`IsPalindrome("été") = false`) }
}

func TestIsPalindrome(t *testing.T){
	var tests = []struct{
		input string
		want bool
	}{
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true}, {"Evil I did dwell; lewd did I live.", true}, {"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true}, {"palindrome", false}, // non-palindrome {"desserts", false}, // semi-palindrome
	}
	
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want { t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}