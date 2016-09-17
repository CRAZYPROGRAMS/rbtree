package main

import (
	"fmt"

	"github.com/crazyprograms/rbtree"
)

type node struct {
	l, r  *node
	color bool
	key   int
}

type tree struct {
	head *node
}

func (t *tree) GetL(n rbtree.Node) rbtree.Node {
	return n.(*node).l
}
func (t *tree) SetL(n rbtree.Node, l rbtree.Node) {
	n.(*node).l = l.(*node)
}
func (t *tree) GetR(n rbtree.Node) rbtree.Node {
	return n.(*node).r
}
func (t *tree) SetR(n rbtree.Node, r rbtree.Node) {
	n.(*node).r = r.(*node)
}
func (t *tree) GetColor(n rbtree.Node) rbtree.NodeColor {
	return rbtree.NodeColor(n.(*node).color)
}
func (t *tree) SetColor(n rbtree.Node, color rbtree.NodeColor) {
	n.(*node).color = bool(color)
}
func (t *tree) GetKey(n rbtree.Node) rbtree.Key {
	return n.(*node).key
}

func (t *tree) LessKey(i, j rbtree.Key) bool {
	return i.(int) < j.(int)
}
func (t *tree) EqKey(i, j rbtree.Key) bool {
	return i.(int) == j.(int)
}
func (t *tree) EqNode(i, j rbtree.Node) bool {
	return i.(*node) == j.(*node)
}
func (t *tree) SetHead(h rbtree.Node) {
	t.head = h.(*node)
}
func (t *tree) GetHead() rbtree.Node {
	return t.head
}
func (t *tree) NewNode(key rbtree.Key) rbtree.Node {
	return &node{key: key.(int)}
}
func (t *tree) DeleteNode(n rbtree.Node) {
}
func (t *tree) IsNull(h rbtree.Node) bool {
	return h.(*node) == nil
}

func showTree(t *tree) {
	rbtree.Look(t, func(node rbtree.Node) {
		var L = "[L:-B]"
		var R = "[R:-B]"
		var C = "[-B]"
		if nL := t.GetL(node); !t.IsNull(nL) {
			if t.GetColor(nL) == rbtree.NodeColorBlack {
				L = fmt.Sprint("[", t.GetKey(nL), "B]")
			} else {
				L = fmt.Sprint("[", t.GetKey(nL), "R]")
			}
		}
		if nR := t.GetR(node); !t.IsNull(nR) {
			if t.GetColor(nR) == rbtree.NodeColorBlack {
				R = fmt.Sprint("[", t.GetKey(nR), "B]")
			} else {
				R = fmt.Sprint("[", t.GetKey(nR), "R]")
			}
		}
		if !t.IsNull(node) {
			if t.GetColor(node) == rbtree.NodeColorBlack {
				C = fmt.Sprint("[", t.GetKey(node), "B]")
			} else {
				C = fmt.Sprint("[", t.GetKey(node), "R]")
			}
		}
		fmt.Println(C, L, R)
	})
}

func main() {
	t := &tree{}
	for i := 0; i < 10000; i++ {
		rbtree.Insert(t, fmt.Sprint(i))
		if node, err := check.CaseAll(t); err != nil {
			fmt.Println(node, err)
		}
	}
	fmt.Println(rbtree.CheckCaseAll(t))
	showTree(t)

	fmt.Println("delete 17")
	rbtree.Delete(t, 17)
	fmt.Println(rbtree.CheckCaseAll(t))
	showTree(t)
	fmt.Println("end")
}
