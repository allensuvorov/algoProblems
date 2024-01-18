package main

import "fmt"

var data []string = []string{
	"/app/components/header",
	"/app/services",
	"/app/tests/components/header",
	"/images/image.png",
	"tsconfig.json",
	"Index.html",
}

type Trie struct {
	root *Node
}

type Node struct {
	isEnd    bool
	Children map[string]*Node
}

func newNode() *Node {
	return &Node{isEnd: false, Children: make(map[string]*Node)}
}

func (this *Trie) Insert(path string) {
	now := this.root
	name := []rune{}

	flush := func() {
		if len(name) != 0 {
			sName := string(name)
			if now.Children[sName] == nil {
				now.Children[sName] = newNode()
			}
			now = now.Children[sName]
			name = name[:0]
		}
	}
	for _, c := range path {
		if c == '/' {
			if len(name) != 0 {
				flush()
			}
		} else {
			name = append(name, c)
		}
	}
	flush()
	now.isEnd = true
}

func (this *Trie) dfsPrint(now *Node, pref []rune) {
	pref = append(pref, '-')
	if !now.isEnd {
		for val, node := range now.Children {
			fmt.Println(string(pref) + val)
			this.dfsPrint(node, pref)
		}
	}
}

func main() {
	t := new(Trie)
	t.root = newNode()
	for _, path := range data {
		t.Insert(path)
	}
	t.dfsPrint(t.root, []rune{})
	println("test")
}
