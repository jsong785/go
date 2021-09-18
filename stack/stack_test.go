package stack

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type internal_items []interface{}

func (in internal_items) Len() int {
	return len(in)
}

func (in internal_items) Less(i, j int) bool {
	return i < j // abuse of sort reverse
}

func (in internal_items) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

type Stack struct {
	items internal_items
}

func (s *Stack) Push(items ...interface{}) {
	if s.items == nil {
		s.items = make(internal_items, 0, 10)
	}
	s.items = append(s.items, items...)
}

func (s *Stack) Pop() (interface{}, error) {
	if ret, err := s.PopNum(1); err != nil {
		return nil, err
	} else {
		return ret[0], nil
	}
}

func (s *Stack) PopNum(amt int) ([]interface{}, error) {
	if len(s.items) < amt {
		return nil, errors.New("error")
	}
	ret := s.items[len(s.items)-amt:]
	sort.Sort(sort.Reverse(internal_items(ret)))

	s.items = s.items[:len(s.items)-len(ret)]
	return ret, nil
}

func (s Stack) Size() int {
	return len(s.items)
}

func TestStack(t *testing.T) {
	t.Run("pop one", func(t *testing.T) {
		s := Stack{}
		assert.Equal(t, s.Size(), 0)

		s.Push(1, 2, 3, 4)
		assert.Equal(t, s.Size(), 4)

		s.Push(5)
		assert.Equal(t, s.Size(), 5)

		i, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, i, 5)
		assert.Equal(t, s.Size(), 4)

		i, err = s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, i, 4)
		assert.Equal(t, s.Size(), 3)

		i, err = s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, i, 3)
		assert.Equal(t, s.Size(), 2)

		i, err = s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, i, 2)
		assert.Equal(t, s.Size(), 1)

		i, err = s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, i, 1)
		assert.Equal(t, s.Size(), 0)

		i, err = s.Pop()
		assert.NotNil(t, err)
		assert.Equal(t, s.Size(), 0)
	})

	t.Run("pop multiple", func(t *testing.T) {
		s := Stack{}
		assert.Equal(t, s.Size(), 0)

		s.Push("abc", "def", "ghi", "jkl", "mno")
		assert.Equal(t, s.Size(), 5)

		ret, err := s.PopNum(3)
		assert.Nil(t, err)
		assert.Equal(t, ret, []interface{}{"mno", "jkl", "ghi"})
		assert.Equal(t, s.Size(), 2)

		ret, err = s.PopNum(3)
		assert.NotNil(t, err)

		ret, err = s.PopNum(2)
		assert.Nil(t, err)
		assert.Equal(t, ret, []interface{}{"def", "abc"})
		assert.Equal(t, s.Size(), 0)
	})
}
