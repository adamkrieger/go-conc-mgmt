package main

import "fmt"

type Worker interface {
	DoWork() error
	Description() string
}

type workerImpl struct {
	description string
}

func (wi *workerImpl) DoWork() error {
	//TODO
	return nil
}

func (wi *workerImpl) Description() string {
	return wi.description
}

func main() {

	// arrOf := []string{"one", "two", "three"}

	// mapOf := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// }

	//workerImpls
	wi1 := workerImpl{description: "1"}
	wi2 := workerImpl{description: "2"}
	wi3 := workerImpl{description: "3"}

	wi1p := &wi1
	wi2p := &wi2
	wi3p := &wi3

	//Array, Slice
	{
		workerArray := []Worker{wi1p, wi2p, wi3p}
		// v--Array       Slice--v
		workerArray = workerArray[1 : len(workerArray)-1]

		for wIndex, worker := range workerArray {
			fmt.Println(workerArray[wIndex].Description())
			_ = worker.DoWork()
		}
	}

	//Stack
	{
		workerStack := []Worker{}
		//Push
		workerStack = append(workerStack, wi1p)
		//Pop
		next := workerStack[len(workerStack)-1]
		workerStack = workerStack[:len(workerStack)-1]
		fmt.Println(next.Description())
	}

	//Queue
	{
		workerQueue := []Worker{}
		//Enqueue
		workerQueue = append(workerQueue, wi1p)
		//Dequeue
		next := workerQueue[0]
		workerQueue = workerQueue[1:]
		fmt.Println(next.Description())
	}

	//Map
	{
		workerMap := map[string]Worker{
			"w1": wi1p,
			"w2": wi2p,
			"w3": wi3p,
		}

		if w, exists := workerMap["w2"]; exists {
			fmt.Println("Yep... ", w.Description())
		}
	}

	//Set
	{
		workerSet := map[Worker]interface{}{
			wi1p: nil,
			wi2p: nil,
		}

		//Delete
		delete(workerSet, wi1p)
		//Add
		workerSet[wi1p] = nil

		fmt.Println("Set len: ", len(workerSet))
	}
}

//Linked List
type workerLinkedList struct {
	head *workerLinkedListItem
}

type workerLinkedListItem struct {
	val  Worker
	next *workerLinkedListItem
	// For double
	prev *workerLinkedListItem
}

//Binary Tree
type workerBTree struct {
	root *workerNode
}

type workerNode struct {
	val   Worker
	left  *workerNode
	right *workerNode
}
