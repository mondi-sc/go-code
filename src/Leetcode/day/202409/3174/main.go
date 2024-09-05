package main

import "unicode"

/*
	给你一个字符串 s 。

	你的任务是重复以下操作删除 所有 数字字符：
	· 删除 第一个数字字符 以及它左边 最近 的 非数字 字符。
	请你返回删除所有数字字符以后剩下的字符串。
*/

func clearDigits(s string) string {
	st := []rune{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}