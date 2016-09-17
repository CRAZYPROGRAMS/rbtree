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
	EqNode(i, j Node) bool
	SetHead(h Node)
	GetHead() Node
	NewNode(key Key) Node
	DeleteNode(Node)
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

func rbInsert(tree Tree, h Node, key Key, sw bool) (head Node, editHead bool) {
	var edit bool
	if tree.IsNull(h) {
		node := tree.NewNode(key)
		tree.SetColor(node, NodeColorRed)
		edit = true
		return node, edit
	}
	if hL := tree.GetL(h); !tree.IsNull(hL) && tree.GetColor(hL) == NodeColorRed {
		if hR := tree.GetR(h); !tree.IsNull(hR) && tree.GetColor(hR) == NodeColorRed {
			tree.SetColor(h, NodeColorRed)
			tree.SetColor(hL, NodeColorBlack)
			tree.SetColor(hR, NodeColorBlack)
		}
	}
	if tree.LessKey(key, tree.GetKey(h)) {
		hL, hLEdit := rbInsert(tree, tree.GetL(h), key, false)
		if hLEdit {
			tree.SetL(h, hL)
		}
		if !tree.IsNull(hL) && sw && tree.GetColor(h) == NodeColorRed && tree.GetColor(hL) == NodeColorRed {
			h = rotR(tree, h)
			edit = true
		}

		if hL := tree.GetL(h); !tree.IsNull(hL) && tree.GetColor(hL) == NodeColorRed {
			if hLL := tree.GetL(hL); !tree.IsNull(hLL) && tree.GetColor(hLL) == NodeColorRed {
				h = rotR(tree, h)
				edit = true
				tree.SetColor(h, NodeColorBlack)
				tree.SetColor(tree.GetR(h), NodeColorRed)
			}
		}
	} else {
		hR, hREdit := rbInsert(tree, tree.GetR(h), key, true)
		if hREdit {
			tree.SetR(h, hR)
		}
		if !tree.IsNull(hR) && !sw && tree.GetColor(h) == NodeColorRed && tree.GetColor(hR) == NodeColorRed {
			h = rotL(tree, h)
			edit = true
		}

		if hR := tree.GetR(h); !tree.IsNull(hR) && tree.GetColor(hR) == NodeColorRed {
			if hRR := tree.GetR(hR); !tree.IsNull(hRR) && tree.GetColor(hRR) == NodeColorRed {
				h = rotL(tree, h)
				edit = true
				tree.SetColor(h, NodeColorBlack)
				tree.SetColor(tree.GetL(h), NodeColorRed)
			}
		}
	}
	return h, edit
}

// Insert - Inserting a new node
func Insert(tree Tree, key Key) {
	h, edit := rbInsert(tree, tree.GetHead(), key, false)
	tree.SetColor(h, NodeColorBlack)
	if edit {
		tree.SetHead(h)
	}
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

// LookItem - Callback delegat for look tree
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
func LookWhere(tree Tree, item LookItem, Min WhereItem, Max WhereItem) uint {
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

func transplant(tree Tree, u Node, uP Node, v Node, left bool) {
	if uP == nil {
		tree.SetHead(v)
	} else {
		if left {
			tree.SetL(uP, v)
		} else {
			tree.SetR(uP, v)
		}
	}
}

func minP(tree Tree, n Node, nP Node) (m Node, mP Node, min bool) {
	m = n
	mP = nP
	min = true
	mL := tree.GetL(m)
	for !tree.IsNull(mL) {
		mP = m
		m = mL
		min = false
		mL = tree.GetL(m)
	}
	return m, mP, min
}

func delete2(tree Tree, h Node, p Node, key Key, left bool) (x, xP Node, fix bool) {
	if tree.IsNull(h) {
		return nil, nil, false
	}
	hKey := tree.GetKey(h)
	if tree.EqKey(key, hKey) {
		var x1, x1P Node
		x1, x1P = deleteNode(tree, h, p, left)
		return x1, x1P, true
	}
	if tree.LessKey(key, hKey) {
		return delete2(tree, tree.GetL(h), h, key, true)
	} else {
		return delete2(tree, tree.GetR(h), h, key, false)
	}
}

func dnodeFixup(tree Tree, x Node, xP Node) (head Node, fix, editHead bool) {
	var edit bool
	if !tree.IsNull(x) && tree.GetColor(x) != NodeColorBlack {
		tree.SetColor(x, NodeColorBlack)
		edit = false
		return xP, false, edit
	}
	hOld := xP
	xPL := tree.GetL(xP)
	if tree.EqNode(x, xPL) {
		w := tree.GetR(xP)
		wColor := tree.GetColor(w)
		if wColor == NodeColorRed {
			tree.SetColor(w, NodeColorBlack)
			tree.SetColor(xP, NodeColorRed)
			xP = rotL(tree, xP)
			edit = true
			w = tree.GetR(hOld)
		}
		wL := tree.GetL(w)
		wR := tree.GetR(w)
		wLColor := NodeColorBlack
		wRColor := NodeColorBlack
		if !tree.IsNull(wL) {
			wLColor = tree.GetColor(wL)
		}
		if !tree.IsNull(wR) {
			wRColor = tree.GetColor(wR)
		}
		if wLColor == NodeColorBlack && wRColor == NodeColorBlack {
			tree.SetColor(w, NodeColorRed)
			if tree.GetColor(hOld) != NodeColorBlack {
				tree.SetColor(hOld, NodeColorBlack)
				return xP, false, edit
			}
			return xP, true, edit
		} else {
			if wRColor == NodeColorBlack {
				tree.SetColor(wL, NodeColorBlack)
				tree.SetColor(w, NodeColorRed)
				w = rotR(tree, w)
				tree.SetR(xP, w)
			}
			tree.SetColor(w, tree.GetColor(xP))
			tree.SetColor(xP, NodeColorBlack)
			tree.SetColor(wR, NodeColorBlack)
			xP = rotL(tree, xP)
			edit = true
			return xP, false, edit
		}
	} else {
		w := tree.GetL(xP)
		wColor := tree.GetColor(w)
		if wColor == NodeColorRed {
			tree.SetColor(w, NodeColorBlack)
			tree.SetColor(xP, NodeColorRed)
			xP = rotR(tree, xP)
			edit = true
			w = tree.GetL(hOld)
		}
		wL := tree.GetL(w)
		wR := tree.GetR(w)
		wLColor := NodeColorBlack
		wRColor := NodeColorBlack
		if !tree.IsNull(wL) {
			wLColor = tree.GetColor(wL)
		}
		if !tree.IsNull(wR) {
			wRColor = tree.GetColor(wR)
		}
		if wLColor == NodeColorBlack && wRColor == NodeColorBlack {
			tree.SetColor(w, NodeColorRed)
			if tree.GetColor(hOld) != NodeColorBlack {
				tree.SetColor(hOld, NodeColorBlack)
				return xP, false, edit
			}
			return xP, true, edit
		} else {
			if wLColor == NodeColorBlack {
				tree.SetColor(wR, NodeColorBlack)
				tree.SetColor(w, NodeColorRed)
				w = rotL(tree, w)
				tree.SetL(xP, w)
			}
			tree.SetColor(w, tree.GetColor(xP))
			tree.SetColor(xP, NodeColorBlack)
			tree.SetColor(wL, NodeColorBlack)
			xP = rotR(tree, xP)
			edit = true
			return xP, false, edit
		}
	}
}

func deleteFixup(tree Tree, x, h Node, key Key) (head Node, fix, editHead bool) {
	var fix1, edit bool
	var x1 Node
	hKey := tree.GetKey(h)
	if tree.EqKey(key, hKey) {
		return dnodeFixup(tree, x, h)
	}
	if tree.LessKey(key, hKey) {
		x1, fix1, edit = deleteFixup(tree, x, tree.GetL(h), key)
		if edit {
			tree.SetL(h, x1)
		}
	} else {
		x1, fix1, edit = deleteFixup(tree, x, tree.GetR(h), key)
		if edit {
			tree.SetR(h, x1)
		}
	}
	if !tree.IsNull(x1) {
		x = x1
	}
	if !fix1 {
		return h, false, false
	}
	return dnodeFixup(tree, x, h)
}

func deleteNode(tree Tree, z Node, zP Node, left bool) (node, parent Node) {
	var y Node
	var x Node
	var yP Node
	var xP Node
	y = z
	yOrigColor := tree.GetColor(y)
	zL := tree.GetL(z)
	zR := tree.GetR(z)
	if tree.IsNull(zL) {
		xP = zP
		x = zR
		transplant(tree, z, zP, zR, left)
	} else if tree.IsNull(zR) {
		xP = zP
		x = zL
		transplant(tree, z, zP, zL, left)
	} else {
		var min bool
		y, yP, min = minP(tree, zR, z)

		yOrigColor = tree.GetColor(y)
		x = tree.GetR(y)
		if min {
			xP = y
			tree.SetL(yP, tree.GetL(yP))
		} else {
			xP = yP
			transplant(tree, y, yP, x, true)
			tree.SetR(y, zR)
		}
		transplant(tree, z, zP, y, left)
		tree.SetL(y, zL)
		tree.SetColor(y, tree.GetColor(z))
	}
	tree.DeleteNode(z)
	if yOrigColor == NodeColorBlack && xP != nil {
		return x, xP
	}
	return nil, nil
}

// Delete - node remove
func Delete(tree Tree, key Key) bool {
	var x, xP Node
	var d bool
	if x, xP, d = delete2(tree, tree.GetHead(), nil, key, false); d && x != nil {
		h, _, edit := deleteFixup(tree, x, tree.GetHead(), tree.GetKey(xP))
		tree.SetColor(h, NodeColorBlack)
		if edit {
			tree.SetHead(h)
		}
	}
	return d
}
