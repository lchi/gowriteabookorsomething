package main

type CharNode struct {
  next *CharNode
  prev *CharNode
  ch byte
}

func MakeNode(b byte) (newNode *CharNode) {
  newNode = new(CharNode)
  newNode.ch = b
  return
}

func (charNode *CharNode) Remove() (*CharNode){
  charNode.prev.next = charNode.next
  charNode.next.prev = charNode.prev
  return charNode.prev
}

func (charNode *CharNode) InsertAfter(b byte) (newNode *CharNode){
  newNode = MakeNode(b)
  newNode.prev, newNode.next = charNode, nil

  if charNode.next == nil {
    charNode.next = newNode
  } else {
    charNode.next.prev = newNode
    newNode.next = charNode.next
    charNode.next = newNode
  }
  return
}

func (charNode *CharNode) InsertBefore(b byte) (newNode *CharNode){
  newNode = MakeNode(b)
  newNode.prev, newNode.next = nil, charNode

  if charNode.prev == nil {
    charNode.prev = newNode
  } else {
    charNode.prev.next = newNode
    newNode.prev = charNode.prev
    charNode.prev = newNode
  }
  return
}
