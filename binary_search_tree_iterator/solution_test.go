package binary_search_tree_iterator

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestBSTIterator_HasNext(t *testing.T) {
    testCases := []struct {
        name    string
        root    *TreeNode
        hasNext bool
    }{
        {
            name:    "has next item",
            root:    NewTree([]int{7, 3, 15, NullItem, NullItem, 9, 20}),
            hasNext: true,
        },
        {
            name:    "has not next item",
            root:    NewTree([]int{}),
            hasNext: false,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            iterator := Constructor(tc.root)

            got := iterator.HasNext()

            assert.Equal(t, tc.hasNext, got)
        })
    }
}

func TestBSTIterator_Next(t *testing.T) {
    testCases := []struct {
        name     string
        root     *TreeNode
        commands []string
        output   []interface{}
    }{
        {
            name:     "example 1",
            root:     NewTree([]int{7, 3, 15, NullItem, NullItem, 9, 20}),
            commands: []string{"BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"},
            output:   []interface{}{NullItem, 3, 7, true, 9, true, 15, true, 20, false},
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            var iter BSTIterator

            for i, command := range tc.commands {
                switch command {
                case "BSTIterator":
                    iter = Constructor(tc.root)
                case "next":
                    got := iter.Next()
                    assert.Equal(t, tc.output[i], got)
                case "hasNext":
                    got := iter.HasNext()
                    assert.Equal(t, tc.output[i], got)
                }
            }
        })
    }
}
