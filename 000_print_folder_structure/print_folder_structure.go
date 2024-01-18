package main

type Trie struct {
	root *Node
}

type Node struct {
	isEnd    bool
	Children map[byte]*Node
}

func (this Trie) Insert(path string) {

}

func (this Trie) Print() {

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
