package data

type QueueString struct {
	data []string
}

func (q *QueueString) EnqueueString(input string) {
	(*q).data = append((*q).data, input)
}

func (q *QueueString) DequeueString() string {
	if q.IsEmptyString() {
		return ""
	}

	temp := (*q).data[0]
	(*q).data = (*q).data[1:]
	return temp
}

func (q *QueueString) IsEmptyString() bool {
	return len((*q).data) == 0
}
