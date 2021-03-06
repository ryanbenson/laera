package isanagram

import (
	"testing"
)

func TestIsAnagram_Success(t *testing.T) {
	result1 := isAnagram("fried", "fired")
	result2 := isAnagram("listen", "silent")
	expected := true

	if result1 != expected {
		t.Errorf("When determing anagram 1, incorrect result given, got: %v, expected: %v", result1, expected)
	}

	if result2 != expected {
		t.Errorf("When determing anagram 2, incorrect result given, got: %v, expected: %v", result2, expected)
	}
}

func TestIsAnagram_Fail(t *testing.T) {
	result1 := isAnagram("bob", "boo")
	result2 := isAnagram("hello", "help")
	expected := false

	if result1 != expected {
		t.Errorf("When determing anagram 1, incorrect result given, got: %v, expected: %v", result1, expected)
	}

	if result2 != expected {
		t.Errorf("When determing anagram 2, incorrect result given, got: %v, expected: %v", result2, expected)
	}
}

func TestIsAnagram_Empty(t *testing.T) {
	result1 := isAnagram("", "")
	expected := true

	if result1 != expected {
		t.Errorf("When determing anagram, and both strings are empty, incorrect result given, got: %v, expected: %v", result1, expected)
	}
}

func TestIsAnagram_EmptyFirst(t *testing.T) {
	result1 := isAnagram("", "hello")
	expected := false

	if result1 != expected {
		t.Errorf("When determing anagram, and first string is empty, incorrect result given, got: %v, expected: %v", result1, expected)
	}
}

func TestIsAnagramStringsManip_Success(t *testing.T) {
	result1 := isAnagramStringsManip("fried", "fired")
	result2 := isAnagramStringsManip("listen", "silent")
	expected := true

	if result1 != expected {
		t.Errorf("When determing anagram 1, incorrect result given, got: %v, expected: %v", result1, expected)
	}

	if result2 != expected {
		t.Errorf("When determing anagram 2, incorrect result given, got: %v, expected: %v", result2, expected)
	}
}

func TestIsAnagramStringsManip_Fail(t *testing.T) {
	result1 := isAnagramStringsManip("bob", "boo")
	result2 := isAnagramStringsManip("hello", "help")
	expected := false

	if result1 != expected {
		t.Errorf("When determing anagram 1, incorrect result given, got: %v, expected: %v", result1, expected)
	}

	if result2 != expected {
		t.Errorf("When determing anagram 2, incorrect result given, got: %v, expected: %v", result2, expected)
	}
}

func TestIsAnagramStringsManip_Empty(t *testing.T) {
	result1 := isAnagramStringsManip("", "")
	expected := true

	if result1 != expected {
		t.Errorf("When determing anagram, and both strings are empty, incorrect result given, got: %v, expected: %v", result1, expected)
	}
}

func TestIsAnagramStringsManip_EmptyFirst(t *testing.T) {
	result1 := isAnagramStringsManip("", "hello")
	expected := false

	if result1 != expected {
		t.Errorf("When determing anagram, and first string is empty, incorrect result given, got: %v, expected: %v", result1, expected)
	}
}

func TestSortString(t *testing.T) {
	result := sortString("hello")
	expected := "ehllo"

	if result != expected {
		t.Errorf("When sorting string, incorrect result given, got: %v, expected: %v", result, expected)
	}
}

func BenchmarkIsAnagram_SortStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram("listen", "silent")
	}
}

func BenchmarkIsAnagramStringsManip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagramStringsManip("listen", "silent")
	}
}
