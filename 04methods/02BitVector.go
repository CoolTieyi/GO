package main

type IntSet struct {
	words []uint64
}

//Has 是否包含
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//Add 添加
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

//UnionWith 并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	/*
		GO中集合通常会用map[T]bool来表示 T代表元素类型
		但是我们有一种更好的形式来表现-----bit数组
		(如在数据流分析领域 很多元素 常用交集并集等)

		bit数组常用一个无符号数或称之为"字"的slice来表示
		每一个元素的每一位都表示集合里的一个值
		当集合的第i位被设置时 我们才说这个集合包含元素i
	*/
}
