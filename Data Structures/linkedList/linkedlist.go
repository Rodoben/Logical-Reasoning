package main

type Node struct {
	Data any
	next *Node
}

type linkedList struct {
	Head   *Node
	length int
}

func (l *linkedList) InsertNode(n *Node) {
	second := l.Head     // taking the address of next node
	l.Head = n           // storing the node to head of linkedList
	l.Head.next = second // storing the adress of second node to the node next
	l.length++           // increaing the size of the node
}
func (l *linkedList) DisplayLinkedList() {

}

func main() {
	l := linkedList{}
	n1 := &Node{Data: 1}

	l.InsertNode(n1)
}
