package avltree

// Node represents a node in the AVL tree
type Node struct {
    Key    int
    Height int
    Left   *Node
    Right  *Node
}

// AVLTree represents an AVL tree
type AVLTree struct {
    Root *Node
}

// NewAVLTree creates a new AVL tree
func NewAVLTree() *AVLTree {
    return &AVLTree{}
}

// Insert inserts a key into the AVL tree
func (t *AVLTree) Insert(key int) {
    t.Root = insert(t.Root, key)
}

// insert recursively inserts a key into the tree and balances it
func insert(node *Node, key int) *Node {
    if node == nil {
        return &Node{Key: key, Height: 1}
    }

    if key < node.Key {
        node.Left = insert(node.Left, key)
    } else if key > node.Key {
        node.Right = insert(node.Right, key)
    } else {
        // Duplicate keys are not allowed
        return node
    }

    node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))

    balance := getBalance(node)

    // Left Left Case
    if balance > 1 && key < node.Left.Key {
        return rightRotate(node)
    }

    // Right Right Case
    if balance < -1 && key > node.Right.Key {
        return leftRotate(node)
    }

    // Left Right Case
    if balance > 1 && key > node.Left.Key {
        node.Left = leftRotate(node.Left)
        return rightRotate(node)
    }

    // Right Left Case
    if balance < -1 && key < node.Right.Key {
        node.Right = rightRotate(node.Right)
        return leftRotate(node)
    }

    return node
}

// Delete removes a key from the AVL tree
func (t *AVLTree) Delete(key int) {
    t.Root = delete(t.Root, key)
}

// delete recursively deletes a key from the tree and balances it
func delete(node *Node, key int) *Node {
    if node == nil {
        return nil
    }

    if key < node.Key {
        node.Left = delete(node.Left, key)
    } else if key > node.Key {
        node.Right = delete(node.Right, key)
    } else {
        if node.Left == nil || node.Right == nil {
            var temp *Node
            if node.Left == nil {
                temp = node.Right
            } else {
                temp = node.Left
            }

            if temp == nil {
                temp = node
                node = nil
            } else {
                *node = *temp
            }
        } else {
            temp := minValueNode(node.Right)
            node.Key = temp.Key
            node.Right = delete(node.Right, temp.Key)
        }
    }

    if node == nil {
        return node
    }

    node.Height = 1 + max(getHeight(node.Left), getHeight(node.Right))

    balance := getBalance(node)

    // Left Left Case
    if balance > 1 && getBalance(node.Left) >= 0 {
        return rightRotate(node)
    }

    // Left Right Case
    if balance > 1 && getBalance(node.Left) < 0 {
        node.Left = leftRotate(node.Left)
        return rightRotate(node)
    }

    // Right Right Case
    if balance < -1 && getBalance(node.Right) <= 0 {
        return leftRotate(node)
    }

    // Right Left Case
    if balance < -1 && getBalance(node.Right) > 0 {
        node.Right = rightRotate(node.Right)
        return leftRotate(node)
    }

    return node
}

// Search searches for a key in the AVL tree
func (t *AVLTree) Search(key int) bool {
    return search(t.Root, key)
}

// search recursively searches for a key in the tree
func search(node *Node, key int) bool {
    if node == nil {
        return false
    }

    if key < node.Key {
        return search(node.Left, key)
    } else if key > node.Key {
        return search(node.Right, key)
    }

    return true
}

// Helper functions

func getHeight(node *Node) int {
    if node == nil {
        return 0
    }
    return node.Height
}

func getBalance(node *Node) int {
    if node == nil {
        return 0
    }
    return getHeight(node.Left) - getHeight(node.Right)
}

func rightRotate(y *Node) *Node {
    x := y.Left
    T2 := x.Right

    x.Right = y
    y.Left = T2

    y.Height = max(getHeight(y.Left), getHeight(y.Right)) + 1
    x.Height = max(getHeight(x.Left), getHeight(x.Right)) + 1

    return x
}

func leftRotate(x *Node) *Node {
    y := x.Right
    T2 := y.Left

    y.Left = x
    x.Right = T2

    x.Height = max(getHeight(x.Left), getHeight(x.Right)) + 1
    y.Height = max(getHeight(y.Left), getHeight(y.Right)) + 1

    return y
}

func minValueNode(node *Node) *Node {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
