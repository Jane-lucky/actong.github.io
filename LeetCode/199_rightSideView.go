package Alth

func BFSrightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		p := []*TreeNode{}
		size := len(q)
		res = append(res, q[size-1].Val)
		for j := 0; j < size; j++ {
			node := q[j]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}

		}
		q = p
	}
	return res
}

var res1 []int
var Maxlen int

func DFSrightSideView(root *TreeNode) []int {
	level := 0
	dfs(root, level)
	return res1
}

func dfs(root *TreeNode, lv int) {
	//统一深度遍历右子树

	if root == nil {
		return
	}
	if lv > Maxlen {
		res1 = append(res1, root.Val)
		Maxlen = lv
	}

	dfs(root.Right, lv+1)

	dfs(root.Left, lv+1)
}
