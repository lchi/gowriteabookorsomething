package main

type CharNode struct {
  next *CharNode
  prev *CharNode
  ch byte
}

func (charNode *CharNode) Remove() {
  charNode.prev.next = charNode.next
  charNode.next.prev = charNode.prev
}
