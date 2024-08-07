package datastruct

import (
	"fmt"
)

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 初始化节点
func NewTreeNode(data int) *TreeNode {
	return &TreeNode{
		Val:   data,
		Left:  nil,
		Right: nil,
	}
}

// // 初始化节点
// func NewTreeNode(data interface{}) *TreeNode {
// 	return &TreeNode{
// 		Val: data,
//         Left: nil,
//         Right: nil,
// 	}
// }

// PreOrderTraverseTreeNode 前序遍历
func PreOrderTraverseTreeNode(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}

	// 先打印当前节点值
	fmt.Printf("%v ", treeNode.Val)

	// 然后依次对左子树和右子树做前序遍历
	PreOrderTraverseTreeNode(treeNode.Left)
	PreOrderTraverseTreeNode(treeNode.Right)
}

// MidOrderTraverseTreeNode 中序遍历
func MidOrderTraverseTreeNode(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}

	// 先从左子树最左侧节点开始遍历
	MidOrderTraverseTreeNode(treeNode.Left)
	// 打印位于中间的根节点
	fmt.Printf("%v ", treeNode.Val)
	// 最后按照和左子树一样的逻辑遍历右子树
	MidOrderTraverseTreeNode(treeNode.Right)
}

// PostOrderTraverseTreeNode 后序遍历
func PostOrderTraverseTreeNode(treeNode *TreeNode) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}

	// 先遍历左子树
	PostOrderTraverseTreeNode(treeNode.Left)
	// 再遍历右子树
	PostOrderTraverseTreeNode(treeNode.Right)
	// 最后访问根节点
	fmt.Printf("%v ", treeNode.Val)
}

// LevelOrderTraverseTreeNode 层次遍历
func LevelOrderTraverseTreeNode(treeNode *TreeNode) {
	if nil == treeNode {
		return
	}

	tmp := []*TreeNode{treeNode}
	for i := 0; i < len(tmp); i++ {
		current := tmp[i]

		fmt.Printf("%v ", current.Val)

		if nil != current.Left {
			tmp = append(tmp, current.Left)
		}
		if nil != current.Right {
			tmp = append(tmp, current.Right)
		}
	}
}

// LevelOrderTraverseTreeNodeToArray 层次遍历二叉树，返回二维数组
func LevelOrderTraverseTreeNodeToArray(treeNode *TreeNode) [][]int {
	if treeNode == nil {
		return nil
	}
	res := make([][]int, 0)
	cur := []*TreeNode{treeNode}
	for len(cur) > 0 {
		// 收集当前层的所有值
		level := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)
		for _, node := range cur {
			level = append(level, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, level)
		cur = nextLevel
	}
	return res
}

var NULL = -1 << 63

// TransformArrayToTreeNode 将数组转化成二叉树
func TransformArrayToTreeNode(data []int) *TreeNode {
	n := len(data)
	if 0 == n {
		return nil
	}

	root := NewTreeNode(data[0])

	queue := make([]*TreeNode, 1)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && data[i] != NULL {
			node.Left = NewTreeNode(data[i])
			queue = append(queue, node.Left)
		}
		i++

		if i < n && data[i] != NULL {
			node.Right = NewTreeNode(data[i])
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
