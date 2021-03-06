package longestsubstringwithoutrepeatingchars

import (
	"testing"
)

func TestGetLongestSubstrLengthWithoutRepeats_NormalEarly(t *testing.T) {
	result := getLongestSubstrLengthWithoutRepeats("abcabcbb")
	expected := 3 // abc

	if result != expected {
		t.Errorf("Length is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestGetLongestSubstrLengthWithoutRepeats_AllRepeats(t *testing.T) {
	result := getLongestSubstrLengthWithoutRepeats("bbbbb")
	expected := 1 // b

	if result != expected {
		t.Errorf("Length is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestGetLongestSubstrLengthWithoutRepeats_NotSubsequence(t *testing.T) {
	result := getLongestSubstrLengthWithoutRepeats("pwwkew")
	expected := 3 // The answer is "wke", with the length of 3. 
	// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

	if result != expected {
		t.Errorf("Length is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestGetLongestSubstrLengthWithoutRepeats_OneChar(t *testing.T) {
	result := getLongestSubstrLengthWithoutRepeats("z")
	expected := 1 // The answer is "z", with the length of 1. 

	if result != expected {
		t.Errorf("Length is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestIsStringUniqueChars_OneChar(t *testing.T) {
	result := isStringUniqueChars("z")
	expected := true

	if result != expected {
		t.Errorf("Is unique string with one char is incorrect, got %v, want %v", result, expected)
	}
}

func TestIsStringUniqueChars_LotsOfCharAllUnique(t *testing.T) {
	result := isStringUniqueChars("zxcvbnmasdfghjklqwertyuiop")
	expected := true

	if result != expected {
		t.Errorf("Is unique string with one char is incorrect, got %v, want %v", result, expected)
	}
}

func TestIsStringUniqueChars_LotsOfCharWithDuplicate(t *testing.T) {
	result := isStringUniqueChars("zxcvbnmasdfghjklqffdafadsadwertyuiop")
	expected := false

	if result != expected {
		t.Errorf("Is unique string with one char is incorrect, got %v, want %v", result, expected)
	}
}