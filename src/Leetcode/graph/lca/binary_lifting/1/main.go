package main

import "math/bits"

// 树上倍增算法求lca（最近公共祖先）

type TreeAncestor [][]int

func Constructor(n int, parent []int) TreeAncestor {
	m := bits.Len(uint(n))
	pa := make([][]int, n)
	// pa[i][0] 父节点 pa[i][1] 爷爷节点
	for i, p := range parent {
		pa[i] = make([]int, m)
		pa[i][0] = p
	}

	for i := 0; i < m-1; i++ {
		for x := 0; x < n; x++ {
			if p := pa[x][i]; p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}
	return pa
}

func (pa *TreeAncestor) GetKthAncestor(node int, k int) int {
	m := bits.Len(uint(k))
	for i := 0; i < m; i++ {
		if k>>i&1 == 1 { // k 的二进制从低到高第 i 位是 1
			node = (*pa)[node][i]
			if node < 0 {
				break
			}
		}
	}
	return node
}

// GetKthAncestor2 另一种写法，不断去掉 k 的最低位的 1
func (pa TreeAncestor) GetKthAncestor2(node, k int) int {
	for ; k > 0 && node != -1; k &= k - 1 {
		node = pa[node][bits.TrailingZeros(uint(k))]
	}
	return node
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */

func main() {

}
