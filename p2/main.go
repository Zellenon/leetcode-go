//Solution passes all tests with 0ms runtime (meets or beats 100% of competitors) and 6.16MB memory
//(meets or beats 93.59% of competitors)
package p2

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	start := &ListNode{};
	current_node := start;
	next_digit := 0;

	for l1 != nil || l2 != nil || next_digit != 0 {
		x, y := 0, 0;
		if l1 != nil {
			x = l1.Val;
			l1 = l1.Next;
		}
		if l2 != nil {
			y = l2.Val;
			l2 = l2.Next;
		}

		sum := x + y + next_digit;
		next_digit = sum / 10;
		current_node.Next = &ListNode{Val: sum % 10};
		current_node = current_node.Next;
	}
	return start.Next;
}
