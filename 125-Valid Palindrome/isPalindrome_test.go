package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	testStr := "A man, aplan, a canal: Panama"
	if !IsPalindrome2(testStr) {
		t.Error("FAIL")
	} else {
		t.Log("PASS")
	}
}

func Test_isPalindrome2(t *testing.T) {

	testStr := "race a car"
	if !IsPalindrome2(testStr) {
		t.Log("PASS")
	} else {
		t.Error("FAIL")
	}
}
