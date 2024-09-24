package easy

import "golc/utils"

func IntersectionOfTwoLinkedListsCase() {
    l1 := []int{4, 1, 8, 4, 5}
    l2 := []int{5, 6, 1, 8, 4, 5}
    intersection := 8
    skipA := 2
    skipB := 3
}

// intersectionOfTwoLinkedListsS1 暴力双重循环求解
func intersectionOfTwoLinkedListsS1(headA, headB *utils.ListNode) *utils.ListNode {
    var ans *utils.ListNode

    if headA == headB {
        return headA
    }

    t := shorterList(headA, headB)
    var k *utils.ListNode

    if t == headA {
        k = headB
    } else {
        k = headA
    }

    memo := t
FindAnswer:
    for k != nil {
        t = memo
        for t != nil {
            if t == k {
                ans = t
                break FindAnswer
            }
            t = t.Next
        }
        k = k.Next
    }

    return ans
}

// shorterList 返回更短的那一条链表
func shorterList(a, b *utils.ListNode) *utils.ListNode {
    lenA := 0
    lenB := 0

    p := a
    q := b

    for p != nil {
        lenA++
        p = p.Next
    }

    for q != nil {
        lenB++
        q = q.Next
    }

    if lenA >= lenB {
        return b
    } else {
        return a
    }
}
