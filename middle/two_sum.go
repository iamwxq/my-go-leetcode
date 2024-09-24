package middle

import (
    "fmt"

    "golc/utils"
)

func TwoSumCase() {
    // s1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
    s1 := []int{2, 4, 3}
    s2 := []int{5, 6, 4}
    l1 := utils.CreateLinkedList(s1)
    l2 := utils.CreateLinkedList(s2)

    // addTwoNumbersS1(l1, l2)
    ans := addTowNumbersS2(l1, l2)
    for ans != nil {
        fmt.Printf("%d ", ans.Val)
        ans = ans.Next
    }
}

// 自己想出来的解决方案 1566/1569 没ac
// l1 = [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
// l2 = [5,6,4]
// 原因是 base 会溢出
func addTwoNumbersS1(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
    ans := new(utils.ListNode)

    sum := 0
    base := 1
    p := l1
    q := l2

    for p != nil || q != nil {
        if p != nil {
            sum += p.Val * base
            p = p.Next
        }

        if q != nil {
            sum += q.Val * base
            q = q.Next
        }

        base *= 10
    }

    temp := ans
    temp.Val = sum % 10
    sum /= 10

    for sum != 0 {
        r := new(utils.ListNode)
        r.Val = sum % 10
        r.Next = nil
        temp.Next = r
        temp = r
        sum /= 10
    }

    return ans
}

// 参考了答案后写的
func addTowNumbersS2(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
    ans := new(utils.ListNode)
    curr := ans
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        v1 := 0
        if l1 != nil {
            v1 = l1.Val
        }

        v2 := 0
        if l2 != nil {
            v2 = l2.Val
        }

        sum := carry + v1 + v2
        if sum > 9 {
            carry = 1
        } else {
            carry = 0
        }
        curr.Val = sum % 10

        if l1 != nil {
            l1 = l1.Next
        }

        if l2 != nil {
            l2 = l2.Next
        }

        if l1 == nil && l2 == nil && carry == 0 {
            break
        }

        curr.Next = new(utils.ListNode)
        curr = curr.Next
    }

    return ans
}
