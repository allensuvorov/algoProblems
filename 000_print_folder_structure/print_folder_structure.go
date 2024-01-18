package main

type Trie struct {
	root *Node
}

type Node struct {
	isEnd    bool
	Children map[rune]*Node
}

func newNode() *Node {
	return &Node{Children: make(map[rune]*Node)}
}

func (this *Trie) Insert(path string) {
	now := this.root
	for _, c := range path {
		if now.Children[c] == nil {
			now.Children[c] = newNode()
		}
		now = now.Children[c]
	}
}

func (this *Trie) Print() {
	now := this.root
}

func main() {
	println("test")
}

// Suppose you have a file named           "/home/files/xyz"

// Print out the directory structure of the file like the following

// - home
// -- files
// --- xyz

// Given the following file names

// /app/components/header
// /app/services
// /app/tests/components/header
// /images/image.png
// tsconfig.json
// Index.html

// - app
// -- components
// --- header
// -- services
// -- tests
// --- components
// ---- header
// - images
// -- Image.png
// - tsconfig.json
// - index.html
