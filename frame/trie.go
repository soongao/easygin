package frame

type node struct {
	val              string
	isLeaf           string
	children         []*node
	childrenHaveStar bool
	isWild           bool
}

func (root *node) matchChild(val string) *node {
	for _, c := range root.children {
		if c.val == val {
			return c
		}
	}
	return nil
}

func (root *node) insert(pattern string, parts []string, hight int) {
	if hight == len(parts) {
		if root.isLeaf != "" {
			panic("pattern confict")
		}
		root.isLeaf = pattern
		return
	}

	part := parts[hight]
	child := root.matchChild(part)
	if child == nil {
		if (len(root.children) > 0 && part == "*") || (root.childrenHaveStar) {
			panic("pattern confict")
		}
		child = &node{
			val:              part,
			childrenHaveStar: part == "*",
			isWild:           part == "*" || part == ":",
		}
		if child.isWild {
			root.children = append(root.children, child)
		} else {
			root.children = append([]*node{child}, root.children...)
		}
	}
	child.insert(pattern, parts, hight+1)
}

func (root *node) matchChildren(val string) []*node {
	children := make([]*node, 0)
	for _, c := range root.children {
		if c.val == val || c.isWild {
			children = append(children, c)
		}
	}
	return children
}

func (root *node) search(parts []string, height int) string {
	if len(parts) == height {
		return root.isLeaf
	}
	part := parts[height]
	children := root.matchChildren(part)

	for _, c := range children {
		result := c.search(parts, height+1)
		if result != "" {
			return result
		}
	}
	return ""
}
