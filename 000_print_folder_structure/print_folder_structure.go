package main

type Trie struct {
	root *Node
}

type Node struct {
	isEnd    bool
	Children map[string]*Node
}

func newNode() *Node {
	return &Node{Children: make(map[string]*Node)}
}

func (this *Trie) Insert(path string) {
	now := this.root
	pathChars := []rune{}
	for _, c := range path {
		if c == '/' && len(pathChars) != 0 {
			pathStr := string(pathChars)
			if now.Children[pathStr] == nil {
				now.Children[pathStr] = newNode()
			}
			now = now.Children[pathStr]
			pathChars = pathChars[:0]
		} else {
			pathChars = append(pathChars, c)
		}
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
