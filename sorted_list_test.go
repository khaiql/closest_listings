package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertToSortedList_WithoutCap(t *testing.T) {
	l := new(SortedList)
	l.Insert(10, 10)
	l.Insert(9, 9)
	l.Insert(15, 15)
	l.Insert(11, 11)
	l.Insert(20, 20)

	assert.Equal(t, []interface{}{9, 10, 11, 15, 20}, l.Top(-1))
	assert.Equal(t, int64(5), l.Len())
}

func TestInsertSortedList_WithCap(t *testing.T) {
	l := NewSortedList(3)
	l.Insert(10, 10)
	l.Insert(9, 9)
	l.Insert(20, 20)
	l.Insert(3, 3)
	l.Insert(11, 11)
	l.Insert(90, 90)

	assert.Equal(t, []interface{}{3, 9, 10}, l.Top(-1))
	assert.Equal(t, int64(3), l.Len())
}

func TestTop(t *testing.T) {
	l := new(SortedList)
	l.Insert(10, 10)
	l.Insert(9, 9)
	l.Insert(20, 20)

	t.Run("n < 0", func(t *testing.T) {
		results := l.Top(-1)
		assert.Equal(t, []interface{}{9, 10, 20}, results)

		resultsWithScore := l.TopWithScore(-1)
		expectedResult := []*Element{
			{Score: 9, Value: 9},
			{Score: 10, Value: 10},
			{Score: 20, Value: 20},
		}
		assert.Equal(t, expectedResult, resultsWithScore)
	})

	t.Run("0 < n < l.Len()", func(t *testing.T) {
		results := l.Top(2)
		assert.Equal(t, []interface{}{9, 10}, results)

		resultsWithScore := l.TopWithScore(2)
		expectedResult := []*Element{
			{Score: 9, Value: 9},
			{Score: 10, Value: 10},
		}
		assert.Equal(t, expectedResult, resultsWithScore)
	})

	t.Run("n > l.Len()", func(t *testing.T) {
		results := l.Top(5)
		assert.Equal(t, []interface{}{9, 10, 20}, results)

		resultsWithScore := l.TopWithScore(5)
		expectedResult := []*Element{
			{Score: 9, Value: 9},
			{Score: 10, Value: 10},
			{Score: 20, Value: 20},
		}
		assert.Equal(t, expectedResult, resultsWithScore)
	})
}
