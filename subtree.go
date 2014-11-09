package sortof

type Tree struct {
	data     string
	children []*Tree
}

func printTree_r(input *Tree, level int) string {
	if input == nil {
		return ""
	}
	var result string
	if level > 1 {
		result += "\n"
		for i := 1; i < level; i += 1 {
			result += ".\t"
		}
	}
	result += input.data
	for _, child := range input.children {
		result += printTree_r(child, level+1)
	}
	return result
}

func printTree(input *Tree) string {
	return printTree_r(input, 1)
}

func compareTrees(t1, t2 *Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.data == t2.data && t1.children == nil && t2.children == nil {
		return true
	}
	if t1.data == t2.data && len(t1.children) == len(t2.children) {
		for i, child := range t1.children {
			return compareTrees(child, t2.children[i])
		}
	}
	return false
}
