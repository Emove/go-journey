package _32_implement_queue_using_stack

type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{in: []int{}, out: []int{}}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
	if len(q.out) == 0 {
		q.in2out()
	}
	v := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return v
}

func (q *MyQueue) Peek() int {
	if len(q.out) == 0 {
		q.in2out()
	}
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func (q *MyQueue) in2out() {
	for len(q.in) > 0 {
		q.out = append(q.out, q.in[len(q.in)-1])
		q.in = q.in[:len(q.in)-1]
	}
}
