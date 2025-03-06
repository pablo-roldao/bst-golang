# Binary Search Tree (BST) in Go

This repository contains an implementation of a **Binary Search Tree (BST)** in Go, with methods for insertion, deletion, searching, traversal, and various utility functions. It also includes unit tests to ensure correct functionality.

## Features
- **Insertion (`Insert`)**: Adds a node to the tree while maintaining BST properties.
- **Deletion (`Remove`)**: Removes a node, handling cases where the node has zero, one, or two children.
- **Search (`Search`)**: Finds a node by its value.
- **Tree Traversals**:
  - **Pre-order (`ToPreOrderString`)**
  - **In-order (`ToInOrderString`)**
  - **Post-order (`ToPosOrderString`)**
  - **Reverse order (`ToReverseString`)**
- **Tree Properties**:
  - **Height of the tree (`GetHeight`)**
  - **Number of leaf nodes (`GetNumberOfLeaves`)**
  - **Find the largest node (`getGreaterNode`)**
  - **Find the path to a node (`GetPath`)**

## Usage
Import the `bst` package and use its methods:
```go
package main

import (
    "fmt"
    "bst"
)

func main() {
    var tree *bst.Node // Initialize as nil

    tree = tree.Insert(50)
    tree = tree.Insert(30)
    tree = tree.Insert(70)
    tree = tree.Insert(20)
    tree = tree.Insert(40)

    fmt.Println("Pre-order Traversal:", tree.ToPreOrderString())
    fmt.Println("Search for 30:", tree.Search(30) != nil)
    tree = tree.Remove(30)
    fmt.Println("After removing 30:", tree.ToPreOrderString())
}
```

## Running Tests
To run the unit tests:
```sh
go test
```

### Test Cases
The test suite (`bst_test.go`) includes:
- **Insertion Test**: Verifies correct tree structure after inserting elements.
- **Deletion Tests**:
  - Removing a leaf node
  - Removing a node with one child
  - Removing a node with two children
  - Removing the root
- **Search Tests**: Checks if searching works correctly.
- **Traversal Tests**: Ensures all traversal methods return correct outputs.

### Sample Test Output
```
=== RUN   TestInsert
--- PASS: TestInsert (0.00s)
=== RUN   TestToInOrderString
--- PASS: TestToInOrderString (0.00s)
=== RUN   TestToPosOrderString
--- PASS: TestToPosOrderString (0.00s)
=== RUN   TestToReverseString
--- PASS: TestToReverseString (0.00s)
=== RUN   TestGetHeight
--- PASS: TestGetHeight (0.00s)
=== RUN   TestGetNumberOfLeaves
--- PASS: TestGetNumberOfLeaves (0.00s)
=== RUN   TestGetPath
--- PASS: TestGetPath (0.00s)
=== RUN   TestRemoveLeafNode
--- PASS: TestRemoveLeafNode (0.00s)
=== RUN   TestRemoveNodeWithOneLeftChild
--- PASS: TestRemoveNodeWithOneLeftChild (0.00s)
=== RUN   TestRemoveNodeWithOneRightChild
--- PASS: TestRemoveNodeWithOneRightChild (0.00s)
=== RUN   TestRemoveNodeWithTwoChildren
--- PASS: TestRemoveNodeWithTwoChildren (0.00s)
=== RUN   TestRemoveRootNode
--- PASS: TestRemoveRootNode (0.00s)
=== RUN   TestRemoveNonExistentNode
--- PASS: TestRemoveNonExistentNode (0.00s)
PASS
ok   github.com/pablo-roldao/bst  0.002s
```

## License
This project is licensed under the MIT License.

