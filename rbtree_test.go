package rbtree

import (
	"testing"
)

type node struct {
	l, r  *node
	color bool
	key   int
}

type tree struct {
	head *node
}

func (t *tree) GetL(n Node) Node {
	return n.(*node).l
}
func (t *tree) SetL(n Node, l Node) {
	n.(*node).l = l.(*node)
}
func (t *tree) GetR(n Node) Node {
	return n.(*node).r
}
func (t *tree) SetR(n Node, r Node) {
	n.(*node).r = r.(*node)
}
func (t *tree) GetColor(n Node) NodeColor {
	return NodeColor(n.(*node).color)
}
func (t *tree) SetColor(n Node, color NodeColor) {
	n.(*node).color = bool(color)
}
func (t *tree) GetKey(n Node) Key {
	return n.(*node).key
}

func (t *tree) LessKey(i, j Key) bool {
	return i.(int) < j.(int)
}
func (t *tree) EqKey(i, j Key) bool {
	return i.(int) == j.(int)
}
func (t *tree) EqNode(i, j Node) bool {
	return i.(*node) == j.(*node)
}
func (t *tree) SetHead(h Node) {
	t.head = h.(*node)
}
func (t *tree) GetHead() Node {
	return t.head
}
func (t *tree) NewNode(key Key) Node {
	return &node{key: key.(int)}
}
func (t *tree) DeleteNode(n Node) {

}
func (t *tree) IsNull(h Node) bool {
	return h.(*node) == nil
}

type ttestInsert struct {
	InsertKey int
	Min, Max  int
}

var insertKeys = []ttestInsert{
	{100, 100, 100},
	{200, 100, 200},
	{110, 100, 200},
	{10, 10, 200},
	{50, 10, 200},
	{60, 10, 200},
	{500, 10, 500},
	{5, 5, 500},
}

func TestInsert(t *testing.T) {
	{
		tree := &tree{}
		var i int
		for i = 0; i < 1000; i++ {
			Insert(tree, i)
			if _, err := CheckCaseAll(tree); err != nil {
				t.Error(err)
			}
			if Search(tree, i) == nil {
				t.Error("insert ", i, "not Search")
			}
		}
		for i = 0; i < 1000; i++ {
			if Search(tree, i) == nil {
				t.Error("insert ", i, "not Search")
			}
		}
	}
	{
		tree := &tree{}
		var i int
		for _, p := range insertKeys {
			Insert(tree, p.InsertKey)
			if _, err := CheckCaseAll(tree); err != nil {
				t.Error(err)
			}
			if Search(tree, p.InsertKey) == nil {
				t.Error("insert ", i, "not Search")
			}
			if m := Min(tree); m.(int) != p.Min {
				t.Error("insert ", i, "min", m, "!=", p.Min)
			}
			if m := Max(tree); m.(int) != p.Max {
				t.Error("insert ", i, "max", m, "!=", p.Max)
			}

		}
		for _, p := range insertKeys {
			if Search(tree, p.InsertKey) == nil {
				t.Error("insert ", p.InsertKey, "not Search")
			}
		}
	}
}
func TestDelete(t *testing.T) {
	{
		const num = 1000
		var i, j int
		for j = 0; j < num; j++ {
			tree := &tree{}
			defer func() {
				if r := recover(); r != nil {
					t.Error("Delete panic ", j, r)
				}
			}()
			for i = 0; i < num; i++ {
				Insert(tree, i)
				if Search(tree, i) == nil {
					t.Error("insert ", i, "not Search")
				}
			}
			if _, err := CheckCaseAll(tree); err != nil {
				t.Error("tree create. check case ", err)
			}
			if !Delete(tree, j) {
				t.Error("Delete ", j)
			}
			if _, err := CheckCaseAll(tree); err != nil {
				t.Error("delete. check case", err)
			}
			for i = 0; i < num; i++ {
				if i == j {
					if Search(tree, i) != nil {
						t.Error("delete:", j, "contain ", i)
					}
				} else {
					if Search(tree, i) == nil {
						t.Error("delete:", j, "not contain ", i)
					}
				}
			}
			if _, err := CheckCaseAll(tree); err != nil {
				t.Error(err)
			}
		}
	}
}
