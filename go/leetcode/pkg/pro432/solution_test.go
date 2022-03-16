package pro432

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestAllOne(t *testing.T) {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("hello")
	assert.Equal(t, "hello", allOne.GetMaxKey())
	assert.Equal(t, "hello", allOne.GetMinKey())
	allOne.Inc("leet")
	assert.Equal(t, "hello", allOne.GetMaxKey())
	assert.Equal(t, "leet", allOne.GetMinKey())
	allOne.Dec("hello")
	allOne.Dec("hello")
	assert.Equal(t, "leet", allOne.GetMaxKey())
	assert.Equal(t, "leet", allOne.GetMinKey())
}

func TestAllOne2(t *testing.T) {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("world")
	allOne.Inc("leet")
	allOne.Inc("code")
	allOne.Inc("ds")
	allOne.Inc("leet")
	assert.Equal(t, "leet", allOne.GetMaxKey())
	allOne.Inc("ds")
	allOne.Dec("leet")
	assert.Equal(t, "ds", allOne.GetMaxKey())
}

// ["AllOne","inc","inc","inc","inc","inc","inc","inc","inc","inc","inc","inc","inc","getMinKey"]
// [[],["a"],["b"],["c"],["d"],["a"],["b"],["c"],["d"],["c"],["d"],["d"],["a"],[]]

func TestAllOne3(t *testing.T) {
	allOne := Constructor()
	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("c")
	allOne.Inc("d")

	allOne.Inc("a")
	allOne.Inc("b")
	allOne.Inc("c")
	allOne.Inc("d")

	allOne.Inc("c")
	allOne.Inc("d")
	allOne.Inc("d")
	allOne.Inc("a")

	assert.Equal(t, "b", allOne.GetMinKey())
}
