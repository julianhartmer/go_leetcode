/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var finalResult *ListNode = nil
	var resultTail *ListNode = finalResult
	if nil == l1 || nil == l2 {
		return nil
	}

	zeroNode := ListNode{
		Val:  0,
		Next: nil,
	}

	var hasCarry bool = false
	for &zeroNode != l1 || &zeroNode != l2 || hasCarry != false {
		// calculate value and carry
		val := l1.Val + l2.Val
		if hasCarry {
			val += 1
		}
		hasCarry = val > 9
		if hasCarry {
			val -= 10
		}

		// create and add node
		nodeResult := &ListNode{
			Val: val,
		}

		if nil == finalResult {
			finalResult = nodeResult
			resultTail = finalResult
		} else {
			resultTail.Next = nodeResult
			resultTail = resultTail.Next
		}

		// prepare next value: When list is empty, use NullNode
		if nil == l1.Next {
			l1 = &zeroNode
		} else {
			l1 = l1.Next
		}
		if nil == l2.Next {
			l2 = &zeroNode
		} else {
			l2 = l2.Next
		}
	}

	return finalResult
}
