type TrieNode struct {
	children [26]*TrieNode
	word bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: NewTrieNode() }
}

func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
    for _, c := range word {
		index := c-'a'
		if cur.children[index] == nil {
			cur.children[index] = NewTrieNode()
		}
		cur = cur.children[index]
	}
	cur.word = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool {
	cur := root
	for i := j; i < len(word); i++ {
		c := word[i]
		if c == '.' {
			for _, child := range cur.children {
				if child != nil &&  this.dfs(word, i+1, child) {
					return true
				}
			}
			return false
		} else {
			index := c-'a'
			if cur.children[index] == nil {
				return false
			}
			cur = cur.children[index]
		}
	}
	return cur.word
}
