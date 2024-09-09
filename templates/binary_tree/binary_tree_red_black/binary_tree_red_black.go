package redblacktree

import "fmt"

type Color bool

const (
    RED   Color = true
    BLACK Color = false
)

// Node represents a node in the Red-Black tree
type Node struct {
    Key    int
    Color  Color
    Left   *Node
    Right  *Node
    Parent *Node
}

// RedBlackTree represents the Red-Black tree
type RedBlackTree struct {
    Root *Node
    NIL  *Node
}

// NewRedBlackTree creates a new Red-Black tree
func NewRedBlackTree() *RedBlackTree {
    nil_node := &Node{Color: BLACK}
    return &RedBlackTree{NIL: nil_node, Root: nil_node}
}

// Insert inserts a new key into the tree
func (t *RedBlackTree) Insert(key int) {
    node := &Node{Key: key, Color: RED, Left: t.NIL, Right: t.NIL, Parent: t.NIL}
    var y *Node = t.NIL
    x := t.Root

    for x != t.NIL {
        y = x
        if node.Key < x.Key {
            x = x.Left
        } else {
            x = x.Right
        }
    }

    node.Parent = y
    if y == t.NIL {
        t.Root = node
    } else if node.Key < y.Key {
        y.Left = node
    } else {
        y.Right = node
    }

    t.insertFixup(node)
}

// insertFixup fixes the Red-Black tree after insertion
func (t *RedBlackTree) insertFixup(z *Node) {
    for z.Parent.Color == RED {
        if z.Parent == z.Parent.Parent.Left {
            y := z.Parent.Parent.Right
            if y.Color == RED {
                z.Parent.Color = BLACK
                y.Color = BLACK
                z.Parent.Parent.Color = RED
                z = z.Parent.Parent
            } else {
                if z == z.Parent.Right {
                    z = z.Parent
                    t.leftRotate(z)
                }
                z.Parent.Color = BLACK
                z.Parent.Parent.Color = RED
                t.rightRotate(z.Parent.Parent)
            }
        } else {
            y := z.Parent.Parent.Left
            if y.Color == RED {
                z.Parent.Color = BLACK
                y.Color = BLACK
                z.Parent.Parent.Color = RED
                z = z.Parent.Parent
            } else {
                if z == z.Parent.Left {
                    z = z.Parent
                    t.rightRotate(z)
                }
                z.Parent.Color = BLACK
                z.Parent.Parent.Color = RED
                t.leftRotate(z.Parent.Parent)
            }
        }
    }
    t.Root.Color = BLACK
}

// leftRotate performs a left rotation
func (t *RedBlackTree) leftRotate(x *Node) {
    y := x.Right
    x.Right = y.Left
    if y.Left != t.NIL {
        y.Left.Parent = x
    }
    y.Parent = x.Parent
    if x.Parent == t.NIL {
        t.Root = y
    } else if x == x.Parent.Left {
        x.Parent.Left = y
    } else {
        x.Parent.Right = y
    }
    y.Left = x
    x.Parent = y
}

// rightRotate performs a right rotation
func (t *RedBlackTree) rightRotate(x *Node) {
    y := x.Left
    x.Left = y.Right
    if y.Right != t.NIL {
        y.Right.Parent = x
    }
    y.Parent = x.Parent
    if x.Parent == t.NIL {
        t.Root = y
    } else if x == x.Parent.Right {
        x.Parent.Right = y
    } else {
        x.Parent.Left = y
    }
    y.Right = x
    x.Parent = y
}

// Search searches for a key in the tree
func (t *RedBlackTree) Search(key int) bool {
    return t.searchHelper(t.Root, key)
}

func (t *RedBlackTree) searchHelper(node *Node, key int) bool {
    if node == t.NIL {
        return false
    }
    if key == node.Key {
        return true
    }
    if key < node.Key {
        return t.searchHelper(node.Left, key)
    }
    return t.searchHelper(node.Right, key)
}

// InorderTraversal performs an inorder traversal of the tree
func (t *RedBlackTree) InorderTraversal() []int {
    var result []int
    t.inorderTraversalHelper(t.Root, &result)
    return result
}

func (t *RedBlackTree) inorderTraversalHelper(node *Node, result *[]int) {
    if node != t.NIL {
        t.inorderTraversalHelper(node.Left, result)
        *result = append(*result, node.Key)
        t.inorderTraversalHelper(node.Right, result)
    }
}

// PrintTree prints the tree structure
func (t *RedBlackTree) PrintTree() {
    t.printTreeHelper(t.Root, 0)
}

func (t *RedBlackTree) printTreeHelper(node *Node, level int) {
    if node == t.NIL {
        return
    }
    t.printTreeHelper(node.Right, level+1)
    for i := 0; i < level; i++ {
        fmt.Print("    ")
    }
    color := "R"
    if node.Color == BLACK {
        color = "B"
    }
    fmt.Printf("%d%s\n", node.Key, color)
    t.printTreeHelper(node.Left, level+1)
}
