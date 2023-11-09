package binary_search_tree_iterator

import (
    "reflect"
    "testing"
)

// [7, 3, 15, null, null, 9, 20]
func TestNewTree(t *testing.T) {
    tree := &TreeNode{
        Val: 7,
        Left: &TreeNode{
            Val: 3,
        },
        Right: &TreeNode{
            Val: 15,
            Left: &TreeNode{
                Val: 9,
            },
            Right: &TreeNode{
                Val: 20,
            },
        },
    }

    arr := []int{7, 3, 15, NullItem, NullItem, 9, 20}

    binaryTree := NewTree(arr)

    if !reflect.DeepEqual(binaryTree, tree) {
        t.Fatalf("constructed binary tree not equal to givent tree")
    }
}
