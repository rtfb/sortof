package main

import "fmt"

/*
diagram from here: http://cgi.cse.unsw.edu.au/~cs2521/18s2/tutes/week10/

             (5)
            /   \
           /     \
          /       \
         /         \
        (3)        (8)
       /   \      /   \
      (1)  (4)   (7)  (9)

infix order: 1 3 4 5 7 8 9
prefix order: 5 3 1 4 8 7 9
postfix order: 1 4 3 7 9 8 5
level order: 5 3 8 1 4 7 9
*/

type node struct {
	payload string
	left    *node
	right   *node
}

func traversePrefix(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.payload)
	traversePrefix(node.left)
	traversePrefix(node.right)
}

func traverseInfix(node *node) {
	if node == nil {
		return
	}
	traverseInfix(node.left)
	fmt.Println(node.payload)
	traverseInfix(node.right)
}

func traversePostfix(node *node) {
	if node == nil {
		return
	}
	traversePostfix(node.left)
	traversePostfix(node.right)
	fmt.Println(node.payload)
}

func traverseInverseLevelOrderStack(n *node) {
	type s struct {
		n     *node
		ldone bool
		rdone bool
	}
	stack := []s{{n: n}}
	for len(stack) != 0 {
		sv := &stack[len(stack)-1]
		n := sv.n
		if !sv.ldone {
			sv.ldone = true
			if n.left != nil {
				stack = append(stack, s{n: n.left})
				continue
			}
		}
		if !sv.rdone {
			sv.rdone = true
			if n.right != nil {
				stack = append(stack, s{n: n.right})
				continue
			}
		}
		fmt.Println(n.payload)
		stack = stack[0 : len(stack)-1]
	}
}

func traverseInfixStack(n *node) {
	type s struct {
		n     *node
		ldone bool
		rdone bool
	}
	stack := []s{{n: n}}
	for len(stack) != 0 {
		sv := &stack[len(stack)-1]
		n := sv.n
		if n.left != nil && !sv.ldone {
			sv.ldone = true
			stack = append(stack, s{n: n.left})
			continue
		}
		if !sv.rdone {
			fmt.Println(n.payload)
		}
		if n.right != nil && !sv.rdone {
			sv.rdone = true
			stack = append(stack, s{n: n.right})
			continue
		}
		stack = stack[0 : len(stack)-1]
	}
}

func main() {
	// tree := node{
	// 	payload: "R",
	// 	left: &node{
	// 		payload: "l",
	// 	},
	// 	right: &node{
	// 		payload: "r",
	// 	},
	// }
	t2 := node{
		payload: "A",
		left: &node{
			payload: "B",
			left: &node{
				payload: "D",
			},
			right: &node{
				payload: "E",
				left: &node{
					payload: "F",
				},
				right: &node{
					payload: "G",
				},
			},
		},
		right: &node{
			payload: "C",
		},
	}
	t3 := node{
		payload: "5",
		left: &node{
			payload: "3",
			left: &node{
				payload: "1",
			},
			right: &node{
				payload: "4",
			},
		},
		right: &node{
			payload: "8",
			left: &node{
				payload: "7",
			},
			right: &node{
				payload: "9",
			},
		},
	}
	// traversePrefix(&tree)
	// fmt.Println("\n")
	// traverseInfix(&tree)
	// fmt.Println("\n")
	traverseInfix(&t2)
	fmt.Println("\n")
	traversePrefix(&t2)
	fmt.Println("\n")
	traversePostfix(&t2)
	fmt.Println("\n")
	traverseInverseLevelOrderStack(&t2)
	fmt.Println("\n")
	traversePrefix(&t3)
	fmt.Println("\n")
	traverseInfixStack(&t3)
}
