package bst_test

import (
	"testing"

	"github.com/pablo-roldao/bst"
)

var tree *bst.Node

func TestInsert(t *testing.T) {
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

func TestToInOrderString(t *testing.T) {
	expected := "[20][30][40][50][60][70][80]"
	if tree.ToInOrderString() != expected {
		t.Fatalf("%v(expected) != %v", expected, tree.ToInOrderString())
	}
}

func TestToPosOrderString(t *testing.T) {
	expected := "[20][40][30][60][80][70][50]"
	if tree.ToPosOrderString() != expected {
		t.Fatalf("%v(expected) != %v", expected, tree.ToPosOrderString())
	}
}

func TestToReverseString(t *testing.T) {
	expected := "[80][70][60][50][40][30][20]"
	if tree.ToReverseString() != expected {
		t.Fatalf("%v(expected) != %v", expected, tree.ToReverseString())
	}
}

func TestGetHeight(t *testing.T) {
	expected := 3
	if tree.GetHeight() != expected {
		t.Fatalf("%v(expected) != %v", expected, tree.GetHeight())
	}
}

func TestGetNumberOfLeaves(t *testing.T) {
	expected := 4
	if tree.GetNumberOfLeaves() != expected {
		t.Fatalf("%v(expected) != %v", expected, tree.GetNumberOfLeaves())
	}
}

func TestGetPath(t *testing.T) {
	expected := "[50][70][60]"
	if tree.GetPath(60) != expected {
		t.Fatalf("%v(e60xpected) != %v", expected, tree.GetPath(60))
	}
}
