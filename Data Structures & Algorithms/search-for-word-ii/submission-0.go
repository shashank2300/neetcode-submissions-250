type TrieNode struct {
	children map[byte]*TrieNode
	word bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

func (this *TrieNode) addWord(word string) {
	cur := this
	for i := range word {
		c := word[i]
		if _, found := cur.children[c]; !found {
			cur.children[c] = NewTrieNode()
		}
		cur = cur.children[c]
	}
	cur.word = true
}


func findWords(board [][]byte, words []string) []string {
    root := NewTrieNode()
	for _, word := range words {
		root.addWord(word)
	}

	rows, cols := len(board), len(board[0])
	var res []string
	visit := make(map[[2]int]bool)
	wordSet := make(map[string]bool)

	var dfs func(r, c int, node *TrieNode, word string)
	dfs = func(r, c int, node *TrieNode, word string) {
		if r < 0 || c < 0 || r >= rows || c >= cols || visit[[2]int{r, c}] {
			return
		}

		char := board[r][c]
		nextNode, found := node.children[char]
		if !found {
			return
		}

		visit[[2]int{r, c}] = true
		word += string(char)
		if nextNode.word {
			wordSet[word] = true
		}

		dfs(r+1, c, nextNode, word)
		dfs(r-1, c, nextNode, word)
		dfs(r, c+1, nextNode, word)
		dfs(r, c-1, nextNode, word)

		visit[[2]int{r, c}] = false
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r, c, root, "")
		}
	}

	for word := range wordSet {
		res = append(res, word)
	}
	return res
}
