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

func parseTree_r(input string) ([]*Tree, int) {
	if input == "" || input == "()" {
		return nil, 0
	}
	result := []*Tree(nil)
	i := 0
	if input[i] == '(' {
		i += 1
	}
	for i < len(input) {
		if input[i] == ' ' {
			i += 1
			continue
		}
		if input[i] == ')' {
			return result, i
		}
		if result == nil {
			result = make([]*Tree, 0, 0)
		}
		if input[i] == '(' {
			i += 1
			kids, charsEaten := parseTree_r(input[i:len(input)])
			result[len(result)-1].children = kids
			i += charsEaten + 1
			continue
		}
		result = append(result, &Tree{input[i : i+1], nil})
		i += 1
	}
	return nil, 0
}

func parseTree(input string) *Tree {
	if input == "" || input == "()" {
		return nil
	}
	if input[0] == '(' && len(input) > 1 && input[1] != ' ' {
		kids, _ := parseTree_r(input[2:len(input)])
		return &Tree{data: input[1:2], children: kids}
	}
	return nil
}
