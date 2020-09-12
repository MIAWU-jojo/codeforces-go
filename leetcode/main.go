package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"sort"
)

func main() {
	toBytes := func(g [][]string) [][]byte {
		n, m := len(g), len(g[0])
		bytes := make([][]byte, n)
		for i := range bytes {
			bytes[i] = make([]byte, m)
			for j := range bytes[i] {
				bytes[i][j] = g[i][j][0]
			}
		}
		return bytes
	}

	_ = MustBuildTreeNode

	_ = []interface{}{toBytes, ListNode{}, TreeNode{}}
}

// LC 39
func combinationSum(a []int, target int) (ans [][]int) {
	b := []int{}
	var f func(p, rest int)
	f = func(p, rest int) {
		if p == len(a) {
			return
		}
		if rest == 0 {
			ans = append(ans, append([]int(nil), b...))
			return
		}
		f(p+1, rest)
		if rest-a[p] >= 0 {
			b = append(b, a[p])
			f(p, rest-a[p])
			b = b[:len(b)-1]
		}
	}
	f(0, target)
	return
}

// LC 40
func combinationSum2(a []int, target int) (ans [][]int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	sort.Ints(a)
	var freq [][2]int
	for _, v := range a {
		if freq == nil || v != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{v, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var b []int
	var f func(p, rest int)
	f = func(p, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), b...))
			return
		}
		if p == len(freq) || rest < freq[p][0] {
			return
		}
		f(p+1, rest)
		most := min(rest/freq[p][0], freq[p][1])
		for i := 1; i <= most; i++ {
			b = append(b, freq[p][0])
			f(p+1, rest-i*freq[p][0])
		}
		b = b[:len(b)-most]
	}
	f(0, target)
	return
}

// LC 41
func firstMissingPositive(a []int) int {
	n := len(a)
	for i, v := range a {
		for 0 < v && v <= n && v != a[v-1] {
			a[i], a[v-1] = a[v-1], a[i]
			v = a[i]
		}
	}
	for i, v := range a {
		if i+1 != v {
			return i + 1
		}
	}
	return n + 1
}

// LC99
func recoverTree(root *TreeNode) {
	nodes := []*TreeNode{}
	var f func(o *TreeNode)
	f = func(o *TreeNode) {
		if o == nil {
			return
		}
		f(o.Left)
		nodes = append(nodes, o)
		f(o.Right)
	}
	f(root)
	so := make([]*TreeNode, len(nodes))
	copy(so, nodes)
	sort.Slice(so, func(i, j int) bool { return so[i].Val < so[j].Val })
	do := []*TreeNode{}
	for i, o := range nodes {
		if o.Val != so[i].Val {
			do = append(do, o)
		}
	}
	do[0].Val, do[1].Val = do[1].Val, do[0].Val
}

// LC 124
func maxPathSum(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := int(-1e18)
	var f func(*TreeNode) int
	f = func(o *TreeNode) int {
		if o == nil {
			return -1e18
		}
		l := max(f(o.Left), 0)
		r := max(f(o.Right), 0)
		ans = max(ans, o.Val+l+r)
		return o.Val + max(l, r)
	}
	f(root)
	return ans
}

// LC 216
func combinationSum3(k int, n int) (ans [][]int) {
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return
}

// LC 332
func findItinerary(tickets [][]string) []string {
	g := map[string][]string{}
	for _, p := range tickets {
		g[p[0]] = append(g[p[0]], p[1])
	}
	for _, vs := range g {
		sort.Strings(vs)
	}

	path := make([]string, 0, len(tickets)+1)
	var f func(string)
	f = func(v string) {
		for len(g[v]) > 0 {
			w := g[v][0]
			g[v] = g[v][1:]
			f(w)
		}
		path = append(path, v)
	}
	f("JFK")

	for i, j := 0, len(path)-1; i < j; i++ {
		path[i], path[j] = path[j], path[i]
		j--
	}
	return path
}
