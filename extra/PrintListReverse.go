package main

import (
    "container/list"
)

//从尾到头打印链表

//使用栈+循环  鲁棒性比起递归更好
func PrintListReverseWithStack(head *ListNode) {
    phead := head
    stack := list.New()
    for phead != nil {
        stack.PushFront(phead.value)
    }
    for stack.Len() > 0 {
        value := stack.Front().Value.(int)
        println(value)
        stack.Remove(stack.Front())
    }
}

//既然使用栈，就想到递归
func PrintListReverseWithRecurse(head *ListNode) {
    if head != nil {
        if head.next != nil {
            PrintListReverseWithRecurse(head.next)
        }
        println(head.value)
    }
}


