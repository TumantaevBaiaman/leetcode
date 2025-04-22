package leetcode

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (stack *MyStack) Push(x int) {
	if len(stack.queue1) == 0 {
		stack.queue1 = append(stack.queue1, x)
		for len(stack.queue2) > 0 {
			stack.queue1 = append(stack.queue1, stack.queue2[0])
			stack.queue2 = stack.queue2[1:]
		}
	} else {
		stack.queue2 = append(stack.queue2, x)
		for len(stack.queue1) > 0 {
			stack.queue2 = append(stack.queue2, stack.queue1[0])
			stack.queue1 = stack.queue1[1:]
		}
	}
}

func (stack *MyStack) Pop() int {
	if stack.Empty() {
		return -1
	}
	if len(stack.queue1) > 0 {
		top := stack.queue1[0]
		stack.queue1 = stack.queue1[1:]
		return top
	}
	top := stack.queue2[0]
	stack.queue2 = stack.queue2[1:]
	return top
}

func (stack *MyStack) Top() int {
	if stack.Empty() {
		return -1
	}
	if len(stack.queue1) > 0 {
		return stack.queue1[0]
	}
	return stack.queue2[0]
}

func (stack *MyStack) Empty() bool {
	return len(stack.queue1) == 0 && len(stack.queue2) == 0
}
