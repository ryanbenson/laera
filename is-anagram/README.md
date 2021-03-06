# Is Anagram

Reference: Issue #3 of [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Write a function that takes in two strings and returns true if they are anagrams

Example:

```console
isAnagram("fried", "fired") => true
isAnagram("listen", "silent") => true
isAnagram("bob", "boo") => false
```

## Thoughts

Overall it was pretty simple to determine. Just need to check every letter in each word to ensure every letter matches. A simple, naive, approach is to just sort them both then check them, but that comes at the cost of sorting them, which ends up splitting them into arrays, sorting them, and merging them back. There are more optimized ways of doing this.

Created a method that manipulates strings as it processes. It goes through each letter, finds the index, and if both strings find the letter, then it removes that letter from the string and continues until there are no more letters in the first string. If they are both empty by the end, then it's a match. It performs dramatically better too.

```console
BenchmarkIsAnagramSortStrings-12     	 1910025	       613 ns/op
BenchmarkIsAnagramStringsManip-12    	 7131518	       166 ns/op
```