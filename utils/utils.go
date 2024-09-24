package utils

func CreateLinkedList(slice []int) *ListNode {
    head := new(ListNode)
    tmp := head

    for _, v := range slice {
        tmp.Next = &ListNode{Val: v}
        tmp = tmp.Next
    }

    return head.Next
}
