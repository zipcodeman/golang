package main

import "fmt"

type SkipList struct {
  next *SkipList
  down *SkipList
  value *int
}

func (sl *SkipList) find (i int) (*SkipList, *SkipList, []*SkipList) {
  var previous []*SkipList
  return sl.findWithPrevious(i, previous)
}

func (sl *SkipList) findWithPrevious (i int, previous []*SkipList) (*SkipList, *SkipList, []*SkipList) {
  current, previous := sl.findInRow(i)
  if current != nil || previous.down == nil {
    return current, previous, previous
  } else {
    return previous.down.find(i)
  }
}

func (sl *SkipList) add (i int) (*SkipList) {
  _, previous := sl.find(i)

  newNode := previous.addToRow(i)

  return newNode
}

func (sl *SkipList) findInRow (i int) (*SkipList, *SkipList) {
  var previous *SkipList
  current := sl

  for current != nil && current.value != nil && *current.value < i {
    previous = current
    current = current.next
  }

  if current != nil && *current.value != i {
    current = nil
  }

  return current, previous
}

func (sl *SkipList) addToRow (i int) (*SkipList) {
  if (sl.value != nil) {
    newNode := new(SkipList)
    newNode.value = &i

    _, loc := sl.findInRow(i)

    if loc != nil {
      newNode.next = loc.next
      loc.next = newNode
    } else {
      newNode.value = sl.value
      newNode.next = sl.next
      sl.next = newNode
      sl.value = &i
      newNode = sl
    }
    return newNode
  } else {
    sl.value = &i
    return sl
  }
}

func (sl SkipList) printNode() {
  if (sl.value != nil) {
    fmt.Printf(" --> %d", *sl.value)
    if (sl.next != nil) {
      sl.next.printNode()
    }
  }
}

func (sl SkipList) printList() {
  if (sl.value != nil) {
    fmt.Printf("%d", *sl.value)
    if (sl.next != nil) {
      sl.next.printNode()
    }
  }
  fmt.Printf("\n")
}

func main() {
  sl := new(SkipList)
  sl.addToRow(2)
  sl.printList()
  sl.addToRow(4)
  sl.printList()
  sl.addToRow(1)
  sl.printList()
  sl.addToRow(5)
  sl.printList()
  sl.addToRow(3)
  sl.printList()

  a, b := sl.find(4)

  fmt.Printf("Head %d, Current %d, Previous %d\n", *sl.value, *a.value, *b.value)
}
