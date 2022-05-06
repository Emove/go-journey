package _33_number_of_recent_calls

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{queue: []int{}}
}

func (counter *RecentCounter) Ping(t int) int {
	counter.queue = append(counter.queue, t)

	for i := 0; i < len(counter.queue); i++ {
		if counter.queue[i] >= t-3000 {
			counter.queue = counter.queue[i:]
			break
		}
	}
	return len(counter.queue)
}
