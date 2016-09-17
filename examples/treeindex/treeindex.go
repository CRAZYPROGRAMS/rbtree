package main

import (
	"fmt"

	"github.com/crazyprograms/rbtree"
)

type node struct {
	l, r  int
	color bool
	key   string
}

type tree struct {
	head  int
	nodes []node
}

func (t *tree) GetL(n rbtree.Node) rbtree.Node {
	return t.nodes[n.(int)].l
}
func (t *tree) SetL(n rbtree.Node, l rbtree.Node) {
	t.nodes[n.(int)].l = l.(int)
}
func (t *tree) GetR(n rbtree.Node) rbtree.Node {
	return t.nodes[n.(int)].r
}
func (t *tree) SetR(n rbtree.Node, r rbtree.Node) {
	t.nodes[n.(int)].r = r.(int)
}
func (t *tree) GetColor(n rbtree.Node) rbtree.NodeColor {
	return rbtree.NodeColor(t.nodes[n.(int)].color)
}
func (t *tree) SetColor(n rbtree.Node, color rbtree.NodeColor) {
	t.nodes[n.(int)].color = bool(color)
}
func (t *tree) GetKey(n rbtree.Node) rbtree.Key {
	return t.nodes[n.(int)].key
}
func (t *tree) LessKey(i, j rbtree.Key) bool {
	return i.(string) < j.(string)
}
func (t *tree) EqKey(i, j rbtree.Key) bool {
	return i.(string) == j.(string)
}

func (t *tree) SetHead(h rbtree.Node) {
	t.head = h.(int)
}
func (t *tree) GetHead() rbtree.Node {
	return t.head
}
func (t *tree) NewNode(key rbtree.Key) rbtree.Node {
	i := int(len(t.nodes))
	n := node{key: key.(string)}
	n.l = -1
	n.r = -1
	t.nodes = append(t.nodes, n)
	return i
}
func (t *tree) IsNull(h rbtree.Node) bool {
	return h.(int) == -1
}
func (t *tree) DeleteNode(n rbtree.Node) {

}
func (t *tree) EqNode(i, j rbtree.Node) bool {
	return i.(int) == j.(int)
}
func main() {
	t := &tree{nodes: make([]node, 0), head: -1}
	for i := 0; i < 100; i++ {
		rbtree.Insert(t, fmt.Sprint(i))
	}
	fmt.Println("Min:", rbtree.Min(t))
	fmt.Println("Max:", rbtree.Max(t))
	rbtree.LookWhere(t, func(node rbtree.Node) {
		fmt.Println(t.GetKey(node).(string))
	}, func(n rbtree.Node) bool {
		return t.nodes[n.(int)].key >= "30"
	}, func(n rbtree.Node) bool {
		return t.nodes[n.(int)].key <= "50"
	})
}
