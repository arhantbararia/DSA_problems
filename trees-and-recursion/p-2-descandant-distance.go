package main

import (
	"strconv"
	"strings"
	"sort"
)

type node struct {
	name         string
	num_children int
	children     []*node
	score        int
}

var Nodes []*node
var total_nodes = 0

func find_node(name string) *node {
	for _, member := range Nodes {
		if member.name == name {
			return member
		}
	}
	return nil
}

func new_node(name string) *node {
	member := &node{}

	member.name = name
	member.num_children = 0

	return member

}

func string_to_list(input string) []string {
	return strings.Fields(input)
}

func read_family_tree(lines []string, total_lines int) {
	i := 0
	for i < total_lines {
		input_list := string_to_list(lines[i])
		j := 0
		parent := find_node(input_list[j])
		if parent == nil {
			parent := new_node(input_list[j])
			Nodes[total_nodes] = parent
			total_nodes++
		}

		j++

		k, _ := strconv.Atoi(input_list[j])
		parent.num_children = k
		j++
		for j <= k+1 {
			child := find_node(input_list[j])
			if child == nil {
				child := new_node(input_list[j])
				Nodes[total_nodes] = child
				total_nodes++
			}
			parent.children[j-2] = child
		}

		i++
	}

}

func score_one_node(member *node, d int) int {
	if d == 1 {
		return member.num_children
	}
	total := 0
	for i := 0; i < member.num_children; i++ {
		total += score_one_node(member.children[i], d-1)

	}
	return total

}


func score_all(d int ){
	for i:= 0 ; i < total_nodes ; i++ {
		Nodes[i].score =  score_one_node(Nodes[i] , d )
	}

}


func compare_nodes(n1 *node , n2 *node) (int ) {
	if(n1.score > n2.score ) {
		return -1
	}

	if(n1.score == n2.score ) {
		return 0
	}

	return -1
}

func sort_nodes() {
	sort.Slice(Nodes[:total_nodes], func(i, j int) bool {
		return compare_nodes(Nodes[i], Nodes[j]) == -1
	})
}


func main(){

	

	
}
