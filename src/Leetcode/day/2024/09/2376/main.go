package main

import "strconv"

/*
	如果一个正整数每一个数位都是 互不相同 的，我们称它是 特殊整数 。

	给你一个 正 整数 n ，请你返回区间 [1, n] 之间特殊整数的数目。
*/

/*
	前置知识：位运算与集合论
	集合可以用二进制表示，二进制从低到高第 i 位为 1 表示 i 在集合中，为 0 表示 i 不在集合中。例如集合 {0,2,3} 对应的二进制数为 1101(2) 。
	设集合对应的二进制数为 x。本题需要用到两个位运算操作：
		1. 判断元素 d 是否在集合中：x >> d & 1 可以取出 x 的第 d 个比特位，如果是 1 就说明 d 在集合中。
		2. 把元素 d 添加到集合中：将 x 更新为 x | (1 << d)。

	思路
	将 n 转换成字符串 s，定义 dfs(i,mask,isLimit,isNum) 表示构造第 i 位及其之后数位的合法方案数，其余参数的含义为：
	· mask 表示前面选过的数字集合，换句话说，第 i 位要选的数字不能在 mask 中。
	· isLimit 表示当前是否受到了 n 的约束（注意要构造的数字不能超过 n）。若为真，则第 i 位填入的数字至多为 s[i]，否则可以是 9。如果在受到约束
	  的情况下填了 s[i]，那么后续填入的数字仍会受到 n 的约束。例如 n=123，如果 i=0 填的是 1 的话，i=1 的这一位至多填 2。如果 i=0 填的是 1，
	  i=1 填的是 2，那么 i=2 的这一位至多填 3。
	· isNum 表示 i 前面的数位是否填了数字。若为假，则当前位可以跳过（不填数字），或者要填入的数字至少为 1；若为真，则要填入的数字可以从 0 开始。
	  例如 n=123，在 i=0 时跳过的话，相当于后面要构造的是一个 99 以内的数字了，如果 i=1 不跳过，那么相当于构造一个 10 到 99 的两位数，如果
	  i=1 跳过，相当于构造的是一个 9 以内的数字。
	· 为什么要定义 isNum？因为 010 和 10 都是 10，如果认为第一个 0 和第三个 0 都是我们填入的数字，这就不符合题目要求了，但 10 显然是符合题目要求的。

	实现细节
	递归入口：dfs(0,0,true,false)，表示：
	· 从 s[0] 开始枚举；
	· 一开始集合中没有数字（空集）；
	· 一开始要受到 n 的约束（否则就可以随意填了，这肯定不行）；
	· 一开始没有填数字。

	递归中：
	· 如果 isNum 为假，说明前面没有填数字，那么当前也可以不填数字。一旦从这里递归下去，isLimit 就可以置为 false 了，这是因为 s[0] 必然是大于 0 的
	  ，后面就不受到 n 的约束了。或者说，最高位不填数字，后面无论怎么填都比 n 小。
	· 如果 isNum 为真，那么当前必须填一个数字。枚举填入的数字，根据 isNum 和 isLimit 来决定填入数字的范围。

	递归终点：当 i 等于 s 长度时，如果 isNum 为真，则表示得到了一个合法数字（因为不合法的数字不会递归到终点），返回 1，否则返回 0。

*/

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int, bool, bool) int
	dfs = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			return
		}
		if !isLimit && isNum {
			p := &memo[i][mask]
			if *p != -1 { // 之前计算过
				return *p
			}
			defer func() { *p = res }() // 记忆化
		}
		if !isNum { // 可以跳过当前数位
			res += dfs(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中，说明之前没有填过 d
				res += dfs(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return dfs(0, 0, true, false)
}