package main

type node struct {
	score float64
	value interface{}
	next  *node
	prev  *node
}

// Element represents an element in the list, with its score and value
type Element struct {
	Score float64
	Value interface{}
}

// SortedList is a data structure that care about orders of elements upon insertion
// SortedList that has capacity always maintains max(number of elements) = cap
type SortedList struct {
	head *node
	tail *node
	len  int64
	cap  int64
}

// NewSortedList inits a SortedList that has capacity (cap = 0 means no cap)
func NewSortedList(cap int64) *SortedList {
	return &SortedList{cap: cap}
}

// Top returns top n elements of the list
// Use n < 0 if you want to get the whole list
func (s *SortedList) Top(n int64) []interface{} {
	elements := s.TopWithScore(n)
	results := make([]interface{}, len(elements))
	for i, el := range elements {
		results[i] = el.Value
	}

	return results
}

// TopWithScore is similar to Top but also returns score of the respective element
func (s *SortedList) TopWithScore(n int64) []*Element {
	if n < 0 || n > s.Len() {
		n = s.len
	}

	results := make([]*Element, n)
	var count int64
	node := s.head
	for count < n && node != nil {
		results[count] = &Element{Value: node.value, Score: node.score}
		node = node.next
		count += 1
	}

	return results
}

// Len returns number of elements in the list
func (s *SortedList) Len() int64 {
	return s.len
}

func (s *SortedList) checkCap() {
	// cap = 0 means no cap
	if s.cap == 0 {
		return
	}

	if s.len > s.cap {
		// cut off tail node
		currentTail := s.tail
		newTail := currentTail.prev
		newTail.next = nil
		currentTail.prev = nil
		s.tail = newTail
		s.len -= 1
	}
}

// Insert a value with a score to the sorted list. The function preserves sorted order of the list.
// It also maintains the capacity of the list if there is a cap
// Time complexity: O(N), N = len of the list
// Space complexity: O(1)
func (s *SortedList) Insert(score float64, value interface{}) {
	n := &node{score: score, value: value}
	s.len += 1
	if s.head == nil {
		s.head = n
		s.tail = n
		return
	}

	// new head
	if s.head.score >= score {
		oldHead := s.head
		s.head = n
		n.next = oldHead
		oldHead.prev = n
		s.checkCap()
		return
	}

	// new tail
	if s.tail.score <= score {
		oldTail := s.tail
		s.tail = n
		oldTail.next = n
		n.prev = oldTail
		s.checkCap()
		return
	}

	pointer := s.head
	for pointer != nil && pointer.score < score {
		pointer = pointer.next
	}
	pointer.prev.next = n
	n.prev = pointer.prev
	n.next = pointer
	pointer.prev = n

	s.checkCap()
}
