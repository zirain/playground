package pro297

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	str := "[]"
	codec := Constructor()
	desRes := codec.deserialize(str)
	if desRes != nil {
		t.Fail()
	}

	serRes := codec.serialize(desRes)
	if serRes != str {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	str := "[1,2,3,null,null,4,5]"
	codec := Constructor()
	desRes := codec.deserialize(str)
	if desRes == nil && desRes.Val == 1 {
		t.Fail()
	}

	serRes := codec.serialize(desRes)
	if serRes != str {
		fmt.Println(serRes)
		t.Fail()
	}
}

func Test3(t *testing.T) {
	str := "[1,null,2]"
	codec := Constructor()
	desRes := codec.deserialize(str)
	if desRes == nil && desRes.Val == 1 {
		t.Fail()
	}

	serRes := codec.serialize(desRes)
	if serRes != str {
		fmt.Println(serRes)
		t.Fail()
	}
}
