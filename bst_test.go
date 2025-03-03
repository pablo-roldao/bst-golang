package bst_test

import (
	"testing"

	"github.com/pablo-roldao/bst"
)

func TestInsert(t *testing.T) {
	var tree *bst.Node
	tree = tree.Insert(50)
	tree = tree.Insert(30)
	tree = tree.Insert(70)
	tree = tree.Insert(20)
	tree = tree.Insert(40)
	tree = tree.Insert(60)
	tree = tree.Insert(80)

	expected := "[50][30][20][40][70][60][80]"
	if tree.ToPreOrderString() != expected {
		t.Fatalf("%v(expected) != %v", expected, tree.ToPreOrderString())
	}
}
