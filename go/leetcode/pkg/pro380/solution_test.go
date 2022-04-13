package pro380

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestRandomizedSet(t *testing.T) {
	s := Constructor()
	set := &s
	assert.Equal(t, true, set.Insert(1))
	assert.Equal(t, false, set.Remove(2))
	assert.Equal(t, true, set.Insert(2))
	r := set.GetRandom()
	if r != 1 && r != 2 {
		t.Fail()
	}

	assert.Equal(t, true, set.Remove(1))
	assert.Equal(t, false, set.Insert(2))
	r = set.GetRandom()
	assert.Equal(t, 2, r)
}

func TestRandomizedSet2(t *testing.T) {
	s := Constructor()
	set := &s
	assert.Equal(t, false, set.Remove(0))
	assert.Equal(t, false, set.Remove(0))
	assert.Equal(t, true, set.Insert(0))
	r := set.GetRandom()
	if r != 0 {
		t.Fail()
	}

	assert.Equal(t, true, set.Remove(0))
	assert.Equal(t, true, set.Insert(0))
}
