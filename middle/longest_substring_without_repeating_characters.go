package middle

import "fmt"

func LongestSubstringWithoutRepeatingCharactersCase() {
    l := longestSubstringWithoutRepeatingCharactersS1("abcabcbb")
    fmt.Println(l)
}

func longestSubstringWithoutRepeatingCharactersS1(s string) int {
    sLength := len(s)
    if sLength <= 1 {
        return sLength
    }

    maxLength := 1
    tempLength := 1
    for i := 1; i < sLength; i++ {
        curr := s[i]
        prev := s[i-1]

        if curr != prev {
            tempLength++
            if tempLength > maxLength {
                maxLength = tempLength
            }
        } else {
            tempLength = 1
        }
    }

    return maxLength
}
