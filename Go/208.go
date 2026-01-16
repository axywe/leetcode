package main

type Trie struct {
	Children map[byte]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{make(map[byte]*Trie), false}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		if cur.Children[word[i]] == nil {
			cur.Children[word[i]] = &Trie{make(map[byte]*Trie), false}
		}
		cur = cur.Children[word[i]]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		if cur.Children[word[i]] == nil {
			return false
		}
		cur = cur.Children[word[i]]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		if cur.Children[prefix[i]] == nil {
			return false
		}
		cur = cur.Children[prefix[i]]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
