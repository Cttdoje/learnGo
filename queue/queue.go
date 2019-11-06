package queue

type Queue []interface{}

// 入栈
func (q *Queue) Push(value interface{}) {
	*q = append(*q, value)
}

//出栈
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//是否是空栈
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
