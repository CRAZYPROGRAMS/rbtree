package main

import (
	"fmt"

	"github.com/crazyprograms/rbtree"
)

type node struct {
	l, r  *node
	color bool
	key   string
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
	return i.(string) < j.(string)
}
func (t *tree) EqKey(i, j rbtree.Key) bool {
	return i.(string) == j.(string)
}

func (t *tree) SetHead(h rbtree.Node) {
	t.head = h.(*node)
}
func (t *tree) GetHead() rbtree.Node {
	return t.head
}
func (t *tree) NewNode(key rbtree.Key) rbtree.Node {
	return &node{key: key.(string)}
}
func (t *tree) IsNull(h rbtree.Node) bool {
	return h.(*node) == nil
}
func (t *tree) DeleteNode(n rbtree.Node) {

}
func (t *tree) EqNode(i, j rbtree.Node) bool {
	return i.(*node) == j.(*node)
}
func main() {
	t := &tree{}
	for i := 0; i < 100; i++ {
		rbtree.Insert(t, fmt.Sprint(i))
	}
	fmt.Println("Min:", rbtree.Min(t))
	fmt.Println("Max:", rbtree.Max(t))
	rbtree.LookWhere(t, func(node rbtree.Node) {
		fmt.Println(t.GetKey(node).(string))
	}, func(n rbtree.Node) bool {
		return n.(*node).key >= "30"
	}, func(n rbtree.Node) bool {
		return n.(*node).key <= "50"
	})
}
