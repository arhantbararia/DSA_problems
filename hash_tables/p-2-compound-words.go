package main

import (
	"fmt"
	"strings"
)

const HASH_TABLE_SIZE = 1 << 18 // 2^18 = 262144

type WordNode struct {
	word string
	next *WordNode
}

func ooat(key []byte, length uint64, bits int) uint64 {
	var hash uint64 = 0
	var i uint64

	for i = 0; i < length; i++ {
		hash += uint64(key[i])
		hash += (hash << 10)
		hash ^= (hash >> 6)

	}

	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)

	return hash & ((uint64(1) << bits) - 1)

}
func InHashTable(hash_table []*WordNode, find string, findLen int) bool {
	word_code := ooat([]byte(find), uint64(findLen), 18)

	word_ptr := hash_table[word_code]

	for word_ptr != nil {
		if int(len(word_ptr.word)) == findLen &&
			(strings.Compare(word_ptr.word, find) == 0) {
			return true
		}
		word_ptr = word_ptr.next

	}
	return false
}

func identify_compound_words(words []string, hash_table []*WordNode, n int) {

	for i := 0; i < n; i++ {
		word_len := len(words[i])
		for j := 1; j < word_len; j++ {
			if InHashTable(hash_table, words[i][:j], j) && InHashTable(hash_table, words[i][j:], word_len-j) {
				fmt.Println(words[i])
				break
			}
		}

	}

}

func main() {
	words := []string{
		"cat",
		"dog",
		"cats",
		"dogs",
		"catdog",
		"dogcat",
		"apple",
		"pie",
		"applepie",
	}
	hashTable := make([]*WordNode, HASH_TABLE_SIZE)

	total_length := len(words)
	words_index := 0

	for words_index != total_length {

		word := words[words_index]

		length := len(word)
		word_code := ooat([]byte(word), uint64(length), 18)

		word_ptr := &WordNode{}

		word_ptr.word = word
		word_ptr.next = hashTable[word_code]

		hashTable[word_code] = word_ptr

		words_index++

	}

	identify_compound_words(words, hashTable, total_length)

}
