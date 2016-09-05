package check

import (
	"errors"

	"github.com/crazyprograms/rbtree"
)

func case4(t rbtree.Tree, n rbtree.Node) (rbtree.Node, error) {
	if t.IsNull(n) {
		return nil, nil
	}
	nL := t.GetL(n)
	nR := t.GetR(n)
	if t.GetColor(n) == rbtree.NodeColorRed {
		if !t.IsNull(nL) && t.GetColor(nL) == rbtree.NodeColorRed {
			return n, errors.New("rbtree case4")
		}
	}
	if n, err := case4(t, nL); err != nil {
		return n, err
	}
	if n, err := case4(t, nR); err != nil {
		return n, err
	}
	return nil, nil
}
func case5(t rbtree.Tree, n rbtree.Node) (num int, node rbtree.Node, err error) {
	if t.IsNull(n) {
		return 1, nil, nil // case3 All leaves (NIL) are black.
	}
	nL := t.GetL(n)
	nR := t.GetR(n)
	var numL int
	var numR int
	if numL, node, err = case5(t, nL); err != nil {
		return 0, node, err
	}
	if numR, node, err = case5(t, nR); err != nil {
		return 0, node, err
	}
	if numL != numR {
		return 0, n, errors.New("rbtree case5")
	}
	return numL, nil, nil
}

// Case2 The root is black.
func Case2(t rbtree.Tree) error {
	if t.GetColor(t.GetHead()) == rbtree.NodeColorRed {
		return errors.New("rbtree case2")
	}
	return nil
}

// Case4 If a node is red, then both its children are black.
func Case4(t rbtree.Tree) (rbtree.Node, error) {
	return case4(t, t.GetHead())
}

// Case5 Every simple path from a given node to any leaf node, being a descendant of his, contains the same number of black nodes.
func Case5(t rbtree.Tree) (rbtree.Node, error) {
	if _, node, err := case5(t, t.GetHead()); err != nil {
		return node, err
	}
	return nil, nil
}

// CaseSort traversal of the tree from left to right gives the sorted keys
func CaseSort(t rbtree.Tree) error {
	var oldKey rbtree.Key
	var set bool
	var err error
	rbtree.Look(t, func(n rbtree.Node) {
		key := t.GetKey(n)
		if set && !t.LessKey(oldKey, key) {
			err = errors.New("rbtree sort")
		}
		oldKey = key
		set = true
	})
	return err
}

// CaseAll Check all conditions
func CaseAll(t rbtree.Tree) (rbtree.Node, error) {
	if err := Case2(t); err != nil {
		return t.GetHead(), err
	}
	if n, err := Case4(t); err != nil {
		return n, err
	}
	if n, err := Case5(t); err != nil {
		return n, err
	}
	if err := Case2(t); err != nil {
		return nil, CaseSort(t)
	}

	return nil, nil
}
