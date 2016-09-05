package rbtree

type NodeColor bool

const NodeColorRed NodeColor = false
const NodeColorBlack NodeColor = true

// Key - abstract type key
type Key interface{}

// Node - abstract type for the node
type Node interface{}

// Tree - Interface nodes store
type Tree interface {
	GetL(n Node) Node
	SetL(n Node, l Node)
	GetR(n Node) Node
	SetR(n Node, r Node)
	GetColor(n Node) NodeColor
	SetColor(n Node, color NodeColor)
	GetKey(n Node) Key
	LessKey(i, j Key) bool
	EqKey(i, j Key) bool
	SetHead(h Node)
	GetHead() Node
	NewNode(key Key) Node
	IsNull(h Node) bool
}

func rotR(tree Tree, h Node) Node {
	x := tree.GetL(h)
	tree.SetL(h, tree.GetR(x))
	tree.SetR(x, h)
	return x
}
func rotL(tree Tree, h Node) Node {
	x := tree.GetR(h)
	tree.SetR(h, tree.GetL(x))
	tree.SetL(x, h)
	return x
}

func rbInsert(tree Tree, h Node, key Key, sw bool) Node {
	if tree.IsNull(h) {
		node := tree.NewNode(key)
		tree.SetColor(node, NodeColorRed)
		return node
	}
	if hL := tree.GetL(h); !tree.IsNull(hL) && tree.GetColor(hL) == NodeColorRed {
		if hR := tree.GetR(h); !tree.IsNull(hR) && tree.GetColor(hR) == NodeColorRed {
			tree.SetColor(h, NodeColorRed)
			tree.SetColor(hL, NodeColorBlack)
			tree.SetColor(hR, NodeColorBlack)
		}
	}
	if tree.LessKey(key, tree.GetKey(h)) {
		hL := rbInsert(tree, tree.GetL(h), key, false)
		tree.SetL(h, hL)
		if !tree.IsNull(hL) && sw && tree.GetColor(h) == NodeColorRed && tree.GetColor(hL) == NodeColorRed {
			h = rotR(tree, h)
		}

		if hL := tree.GetL(h); !tree.IsNull(hL) && tree.GetColor(hL) == NodeColorRed {
			if hLL := tree.GetL(hL); !tree.IsNull(hLL) && tree.GetColor(hLL) == NodeColorRed {
				h = rotR(tree, h)
				tree.SetColor(h, NodeColorBlack)
				tree.SetColor(tree.GetR(h), NodeColorRed)
			}
		}
	} else {
		hR := rbInsert(tree, tree.GetR(h), key, true)
		tree.SetR(h, hR)
		if !tree.IsNull(hR) && !sw && tree.GetColor(h) == NodeColorRed && tree.GetColor(hR) == NodeColorRed {
			h = rotL(tree, h)
		}

		if hR := tree.GetR(h); !tree.IsNull(hR) && tree.GetColor(hR) == NodeColorRed {
			if hRR := tree.GetR(hR); !tree.IsNull(hRR) && tree.GetColor(hRR) == NodeColorRed {
				h = rotL(tree, h)
				tree.SetColor(h, NodeColorBlack)
				tree.SetColor(tree.GetL(h), NodeColorRed)
			}
		}
	}
	return h
}

// Insert - Inserting a new node
func Insert(tree Tree, key Key) {
	h := rbInsert(tree, tree.GetHead(), key, false)
	tree.SetColor(h, NodeColorBlack)
	tree.SetHead(h)
}

func search(tree Tree, h Node, key Key) Node {
	if tree.IsNull(h) {
		return nil
	}
	key2 := tree.GetKey(h)
	for !tree.EqKey(key, key2) {
		if tree.LessKey(key, key2) {
			h = tree.GetL(h)
		} else {
			h = tree.GetR(h)
		}
		if tree.IsNull(h) {
			return nil
		}
		key2 = tree.GetKey(h)
	}
	return h
}

// Search - node search
func Search(tree Tree, key Key) Node {
	return search(tree, tree.GetHead(), key)
}

// LoopItem - Callback delegat for look tree
type LookItem func(node Node)

type WhereItem func(node Node) bool

func look(tree Tree, h Node, item LookItem, Min WhereItem, Max WhereItem) (count uint) {
	count = 0
	L := tree.GetL(h)
	R := tree.GetR(h)
	min := Min(h)
	max := Max(h)
	if !tree.IsNull(L) && min {
		count += look(tree, L, item, Min, Max)
	}
	if min && max {
		count++
		item(h)
	}
	if !tree.IsNull(R) && max {
		count += look(tree, R, item, Min, Max)
	}
	return count
}

//Look - View whole tree
func Look(tree Tree, item LookItem) uint {
	h := tree.GetHead()
	if tree.IsNull(h) {
		return 0
	}
	return look(tree, h, item, func(Node) bool { return true }, func(Node) bool { return true })
}

// LoopWhere - View of the tree from the min to the max
func LoopWhere(tree Tree, item LookItem, Min WhereItem, Max WhereItem) uint {
	h := tree.GetHead()
	if tree.IsNull(h) {
		return 0
	}
	return look(tree, h, item, Min, Max)
}

//Min - min key
func Min(tree Tree) Key {
	h := tree.GetHead()
	if tree.IsNull(h) {
		return nil
	}
	hL := tree.GetL(h)
	for !tree.IsNull(hL) {
		h = hL
		hL = tree.GetL(h)
	}
	return tree.GetKey(h)
}

//Max - max key
func Max(tree Tree) Key {
	h := tree.GetHead()
	if tree.IsNull(h) {
		return nil
	}
	hR := tree.GetR(h)
	for !tree.IsNull(hR) {
		h = hR
		hR = tree.GetR(h)
	}
	return tree.GetKey(h)
}
