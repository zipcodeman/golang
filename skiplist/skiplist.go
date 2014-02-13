package main

import "fmt"
import "math/rand"

type SkipList struct {
  next *SkipList
  down *SkipList
  value *int
}

func (sl *SkipList) height () int {
  if (sl == nil) {
    return 0
  } else {
    return 1 + sl.down.height()
  }
}

func (sl *SkipList) find (i int) (*SkipList, *SkipList, []*SkipList) {
  previous := make([]*SkipList, sl.height())
  return sl.findWithPrevious(i, previous, 0)
}

func (sl *SkipList) findWithPrevious (i int, previous []*SkipList, index int) (*SkipList, *SkipList, []*SkipList) {
  current, prev := sl.findInRow(i)
  if current != nil || prev.down == nil {
    return current, prev, previous
  } else {
    previous[index] = prev
    return prev.down.findWithPrevious(i, previous, index + 1)
  }
}

func (sl *SkipList) add (i int) (*SkipList) {
  _, prev, _ := sl.find(i)

  newNode := prev.addToRow(i)

  for rand.Intn(3) == 0 {
    fmt.Printf("%d would be promoted\n", i)
  }

  return newNode
}

func (sl *SkipList) findInRow (i int) (*SkipList, *SkipList) {
  var previous *SkipList
  current := sl

  for current != nil && current.value != nil && *current.value < i {
    previous = current
    current = current.next
  }

  if current != nil && current.value != nil && *current.value != i {
    current = nil
  }

  return current, previous
}

func (sl *SkipList) addToRow (i int) (*SkipList) {
  if sl != nil && sl.value != nil {
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
  } else if sl != nil {
    sl.value = &i
    return sl
  }
  return nil
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
  sl.add(2)
  sl.printList()
  sl.add(4)
  sl.printList()
  sl.add(1)
  sl.printList()
  sl.add(5)
  sl.printList()
  sl.add(3)
  sl.printList()

  a, b, _ := sl.find(4)

  fmt.Printf("Head %d\n", *sl.value)
  fmt.Printf("Current %d\n", *a.value)
  fmt.Printf("Previous %d\n", *b.value)
}
