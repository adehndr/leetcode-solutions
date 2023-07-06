package data

type QueueNodeBinaryTreeLeetCode struct {
	data []NodeBinarySearchTreeLeetCode
}

func (q *QueueNodeBinaryTreeLeetCode) EnqueueNodeBinaryTreeLeetCode(input NodeBinarySearchTreeLeetCode) {
	(*q).data = append((*q).data, input)
}

func (q *QueueNodeBinaryTreeLeetCode) DequeueNodeBinaryTreeLeetCode() *NodeBinarySearchTreeLeetCode {
	if q.IsEmptyNodeBinaryTreeLeetCode() {
		return nil
	}

	temp := (*q).data[0]
	(*q).data = (*q).data[1:]
	return &temp
}

func (q *QueueNodeBinaryTreeLeetCode) IsEmptyNodeBinaryTreeLeetCode() bool{
	return len((*q).data) == 0
}

func (q *QueueNodeBinaryTreeLeetCode) Length() int {
	return len((*q).data)
}