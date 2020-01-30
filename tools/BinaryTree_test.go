package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateBinaryTree() *BinaryTree {
	// generateBinaryTree() should create the following binary tree
	//             10
	//           /    \
	//          5       15
	//         / \    /    \
	//        3  7   13     17
	//       /\  /\  / \    / \
	//      1 4  6 8 11 14 16 18
	bt := NewBinaryTree(10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(13)
	bt.Insert(17)
	bt.Insert(1)
	bt.Insert(4)
	bt.Insert(6)
	bt.Insert(8)
	bt.Insert(11)
	bt.Insert(14)
	bt.Insert(16)
	bt.Insert(18)
	return bt
}

func TestGenerateBinaryTree(t *testing.T) {
	assert := assert.New(t)
	bt := generateBinaryTree()
	// is it actually equal?
	assert.EqualValues(bt.Get(), 10, "10")
	assert.EqualValues(bt.GetLeft().Get(), 5, "5")
	assert.EqualValues(bt.GetRight().Get(), 15, "15")
	assert.EqualValues(bt.GetLeft().GetLeft().Get(), 3, "3")
	assert.EqualValues(bt.GetLeft().GetRight().Get(), 7, "7")
	assert.EqualValues(bt.GetRight().GetLeft().Get(), 13, "13")
	assert.EqualValues(bt.GetRight().GetRight().Get(), 17, "17")
	assert.EqualValues(bt.GetLeft().GetLeft().GetLeft().Get(), 1, "1")
	assert.EqualValues(bt.GetLeft().GetLeft().GetRight().Get(), 4, "4")
	assert.EqualValues(bt.GetLeft().GetRight().GetLeft().Get(), 6, "6")
	assert.EqualValues(bt.GetLeft().GetRight().GetRight().Get(), 8, "8")
	assert.EqualValues(bt.GetRight().GetLeft().GetLeft().Get(), 11, "11")
	assert.EqualValues(bt.GetRight().GetLeft().GetRight().Get(), 14, "14")
	assert.EqualValues(bt.GetRight().GetRight().GetLeft().Get(), 16, "16")
	assert.EqualValues(bt.GetRight().GetRight().GetRight().Get(), 18, "18")
	assert.Nil(bt.GetLeft().GetLeft().GetLeft().GetLeft(), "leaf!")
	assert.Nil(bt.GetLeft().GetLeft().GetLeft().GetRight(), "leaf!")
	assert.Nil(bt.GetLeft().GetLeft().GetRight().GetLeft(), "leaf!")
	assert.Nil(bt.GetLeft().GetLeft().GetRight().GetRight(), "leaf!")
	assert.Nil(bt.GetLeft().GetRight().GetLeft().GetLeft(), "leaf!")
	assert.Nil(bt.GetLeft().GetRight().GetLeft().GetRight(), "leaf!")
	assert.Nil(bt.GetLeft().GetRight().GetRight().GetLeft(), "leaf!")
	assert.Nil(bt.GetLeft().GetRight().GetRight().GetRight(), "leaf!")
	assert.Nil(bt.GetRight().GetLeft().GetLeft().GetLeft(), "leaf!")
	assert.Nil(bt.GetRight().GetLeft().GetLeft().GetRight(), "leaf!")
	assert.Nil(bt.GetRight().GetLeft().GetRight().GetLeft(), "leaf!")
	assert.Nil(bt.GetRight().GetLeft().GetRight().GetRight(), "leaf!")
	assert.Nil(bt.GetRight().GetRight().GetLeft().GetLeft(), "leaf!")
	assert.Nil(bt.GetRight().GetRight().GetLeft().GetRight(), "leaf!")
	assert.Nil(bt.GetRight().GetRight().GetRight().GetLeft(), "leaf!")
	assert.Nil(bt.GetRight().GetRight().GetRight().GetRight(), "leaf!")
}

func TestNewBinaryTree(t *testing.T) {
	assert := assert.New(t)
	bt := NewBinaryTree(42)
	assert.EqualValues(bt.Get(), 42)
	assert.Nil(bt.GetLeft())
	assert.Nil(bt.GetRight())
}

func TestBinaryTree_Set(t *testing.T) {
	assert := assert.New(t)
	original := NewBinaryTree(42)
	original.Set(15)
	assert.EqualValues(original.Get(), 15)
}

func TestBinaryTree_SetLeft(t *testing.T) {
	assert := assert.New(t)
	original := NewBinaryTree(42)
	assist := NewBinaryTree(15)
	original.SetLeft(assist)
	assert.Nil(original.GetRight())
	assert.EqualValues(original.GetLeft(), assist)
}

func TestBinaryTree_SetRight(t *testing.T) {
	assert := assert.New(t)
	original := NewBinaryTree(42)
	assist := NewBinaryTree(15)
	original.SetRight(assist)
	assert.Nil(original.GetLeft())
	assert.EqualValues(original.GetRight(), assist)
}

func TestBinaryTree_Lookup(t *testing.T) {
	assert := assert.New(t)
	bt := generateBinaryTree()
	var lookedUp *BinaryTree

	// root?
	lookedUp = bt.Lookup(10)
	assert.NotNil(lookedUp)
	assert.EqualValues(lookedUp.Get(), 10, "root: 10")
	assert.EqualValues(lookedUp.GetLeft().Get(), 5, "root: 5")
	assert.EqualValues(lookedUp.GetRight().Get(), 15, "root: 15")

	// middle?
	lookedUp = bt.Lookup(13)
	assert.NotNil(lookedUp)
	assert.EqualValues(lookedUp.Get(), 13, "somewhere in the middle: 13")
	assert.EqualValues(lookedUp.GetLeft().Get(), 11, "somewhere in the middle: 11")
	assert.EqualValues(lookedUp.GetRight().Get(), 14, "somewhere in the middle: 14")

	// leaf?
	lookedUp = bt.Lookup(8)
	assert.NotNil(lookedUp)
	assert.EqualValues(lookedUp.Get(), 8, "leaf: 8")
	assert.Nil(lookedUp.GetLeft(), "leaf")
	assert.Nil(lookedUp.GetRight(), "leaf")

	// none
	lookedUp = bt.Lookup(9)
	assert.Nil(lookedUp, "can't find")
}

func TestBinaryTree_In(t *testing.T) {
	assert := assert.New(t)
	bt := generateBinaryTree()
	assert.True(bt.In(10))
	assert.True(bt.In(13))
	assert.True(bt.In(8))
	assert.False(bt.In(9))
}

func TestBinaryTree_Update(t *testing.T) {
	assert := assert.New(t)
	bt := NewBinaryTree(10)
	assert.True(bt.Update(15))
	assert.Nil(bt.GetLeft())
	assert.EqualValues(bt.GetRight().Get(), 15)
	assert.True(bt.Update(13))
	assert.Nil(bt.GetLeft())
	assert.EqualValues(bt.GetRight().GetLeft().Get(), 13)
}

func TestBinaryTree_Delete(t *testing.T) {
	assert := assert.New(t)
	var bt *BinaryTree
	var err error

	// note the original
	//             10
	//           /    \
	//          5       15
	//         / \    /    \
	//        3  7   13     17
	//       /\  /\  / \    / \
	//      1 4  6 8 11 14 16 18

	// delete node 10
	// new tree should be
	//             11
	//           /    \
	//          5       15
	//         / \    /    \
	//        3  7   13     17
	//       /\  /\  / \    / \
	//      1 4  6 8 - 14 16 18
	bt = generateBinaryTree()
	bt, err = bt.Delete(10)
	assert.NoError(err)
	assert.EqualValues(bt.Get(), 11)
	assert.EqualValues(bt.GetRight().Get(), 15)
	assert.EqualValues(bt.GetRight().GetLeft().Get(), 13)
	assert.Nil(bt.GetRight().GetLeft().GetLeft())
	assert.EqualValues(bt.GetRight().GetLeft().GetRight().Get(), 14)

	// delete node 8
	// new tree should be
	//             10
	//           /     \
	//          5       15
	//         / \    /    \
	//        3  7   13     17
	//       /\  /\  / \    / \
	//      1 4  6 - 11 14 16 18
	bt = generateBinaryTree()
	bt, err = bt.Delete(8)
	assert.NoError(err)
	assert.EqualValues(bt.Get(), 10)
	assert.EqualValues(bt.GetLeft().GetRight().Get(), 7)
	assert.EqualValues(bt.GetLeft().GetRight().GetLeft().Get(), 6)
	assert.Nil(bt.GetLeft().GetRight().GetRight())

	// delete node 13
	// new tree should be
	//             10
	//           /    \
	//          5       15
	//         / \    /    \
	//        3  7   14     17
	//       /\  /\  / \    / \
	//      1 4  6 8 11 -  16 18
	bt = generateBinaryTree()
	bt, err = bt.Delete(13)
	assert.NoError(err)
	assert.EqualValues(bt.Get(), 10)
	assert.EqualValues(bt.GetRight().Get(), 15)
	assert.EqualValues(bt.GetRight().GetLeft().Get(), 14)
	assert.EqualValues(bt.GetRight().GetLeft().GetLeft().Get(), 11)
	assert.Nil(bt.GetRight().GetLeft().GetRight())

	// delete node 12
	// this should be invalid
	bt = generateBinaryTree()
	bt, err = bt.Delete(12)
	assert.Error(err)
	assert.EqualValues(bt, generateBinaryTree())
}
