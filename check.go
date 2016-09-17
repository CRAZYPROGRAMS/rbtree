package rbtree

import (
	"errors"
)

func case4(t Tree, n Node) (Node, error) {
	if t.IsNull(n) {
		return nil, nil
	}
	nL := t.GetL(n)
	nR := t.GetR(n)
	if t.GetColor(n) == NodeColorRed {
		if !t.IsNull(nL) && t.GetColor(nL) == NodeColorRed {
			return n, errors.New("rbtree case4")
		}
		if !t.IsNull(nR) && t.GetColor(nR) == NodeColorRed {
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
func case5(t Tree, n Node) (num int, node Node, err error) {
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
	if t.GetColor(n) == NodeColorBlack {
		numL++
	}
	return numL, nil, nil
}

// Case2 The root is black.
func CheckCase2(t Tree) error {
	if t.GetColor(t.GetHead()) == NodeColorRed {
		return errors.New("rbtree case2")
	}
	return nil
}

// Case4 If a node is red, then both its children are black.
func CheckCase4(t Tree) (Node, error) {
	return case4(t, t.GetHead())
}

// Case5 Every simple path from a given node to any leaf node, being a descendant of his, contains the same number of black nodes.
func CheckCase5(t Tree) (Node, error) {
	if _, node, err := case5(t, t.GetHead()); err != nil {
		return node, err
	}
	return nil, nil
}

// CaseSort traversal of the tree from left to right gives the sorted keys
func CheckCaseSort(t Tree) error {
	var oldKey Key
	var set bool
	var err error
	Look(t, func(n Node) {
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
func CheckCaseAll(t Tree) (Node, error) {
	if err := CheckCase2(t); err != nil {
		return t.GetHead(), err
	}
	if n, err := CheckCase4(t); err != nil {
		return n, err
	}
	if n, err := CheckCase5(t); err != nil {
		return n, err
	}
	if err := CheckCaseSort(t); err != nil {
		return nil, err
	}

	return nil, nil
}
