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
