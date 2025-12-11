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

type stack struct {
	values        []*node
	highest_index int
}

func NewStack() {
	node_stack := new(stack)
	node_stack.highest_index = -1
}

func (s *stack) Push(tree_node *node) {
	s.highest_index += 1
	s.values[s.highest_index] = tree_node

}

func (s *stack) Pop() (tree_node *node) {
	ret_node := s.values[s.highest_index]
	s.highest_index -= 1
	return ret_node
}

func (s *stack) is_empty() bool {
	return s.highest_index == -1
}

func candy_sum(tree_node *node, s stack) (total int) {

	for tree_node != nil {

		if tree_node.left != nil && tree_node.right != nil {
			s.Push(tree_node.left)
			tree_node = tree_node.right
		} else {
			total += tree_node.candy

			if s.is_empty() {
				tree_node = nil
			} else {
				tree_node = s.Pop()
			}
		}

	}
	return total

}
