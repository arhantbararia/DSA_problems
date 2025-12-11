package main

type node struct {
	candy int
	left  *node
	right *node
}

func new_house(candy int) *node {
	new_node := new(node)

	new_node.candy = candy
	new_node.left = nil
	new_node.right = nil

	return new_node
}

func new_non_house(left, right *node) *node {
	new_node := new(node)
	new_node.left = left
	new_node.right = right
	new_node.candy = -1
}

func tree_candy(tree *node) int {
	if tree.left == nil && tree.right == nil {
		return tree.candy
	}

	return tree_candy(tree.left) + tree_candy(tree.right)
}

////////////////]

func total_nodes(tree *node) int {
	if tree.left == nil && tree.right == nil {
		return 1
	}

	return 1 + total_nodes(tree.left) + total_nodes(tree.right)
}

func total_leaf_nodes(tree *node) int {
	if tree.left == nil && tree.right == nil {
		return 1
	}

	return total_leaf_nodes(tree.left) + total_leaf_nodes(tree.right)
}

////////////////]

// Walking min streets

//to calculate all candies we have to travel the whole tree

func tree_streets(tree *node) int {
	if tree.left == nil && tree.right == nil {
		return 0
	}

	return 4 + tree_streets(tree.left) + tree_streets(tree.right)

}

// to end jorney at farthest house from node. min_streets = total_streets - height

func height(tree *node) int {
	if tree.left == nil && tree.right == nil {
		return 0
	}

	return 1 + max(height(tree.left), height(tree.right))
}

func tree_solve(tree *node) (streets int, candy int) {

	candy = tree_candy(tree)

	min_streets := tree_streets(tree) - height(tree)

	return min_streets, candy
}

// input binary tree given like input = ((4,9) , 15)
func readTreeHelper(input string, pos *int) *node {
	tree := &node{}

	if input[*pos] == '(' {
		*pos++
		//skip '('
		tree.left = readTreeHelper(input, pos)

		*pos++ //skip seperator ' '
		tree.right = readTreeHelper(input, pos)

		*pos++ //after ')'
		return tree

	}

	tree.left = nil
	tree.right = nil

	//First Digit
	tree.candy = int(input[*pos] - '0')
	*pos++

	for *pos < len(input) && input[*pos] != ')' && input[*pos] != ' ' && input[*pos] != '\x00' {
		tree.candy = tree.candy*10 + int(input[*pos]-'0')
		*pos++
	}

	return tree

}

func main() {

	//binary tree given like ((4,9) , 15)

	inputs = []string{
		""
	}


	for _ , inp := inputs {
		
	}


}
