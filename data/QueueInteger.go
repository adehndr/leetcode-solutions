package data

type QueueInteger struct {
	data []int
}

func (q *QueueInteger) EnqueueInteger(input int) {
	(*q).data = append((*q).data, input)
}

func (q *QueueInteger) DequeueInteger() int {
	if q.IsEmptyInteger() {
		return 0
	}

	temp := (*q).data[0]
	(*q).data = (*q).data[1:]
	return temp
}

func (q *QueueInteger) Contains(n int) bool {
	for _, v := range q.data {
		if v == n {
			return true
		}
	}
	return false
}

func (q *QueueInteger) IsEmptyInteger() bool {
	return len((*q).data) == 0
}

func (q *QueueInteger) Length() int {
	return len((*q).data)
}