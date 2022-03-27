package pro2043

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestBank(t *testing.T) {
	bank := Constructor([]int64{0})
	b := &bank
	got := b.Deposit(1, 2)
	assert.Equal(t, true, got)

	got = b.Transfer(1, 1, 1)
	assert.Equal(t, true, got)

	got = b.Transfer(1, 1, 3)
	assert.Equal(t, false, got)
}
