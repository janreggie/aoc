package structs

import "fmt"

// Based from my ProgrammingInterviews repository

// BinaryTree is an implementation of a binary tree
type BinaryTree struct {
	value int64
	left  *BinaryTree
	right *BinaryTree
}

// NewBinaryTree returns an empty binary tree
func NewBinaryTree(value int64) *BinaryTree {
	return &BinaryTree{value: value, left: nil, right: nil}
}

// Set sets the value of a binary tree
func (btree *BinaryTree) Set(value int64) {
	btree.value = value
}

// Get gets the value of a binary tree
func (btree *BinaryTree) Get() int64 {
	return btree.value
}

// SetLeft sets the left child of a binary tree
func (btree *BinaryTree) SetLeft(left *BinaryTree) {
	btree.left = left
}

// SetRight sets the right child of a binary tree
func (btree *BinaryTree) SetRight(right *BinaryTree) {
	btree.right = right
}

// GetLeft gets the left child of a binary tree
func (btree *BinaryTree) GetLeft() *BinaryTree {
	return btree.left
}

// GetRight gets the right child of a binary tree
func (btree *BinaryTree) GetRight() *BinaryTree {
	return btree.right
}

// Lookup looks for a particular value from a BinaryTree.
// It returns the *BinaryTree object with a particular value
func (btree *BinaryTree) Lookup(value int64) *BinaryTree {
	if value == btree.Get() {
		return btree
	}
	if value < btree.Get() && btree.GetLeft() != nil {
		return btree.GetLeft().Lookup(value)
	}
	if value > btree.Get() && btree.GetRight() != nil {
		return btree.GetRight().Lookup(value)
	}
	return nil // if no such value found!
}

// In returns true if value is in BinaryTree and false otherwise
func (btree *BinaryTree) In(value int64) bool {
	ptr := btree
	for ptr != nil {
		if current := ptr.Get(); value == current {
			return true
		} else if value < current {
			ptr = ptr.GetLeft()
		} else {
			ptr = ptr.GetRight()
		}
	}
	return false
}

// Insert inserts a value in a BinaryTree
func (btree *BinaryTree) Insert(value int64) {
	if value < btree.Get() {
		if left := btree.GetLeft(); left != nil {
			left.Insert(value)
		} else {
			btree.SetLeft(NewBinaryTree(value))
		}
	} else {
		if right := btree.GetRight(); right != nil {
			right.Insert(value)
		} else {
			btree.SetRight(NewBinaryTree(value))
		}
	}
}

// Update checks if a value is in the binary tree and returns false
// or inserts the value and returns true
func (btree *BinaryTree) Update(value int64) bool {
	ptr := btree
	for ptr != nil {
		if current := ptr.Get(); value == current {
			return false // no point in doing anything else
		} else if value > current {
			if right := ptr.GetRight(); right != nil {
				ptr = right
			} else {
				ptr.SetRight(NewBinaryTree(value))
				return true
			}
		} else {
			if left := ptr.GetLeft(); left != nil {
				ptr = left
			} else {
				ptr.SetLeft(NewBinaryTree(value))
				return true
			}
		}
	}
	return false
}

// Delete deletes a value from a BinaryTree
// and returns the new parent and an error if applicable
func (btree *BinaryTree) Delete(value int64) (*BinaryTree, error) {
	if btree == nil {
		return btree, fmt.Errorf("%v not found", value)
	}
	if current := btree.Get(); value == current {
		// we would want the leftest of the right branch to be the new parent
		if right := btree.GetRight(); right != nil {
			left := btree.GetLeft()
			leftest := right
			preleft := right
			for lefter := leftest.GetLeft(); lefter != nil; lefter = lefter.GetLeft() {
				preleft = leftest
				leftest = lefter
			}
			btree = leftest
			if preleft != leftest {
				preleft.SetLeft(nil)
				btree.SetRight(right)
			} else {
				btree.SetRight(nil)
			}
			btree.SetLeft(left)
		} else if left := btree.GetLeft(); left != nil {
			right := btree.GetRight()
			rightest := left
			preright := left
			for righter := rightest.GetRight(); righter != nil; righter = righter.GetRight() {
				preright = rightest
				rightest = righter
			}
			btree = rightest
			if preright != rightest {
				preright.SetRight(nil)
				btree.SetLeft(left)
			} else {
				btree.SetLeft(nil)
			}
			btree.SetRight(right)
		} else {
			btree = nil
		}
		return btree, nil
	} else if value < current {
		stree, err := btree.GetLeft().Delete(value)
		if err != nil {
			return btree, err
		}
		// if btree.GetLeft().Get() == value {
		// 	btree.SetLeft(nil)
		// } else {
		// 	btree.SetLeft(stree)
		// }
		if stree != btree.GetLeft() {
			// some deletion has taken place
			btree.SetLeft(stree)
		}
		return btree, nil
	}
	stree, err := btree.GetRight().Delete(value)
	if err != nil {
		return btree, err
	}
	if stree != btree.GetRight() {
		// some deletion has taken place
		btree.SetRight(stree)
	}
	return btree, nil
}
