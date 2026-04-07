type PrefixTree struct {
	next [26]*PrefixTree
	ends bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	for _, c := range word {
		if this.next[c-'a'] == nil {
			this.next[c-'a'] = &PrefixTree{}
		}
		this = this.next[c-'a']
	}
	this.ends = true
}

func (this *PrefixTree) Search(word string) bool {
	for _, c := range word {
		if this.next[c-'a'] == nil {
			return false
		}
		this = this.next[c-'a']
	}
	if this.ends {
		return true
	}
	return false
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.next[c-'a'] == nil {
			return false
		}
		this = this.next[c-'a']
	}
	return true
}
