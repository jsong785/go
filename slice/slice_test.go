package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceIsJustAView(t *testing.T) {

	data := make([]int, 0, 5)
	data = append(data, 1)
	assert.Equal(t, data, []int{1})

	func(d []int) {
		d = append(d, 2)
		assert.Equal(t, d, []int{1, 2})
	}(data)
	assert.Equal(t, data, []int{1})

	func(d *[]int) {
		*d = append(*d, 2)
		assert.Equal(t, *d, []int{1, 2})
	}(&data)
	assert.Equal(t, data, []int{1, 2})
}
