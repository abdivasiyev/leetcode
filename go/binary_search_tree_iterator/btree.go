package binary_search_tree_iterator

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

const NullItem = -1

func NewTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    root := &TreeNode{Val: nums[0]}
    createTree(root, 0, nums)

    return root
}

func createTree(root *TreeNode, i int, nums []int) {
    if root == nil {
        return
    }
    if 2*i+1 < len(nums) {
        if nums[2*i+1] != NullItem {
            root.Left = &TreeNode{
                Val: nums[2*i+1],
            }
        }
        createTree(root.Left, 2*i+1, nums)
    }

    if 2*i+2 < len(nums) {
        if nums[2*i+2] != NullItem {
            root.Right = &TreeNode{
                Val: nums[2*i+2],
            }
        }
        createTree(root.Right, 2*i+2, nums)
    }

    return
}
