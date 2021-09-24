package method_expressions

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

type Data struct {
    val int
}

func (d *Data)  Add(val int) {
    d.val += val
}

func (d *Data) Multiply(val int) {
    d.val *= val
}

func TestMethodExpressions(t *testing.T) {

    data := Data{ val: 1 }
    for _, bleh := range [...]func(*Data, int){ (*Data).Add, (*Data).Multiply } {
        bleh(&data, 2)
    }
    assert.Equal(t, 6, data.val)

    addFunc := data.Add
    addFunc(10)
    assert.Equal(t, 16, data.val)
}

