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
	head                                    *node
	statGetL, statSetL                      int
	statGetR, statSetR                      int
	statGetColor, statSetColor              int
	statGetHead, statSetHead                int
	statGetKey, statEqKey, statLessKey      int
	statEqNode, statNewNode, statDeleteNode int
	statIsNull                              int
}

func (t *tree) ClearStat() {
	t.statGetL = 0
	t.statSetL = 0
	t.statGetR = 0
	t.statSetR = 0
	t.statGetColor = 0
	t.statSetColor = 0
	t.statGetHead = 0
	t.statSetHead = 0
	t.statGetKey = 0
	t.statEqKey = 0
	t.statLessKey = 0
	t.statEqNode = 0
	t.statNewNode = 0
	t.statDeleteNode = 0
	t.statIsNull = 0
}
func (t *tree) ShowStat() {
	fmt.Println("GetL      ", t.statGetL)
	fmt.Println("SetL      ", t.statSetL)
	fmt.Println("GetR      ", t.statGetR)
	fmt.Println("SetR      ", t.statSetR)
	fmt.Println("GetColor  ", t.statGetColor)
	fmt.Println("SetColor  ", t.statSetColor)
	fmt.Println("GetHead   ", t.statGetHead)
	fmt.Println("SetHead   ", t.statSetHead)
	fmt.Println("GetKey    ", t.statGetKey)
	fmt.Println("EqKey     ", t.statEqKey)
	fmt.Println("LessKey   ", t.statLessKey)
	fmt.Println("EqNode    ", t.statEqNode)
	fmt.Println("NewNode   ", t.statNewNode)
	fmt.Println("DeleteNode", t.statDeleteNode)
	fmt.Println("IsNull    ", t.statIsNull)
}
func (t *tree) GetL(n rbtree.Node) rbtree.Node {
	t.statGetL++
	return n.(*node).l
}
func (t *tree) SetL(n rbtree.Node, l rbtree.Node) {
	t.statSetL++
	n.(*node).l = l.(*node)
}
func (t *tree) GetR(n rbtree.Node) rbtree.Node {
	t.statGetR++
	return n.(*node).r
}
func (t *tree) SetR(n rbtree.Node, r rbtree.Node) {
	t.statSetR++
	n.(*node).r = r.(*node)
}
func (t *tree) GetColor(n rbtree.Node) rbtree.NodeColor {
	t.statGetColor++
	return rbtree.NodeColor(n.(*node).color)
}
func (t *tree) SetColor(n rbtree.Node, color rbtree.NodeColor) {
	t.statSetColor++
	n.(*node).color = bool(color)
}
func (t *tree) GetKey(n rbtree.Node) rbtree.Key {
	t.statGetKey++
	return n.(*node).key
}

func (t *tree) LessKey(i, j rbtree.Key) bool {
	t.statLessKey++
	return i.(int) < j.(int)
}
func (t *tree) EqKey(i, j rbtree.Key) bool {
	t.statEqKey++
	return i.(int) == j.(int)
}
func (t *tree) EqNode(i, j rbtree.Node) bool {
	t.statEqNode++
	return i.(*node) == j.(*node)
}
func (t *tree) SetHead(h rbtree.Node) {
	t.statSetHead++
	t.head = h.(*node)
}
func (t *tree) GetHead() rbtree.Node {
	t.statGetHead++
	return t.head
}
func (t *tree) NewNode(key rbtree.Key) rbtree.Node {
	t.statNewNode++
	return &node{key: key.(int)}
}
func (t *tree) DeleteNode(n rbtree.Node) {
	t.statDeleteNode++
}
func (t *tree) IsNull(h rbtree.Node) bool {
	t.statIsNull++
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

/*
insert 1000
GetL       14812
SetL       984
GetR       33539
SetR       13828
GetColor   31065
SetColor   6917
GetHead    1000
SetHead    1000
GetKey     12844
EqKey      0
LessKey    12844
EqNode     0
NewNode    1000
DeleteNode 0
IsNull     59243

delete 17
GetL       34
SetL       10
GetR       20
SetR       2
GetColor   40
SetColor   13
GetHead    2
SetHead    1
GetKey     17
EqKey      16
LessKey    14
EqNode     8
NewNode    0
DeleteNode 1
IsNull     42
*/
func main() {
	t := &tree{}

	for i := 0; i < 1000; i++ {
		rbtree.Insert(t, i)
		/*if node, err := rbtree.CheckCaseAll(t); err != nil {
			fmt.Println(node, err)
		}*/
	}
	t.ShowStat()
	fmt.Println(rbtree.CheckCaseAll(t))
	//showTree(t)

	fmt.Println("delete 17")
	t.ClearStat()
	rbtree.Delete(t, 17)
	t.ShowStat()
	fmt.Println(rbtree.CheckCaseAll(t))
	//showTree(t)
	fmt.Println("end")
}
