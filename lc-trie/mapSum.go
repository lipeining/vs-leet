package lc

const MAXCAP = 26 // a-z
type Trie struct {
	next   map[rune]*Trie
	cnt    int
	isWord bool
}
type MapSum struct {
	root *Trie
}

// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["apple",3], ["ap"], ["app",2], ["ap"]]

// ["MapSum", "insert", "sum", "insert", "sum"]
// [[], ["aa",3], ["a"], ["aa",2], ["a"]]

// ["MapSum", "insert", "sum", "insert", "sum", "sum", "insert", "sum", "sum", "sum", "insert", "sum", "sum", "sum", "sum", "sum", "insert", "insert", "insert", "sum", "sum", "sum"]
// [[], ["aa",3], ["a"], ["aa",2], ["a"], ["aa"], ["aaa", 3], ["aaa"], ["bbb"], ["ccc"], ["aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj", 111], ["aa"], ["a"], ["b"], ["c"], ["aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj"], ["bb", 143], ["cc", 144], ["aew", 145], ["bb"], ["cc"], ["aew"]]

// 输出：
// [null,null,3,null,2,2,null,3,0,0,null,5,116,0,0,111,null,null,null,143,144,145]
// 预期结果：
// [null,null,3,null,2,2,null,3,0,0,null,5,116,0,0,111,null,null,null,143,144,256]

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string, val int) {
	for _, v := range word {
		if this.next[v] == nil {
			node := new(Trie)
			node.next = make(map[rune]*Trie)
			//初始化节点单词标志为假
			node.isWord = false
			node.cnt = 0
			this.next[v] = node
		}
		this = this.next[v]
	}
	// 只让最终节点有值
	this.cnt = val
	this.isWord = true
}

/** Returns if the prefix is in the trie. */
func (this *Trie) Sum(prefix string) int {
	for _, v := range prefix {
		if this.next[v] == nil {
			return 0
		}
		this = this.next[v]
	}
	return this.SubSum()
}

func (this *Trie) SubSum() int {
	// 从这里开始，之后的单词都可以计算和
	ans := 0
	if this.isWord {
		ans += this.cnt
	}
	for _, v := range this.next {
		ans += v.SubSum()
	}
	return ans
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return this.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v] == nil {
			return false
		}
		this = this.next[v]
	}
	return true
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	root := &Trie{next: make(map[rune]*Trie), cnt: 0, isWord: false}
	return MapSum{root}
}

func (this *MapSum) Insert(key string, val int) {
	this.root.Insert(key, val)
}

func (this *MapSum) Sum(prefix string) int {
	return this.root.Sum(prefix)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
