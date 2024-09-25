package easy

import (
    "fmt"
)

func AddBinaryCase() {
    // s1 := "1010"
    // s2 := "1011"

    // s1 := "11"
    // s2 := "1"

    s1 := "100"
    s2 := "110010"

    result := addBinaryS1(s1, s2)
    fmt.Println(result)
}

// addBinaryS1 比较麻烦的思路；从较短的那个字符串的末尾开始相加
// 把字符转为整型然后判断是否有进位
// 最后再把较长的字符串的剩余位数依次加入到结果字符串中
func addBinaryS1(a, b string) string {
    var ans string
    lenA := len(a)
    lenB := len(b)
    carry := 0
    base := 48 // ascii 48 == '0'
    radix := 2

    var shorterLen int
    var longerLen int
    if lenA < lenB {
        shorterLen = lenA
        longerLen = lenB
    } else {
        shorterLen = lenB
        longerLen = lenA
    }

    for i := shorterLen - 1; i >= 0; i-- {
        var chANum int
        var chBNum int
        if lenA < lenB {
            chANum = int(a[i]) - base
            chBNum = int(b[longerLen-shorterLen+i]) - base
        } else {
            chANum = int(a[longerLen-shorterLen+i]) - base
            chBNum = int(b[i]) - base
        }

        sum := carry + chANum + chBNum
        if sum >= radix {
            carry = 1
        } else {
            carry = 0
        }
        ans = string(rune((sum%radix)+base)) + ans
    }

    if lenA == lenB && carry == 1 {
        ans = "1" + ans
    } else if lenA != lenB {
        var aOrBStr string
        if longerLen == lenA {
            aOrBStr = a
        } else {
            aOrBStr = b
        }

        for i := longerLen - shorterLen - 1; i >= 0; i-- {
            chNum := int(aOrBStr[i]) - base
            sum := carry + chNum
            if sum >= radix {
                carry = 1
            } else {
                carry = 0
            }
            ans = string(rune((sum%radix)+base)) + ans
        }

        if carry == 1 {
            ans = "1" + ans
        }
    }

    return ans
}
