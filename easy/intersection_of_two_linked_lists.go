package easy

import (
    "fmt"

    "golc/utils"
)

func IntersectionOfTwoLinkedListsCase() {
    // l1 := []int{4, 1, 8, 4, 5}
    // l2 := []int{5, 6, 1, 8, 4, 5}
    // intersection := 8
    // skipA := 2
    // skipB := 3

    // l1 := []int{2, 6, 4}
    // l2 := []int{1, 5}
    // intersection := 0
    // skipA := 3
    // skipB := 2

    l1 := []int{1, 9, 1, 2, 4}
    l2 := []int{3, 2, 4}
    intersection := 2
    skipA := 3
    skipB := 1

    headA, headB := createIntersectionLinkedList(l1, l2, intersection, skipA, skipB)

    target := intersectionOfTwoLinkedListsS1(headA, headB)
    if target != nil {
        fmt.Printf("\nIntersection of two linked lists is %v\n", target.Val)
    } else {
        fmt.Printf("\nIntersection of two linked lists is nil\n")
    }
}

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

func createIntersectionLinkedList(
    a, b []int,
    intersection, skipA, skipB int) (*utils.ListNode, *utils.ListNode) {
    var headA, headB = new(utils.ListNode), new(utils.ListNode)
    var curA, curB int

    tmpA := headA
    tmpB := headB

    for curA < skipA || curA < len(a) {
        if a[curA] == intersection && curA == skipA {
            break
        }

        tmpA.Next = &utils.ListNode{Val: a[curA]}
        tmpA = tmpA.Next
        curA += 1
    }

    for curB < skipB || curB < len(b) {
        if b[curB] == intersection && curB == skipB {
            break
        }

        tmpB.Next = &utils.ListNode{Val: b[curB]}
        tmpB = tmpB.Next
        curB += 1
    }

    if skipA >= len(a) || skipB >= len(b) {
        return headA.Next, headB.Next
    }

    var intersectNode *utils.ListNode
    hasIntersection := a[skipA] == intersection && b[skipB] == intersection
    if curA == skipA && curB == skipB && hasIntersection {
        intersectNode = new(utils.ListNode)
        intersectNode.Val = intersection
        tmpA.Next = intersectNode
        tmpB.Next = intersectNode
    }

    curI := curA + 1
    tmpI := intersectNode
    if tmpI != nil {
        for curI < len(a) {
            val := a[curI]
            tmpI.Next = &utils.ListNode{Val: val}
            curI++
            tmpI = tmpI.Next
        }
    }

    return headA.Next, headB.Next
}

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
