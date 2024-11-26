package main

/*
	给你一个整数数组 colors ，它表示一个由红色和蓝色瓷砖组成的环，第 i 块瓷砖的颜色为 colors[i] ：
	· colors[i] == 0 表示第 i 块瓷砖的颜色是 红色 。
	· colors[i] == 1 表示第 i 块瓷砖的颜色是 蓝色 。
	环中连续 3 块瓷砖的颜色如果是 交替 颜色（也就是说中间瓷砖的颜色与它 左边 和 右边 的颜色都不同），那么它被称为一个 交替 组。
	请你返回 交替 组的数目。

	注意 ，由于 colors 表示一个 环 ，第一块 瓷砖和 最后一块 瓷砖是相邻的。
*/

func numberOfAlternatingGroups(colors []int) (ans int) {
	n := len(colors)
	cnt := 0
	for i := 0; i < n*2; i++ {
		if i > 0 && colors[i%n] == colors[(i-1)%n] {
			cnt = 0
		}
		cnt++
		if i >= n && cnt >= 3 {
			ans++
		}
	}
	return
}
