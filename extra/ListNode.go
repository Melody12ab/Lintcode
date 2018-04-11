package main

import "fmt"

type ListNode struct {
    value int
    next  *ListNode
}

//链表末尾添加值
func AddToTail(phead *ListNode, value int) {
    pnew := &ListNode{}
    pnew.value = value
    pnew.next = nil
    //空链表时
    if phead == nil {
        phead = pnew
    } else {
        pnode := phead
        //找到链表的最后一个节点
        for pnode.next != nil {
            pnode = pnode.next
        }
        pnode.next = pnew
    }
}

//删除第一个找到的值
func RemoveNode(phead *ListNode, value int) {
    if phead == nil {
        return
    }
    if phead.value == value {
        phead = phead.next
    } else {
        pnode := phead
        for pnode.next != nil && pnode.next.value != value {
            pnode = pnode.next
        }
        if pnode.next != nil && pnode.next.value == value {
            pnode.next = pnode.next.next
        }
    }
}

func PrintNode(head *ListNode) {
    pnode := head
    for pnode != nil {
        fmt.Println(pnode.value)
        pnode = pnode.next
    }
}

func main() {
    head := &ListNode{10, nil}
    AddToTail(head, 120)
    AddToTail(head, 14)
    AddToTail(head, 27)
    PrintNode(head)
    RemoveNode(head, 14)
    fmt.Println("===============")
    PrintNode(head)
}
