package main

import (
	"fmt"
)

// Notes about binary search tree

// First node in a tree is called the root
// parent node is a node that has 1 or 2 children
// children nodes are nodes that have parents above them.
// nodes with no children are called leaves.
// The reason they called it a binary tree is becuase it can no more 1 or 2 children.
//  Left side is always smaller then its parent
// Right side is always greater then its parent.
// Nodes are always inserted as a leaf node.

// Speed
// Best & average time is O(logn) & worst case is O(n)

// Traversal
// Inorder (left, root, right)
// Preorder (Root, Left, Right)
// Postorder (Left, Right Root)
// Levelorder (print root, left, right)

//  Binary Search Tree
type BinaryTree struct {
	Root *TreeNode
}

// Insert
func (bt *BinaryTree) Insert(insertMe int) {
	if bt.Root == nil {
		bt.Root = &TreeNode{Value: insertMe}
		fmt.Println("successfully inserted root node")
		return
	}

	bt.Root.Insert(insertMe)
	return
}

// Delete
func (bt *BinaryTree) Delete(deleteMe int) {
	if bt.Root == nil {
		fmt.Println("Value not found...")
		return
	}
	bt.Root.Delete(deleteMe)
}

// Search
func (bt *BinaryTree) Search(findMe int) bool {
	if bt.Root == nil {
		fmt.Println("Value not found...")
		return false
	}
	return bt.Root.Search(findMe)
}

// Print Inorder
func (bt *BinaryTree) Inorder() {
	if bt.Root == nil {
		fmt.Println("Binary search tree is empty.")
		return
	}
	bt.Root.Inorder()
	return
}

// Print Preorder
func (bt *BinaryTree) Preorder() {
	if bt.Root == nil {
		fmt.Println("Binary search tree is empty.")
		return
	}
	bt.Root.Preorder()
}

// Print Postorder
func (bt *BinaryTree) Postorder() {
	if bt.Root == nil {
		fmt.Println("Binary search tree is empty.")
		return
	}
	bt.Root.Postorder()
}

// Get Inorder Array
func (bt *BinaryTree) GetInorderArray() []int {
	if bt.Root == nil {
		fmt.Println("Binary search tree is empty.")
		return nil
	}
	return bt.Root.GetInorderArray()
}

// Get Preorder Array
func (bt *BinaryTree) GetPreorderArray() []int {
	if bt.Root == nil {
		fmt.Println("Binary search tree is empty.")
		return nil
	}
	return bt.Root.GetPreorderArray()
}

// Get Postorder Array
func (bt *BinaryTree) GetPostorderArray() []int {
	if bt.Root == nil {
		fmt.Println("Binary search tree is empty.")
		return nil
	}
	return bt.Root.GetPostorderArray()
}

// Tree Node
type TreeNode struct {
	LeftChild  *TreeNode
	RightChild *TreeNode
	Value      int
}

// Insert
func (tn *TreeNode) Insert(insertMe int) {
	if insertMe == tn.Value {
		fmt.Println("node already, not inserting...")
		return
	}
	if insertMe < tn.Value {
		if tn.LeftChild == nil {
			tn.LeftChild = &TreeNode{Value: insertMe}
		} else {
			tn.LeftChild.Insert(insertMe)
		}
	} else {
		if tn.RightChild == nil {
			tn.RightChild = &TreeNode{Value: insertMe}
		} else {
			tn.RightChild.Insert(insertMe)
		}
	}
}

// Delete
func (tn *TreeNode) Delete(deleteMe int) *TreeNode {
	if tn == nil {
		return nil
	}

	if deleteMe < tn.Value {
		tn.LeftChild = tn.LeftChild.Delete(deleteMe)
	} else if deleteMe > tn.Value {
		tn.RightChild = tn.RightChild.Delete(deleteMe)
	} else {
		// if key is same a current tree node
		if tn.LeftChild == nil {
			return tn.RightChild
		} else if tn.RightChild == nil {
			return tn.LeftChild
		} else { // there are two children
			// Set value
			tn.Value = minValue(tn.RightChild)
			// Delete node
			tn.RightChild = tn.RightChild.Delete(tn.Value)
		}
	}

	return tn
}

func minValue(node *TreeNode) int {
	minVal := node.Value
	for node.LeftChild != nil {
		minVal = node.LeftChild.Value
		node = node.LeftChild
	}
	return minVal
}

// Inorder (Left, Root, Right)
func (tn *TreeNode) Inorder() {
	if tn.LeftChild != nil {
		tn.LeftChild.Inorder()
	}

	fmt.Println(tn.Value)

	if tn.RightChild != nil {
		tn.RightChild.Inorder()
	}
}

// Preorder (Root, Left, Right)
func (tn *TreeNode) Preorder() {
	fmt.Println(tn.Value)
	if tn.LeftChild != nil {
		tn.LeftChild.Preorder()
	}
	if tn.RightChild != nil {
		tn.RightChild.Preorder()
	}
}

// Postorder (Left, Root, Right)
func (tn *TreeNode) Postorder() {
	if tn.LeftChild != nil {
		tn.LeftChild.Postorder()
	}
	if tn.RightChild != nil {
		tn.RightChild.Postorder()
	}
	fmt.Println(tn.Value)
}

// Returns an array of the BST inorder
func (tn *TreeNode) GetInorderArray() []int {
	returnMe := []int{}
	if tn.LeftChild != nil {
		returnMe = append(returnMe, tn.LeftChild.GetInorderArray()...)
	}

	returnMe = append(returnMe, tn.Value)

	if tn.RightChild != nil {
		returnMe = append(returnMe, tn.RightChild.GetInorderArray()...)
	}
	return returnMe
}

// Returns a array of the BST in preorder
func (tn *TreeNode) GetPreorderArray() []int {
	returnMe := []int{}
	returnMe = append(returnMe, tn.Value)

	if tn.LeftChild != nil {
		returnMe = append(returnMe, tn.LeftChild.GetPreorderArray()...)
	}

	if tn.RightChild != nil {
		returnMe = append(returnMe, tn.RightChild.GetPreorderArray()...)
	}
	return returnMe
}

// Returns a array of the BST in postorder
func (tn *TreeNode) GetPostorderArray() []int {
	returnMe := []int{}

	if tn.LeftChild != nil {
		returnMe = append(returnMe, tn.LeftChild.GetPostorderArray()...)
	}

	if tn.RightChild != nil {
		returnMe = append(returnMe, tn.RightChild.GetPostorderArray()...)
	}

	returnMe = append(returnMe, tn.Value)
	return returnMe
}

func (tn *TreeNode) Search(findMe int) bool {
	if findMe == tn.Value {
		return true
	}

	if findMe < tn.Value {
		if tn.LeftChild != nil {
			return tn.LeftChild.Search(findMe)
		}
	} else {
		if tn.RightChild != nil {
			return tn.RightChild.Search(findMe)
		}
	}

	return false
}

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) Insert(insertMe int) {
	if ll.Head == nil {
		ll.Head = &ListNode{Val: insertMe}
		return
	}
	ll.Head.Insert(insertMe)
}

func (ll *LinkedList) Print() {
	if ll.Head == nil {
		fmt.Println("The linked list is empty nothing to print.")
	}
	ll.Head.Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Insert(insertMe int) {
	current := ln
	for current.Next != nil {
		if current.Val == insertMe {
			return
		}
		current = current.Next
	}
	current.Next = &ListNode{Val: insertMe}
}

func (ln *ListNode) Print() {
	current := ln

	for current != nil {
		fmt.Printf("%d\t", current.Val)
		current = current.Next
	}
	fmt.Println("\n")
	return
}

func ConvertLinkedListToBalancedBST(ll *LinkedList) *BinaryTree {
	if ll.Head == nil {
		fmt.Println("Linked List empty nothing to print")
		return nil
	}

	root := convertLinkedList(ll.Head, nil)
	return &BinaryTree{
		Root: root,
	}
}

func convertLinkedList(head, end *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == end {
		return &TreeNode{Value: head.Val}
	}

	mid := findMid(head, end)
	return &TreeNode{
		Value:      mid.Val,
		LeftChild:  convertLinkedList(head, mid),
		RightChild: convertLinkedList(mid.Next, end),
	}
}

func findMid(head, end *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func main() {
	bst := &BinaryTree{}
	// Insert
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(6)
	bst.Insert(11)

	// Print Inorder should print [3,5,6,10,11]
	bst.Inorder()

	// Get in order array should hold [3,5,6,10,11] in that order
	inorderArray := bst.GetInorderArray()
	fmt.Println(inorderArray)
	fmt.Println("====================================")

	// Look up item to see if it exist in the binary search tree
	// if exist := bst.Search(10); !exist {
	// 	fmt.Println("The element you searched for doesn't exist")
	// } else {
	// 	fmt.Println("The element you searched for does exist")
	// }

	// if exist := bst.Search(6); !exist {
	// 	fmt.Println("The element you search for doesn't exist")
	// } else {
	// 	fmt.Println("The element you search for does exist.")
	// }

	// if exist := bst.Search(7); !exist {
	// 	fmt.Println("The element you searched for doesn't exist")
	// } else {
	// 	fmt.Println("The element you searched for does exist")
	// }

	// Printing Preorder
	bst.Preorder()
	fmt.Println(bst.GetPreorderArray())
	fmt.Println("====================================")

	// Printing post order
	bst.Postorder()
	fmt.Println(bst.GetPostorderArray())

	// Inserting into linked list.
	ll := &LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	ll.Insert(6)
	ll.Insert(7)

	// Convert linked list into a balanced binary search tree
	balancedTree := ConvertLinkedListToBalancedBST(ll)
	balancedTree.Preorder()
}
