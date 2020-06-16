package pro297

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	nodeVals := make([]string, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur == nil {
			nodeVals = append(nodeVals, "null")
		} else {
			nodeVals = append(nodeVals, strconv.Itoa(cur.Val))
		}

		if cur != nil {
			q = append(q, []*TreeNode{cur.Left, cur.Right}...)
		}
	}

	tail := len(nodeVals)
	// 去掉尾部多余的null
	for {
		tail--
		if nodeVals[tail] == "null" {
			nodeVals = nodeVals[0:tail]
		} else {
			break
		}
	}

	return "[" + strings.Join(nodeVals, ",") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = strings.Replace(data, "[", "", -1)
	data = strings.Replace(data, "]", "", -1)
	if data == "" {
		return nil
	}
	nodeVals := strings.Split(data, ",")
	if len(nodeVals) == 0 {
		return nil
	}

	root := NewNode(nodeVals[0])
	q := make([]*TreeNode, 1)
	q[0] = root
	nodeVals = nodeVals[1:]
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur == nil {
			continue
		}

		if len(nodeVals) > 0 {
			cur.Left = NewNode(nodeVals[0])
			nodeVals = nodeVals[1:]
		}
		if len(nodeVals) > 0 {
			cur.Right = NewNode(nodeVals[0])
			nodeVals = nodeVals[1:]
		}

		q = append(q, []*TreeNode{cur.Left, cur.Right}...)
	}

	return root
}

func NewNode(val string) *TreeNode {
	if val == "null" {
		return nil
	}

	intVal, _ := strconv.ParseInt(val, 10, 32)
	return &TreeNode{
		Val: int(intVal),
	}
}
