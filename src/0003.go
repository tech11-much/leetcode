package src

//滑动窗口加哈希表
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := map[byte]int{}
	rk, res := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			//当前指针已经移动到下一格了，所以需要减一，删除上一个字母
			delete(m, s[i-1])
		}
		for rk < n && m[s[rk]] == 0 {
			m[s[rk]]++
			rk++
		}
		res = max(res, rk-i)
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
