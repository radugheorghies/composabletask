package main

import (
	"container/list"
	"log"
)

func main() {
	// default values to check
	log.Println("solution=", Solution([]int{3, 4, 5, 3, 7}))
	log.Println("solution=", Solution([]int{1, 2, 3, 4}))
	log.Println("solution=", Solution([]int{1, 3, 1, 2}))
	// all cases
	log.Println("solution=", Solution([]int{1, 3, 2, 4}))
	log.Println("solution=", Solution([]int{1, 3, 2, 1}))
	log.Println("solution=", Solution([]int{3, 5, 2, 1}))
	log.Println("solution=", Solution([]int{3, 5, 4, 3}))
	log.Println("solution=", Solution([]int{3, 5, 4, 4}))
	log.Println("solution=", Solution([]int{3, 5, 5, 4}))
	log.Println("solution=", Solution([]int{3, 4, 5, 4, 7}))
	log.Println("solution=", Solution([]int{7, 4, 5, 3, 4}))
	log.Println("solution=", Solution([]int{9, 2, 7, 5, 4, 3}))
	log.Println("solution=", Solution([]int{9, 2, 7, 5, 3, 4}))
	log.Println("solution=", Solution([]int{9, 2, 7, 5, 3, 8}))
	log.Println("solution=", Solution([]int{1, 9, 2, 8, 3, 7}))
	log.Println("solution=", Solution([]int{1, 9, 2, 4, 6, 7}))
	log.Println("solution=", Solution([]int{1, 9, 2, 4, 6, 2}))
	log.Println("solution=", Solution([]int{1, 9, 2, 4, 6, 5}))
	log.Println("solution=", Solution([]int{1, 9, 2, 4, 6, 7, 1}))
}

func Solution(A []int) int {
	log.Println("For the slice A:", A)

	var solution int
	var checkIfSafeToRemoveTheLast bool // it is used for a special condition related to the last tree cut (see Step 3.2)
	var direction string                // used for the same porpuse as checkIfSafeToRemoveTheLast (it will give the direction to check)

	// we will use a list to keep only the data that we nee to analize
	// other correct data is not important for us
	l := list.New()

	// check the basic requirements
	if len(A) < 4 || len(A) > 200 {
		log.Printf("N out of range [4..200]: %v\n", len(A))
		return -1
	}

	for i, v := range A {
		if v < 1 || v > 1000 {
			log.Printf("Element out of range [1..1000]: %v\n", v)
			return -1
		}

		l.PushBack(v)

		// cehck for removal of the precedent last tree contidion
		if checkIfSafeToRemoveTheLast {
			if direction == "down" {
				if l.Back().Value.(int) < l.Back().Prev().Value.(int) {
					// we need to stop because we have the last 4 with the same orientation
					log.Println("We have more than one cut. Wee need to exit.")
					return -1
				}
				if l.Back().Value.(int) > l.Back().Prev().Prev().Value.(int) {
					// we are ok
					solution++
				}
			} else {
				// direction is up
				if l.Back().Value.(int) > l.Back().Prev().Value.(int) {
					// we need to stop because we have the last 4 with the same orientation
					log.Println("We have more than one cut. Wee need to exit.")
					return -1
				}
				if l.Back().Value.(int) < l.Back().Prev().Prev().Value.(int) {
					// we are ok
					solution++
				}
			}
			checkIfSafeToRemoveTheLast = false
			// it is always safe to remove the middle tree
			l.Remove(l.Back().Prev().Prev())
			// we will also remove the first tree from the list, because now we have all conditions ok
			l.Remove(l.Front())
			continue
		}

		// we put the first, the second, the third and the fourth element inside the list and we continue,
		// because we don't have to do anything right now
		if l.Len() < 4 {
			continue
		}

		// it means that we already have at least 4 elements
		firstEl := l.Front().Value.(int)
		secondEl := l.Front().Next().Value.(int)
		thirdEl := l.Front().Next().Next().Value.(int)
		fourthEl := l.Front().Next().Next().Next().Value.(int)

		switch {
		case firstEl < secondEl:
			switch {
			case secondEl > thirdEl:
				switch {
				case thirdEl < fourthEl:
					// we are good, we can continue, we remove the first
					// element, because we have good conditions and we don't need it
					l.Remove(l.Front())
					continue
				case thirdEl == fourthEl:
					log.Println("Wrong dataset (2 adiacent trees equals).")
					return -1
				case thirdEl > fourthEl:
					// first we check if we did not made another cut until here
					if solution > 0 {
						log.Println("We have more than one cut. Wee need to exit.")
						return -1
					}
					// we need to remove one of the last 3 elemens, to have a good condition here
					// step 1: check if we can remove the first of the last 3 elements
					if firstEl < thirdEl {
						solution++
					}
					// step 2: it is always safe to cut the treee in the middle
					solution++

					// step 3: check if we can remove the last tree
					// we have 2 situations here:

					// step 3.1: this is the last elemen from A and it is safe to be cut
					if i == len(A)-1 {
						solution++
						return solution
					}

					// step 3.2: we need to load another element to cehck the condition
					checkIfSafeToRemoveTheLast = true
					direction = "down"
				}
			case secondEl == thirdEl:
				log.Println("Wrong dataset (2 adiacent trees equals).")
				return -1
			case secondEl < thirdEl:
				switch {
				case thirdEl < fourthEl:
					log.Println("We have more than one cut. Wee need to exit.")
					return -1
				case thirdEl == fourthEl:
					log.Println("Wrong dataset (2 adiacent trees equals).")
					return -1
				case thirdEl > fourthEl:
					if solution > 0 {
						log.Println("We have more than one cut. Wee need to exit.")
						return -1
					}
					solution = 2
					if l.Back().Value.(int) < l.Back().Prev().Prev().Value.(int) {
						solution++
					}
					l.Remove(l.Front())
					continue
				}
			}
		case firstEl == secondEl:
			log.Println("Wrong dataset (2 adiacent trees equals).")
			return -1
		case firstEl > secondEl:
			switch {
			case secondEl > thirdEl:
				switch {
				case thirdEl > fourthEl:
					log.Println("We have more than one cut. Wee need to exit.")
					return -1
				case thirdEl == fourthEl:
					log.Println("Wrong dataset (2 adiacent trees equals).")
					return -1
				case thirdEl < fourthEl:
					if solution > 0 {
						log.Println("We have more than one cut. Wee need to exit.")
						return -1
					}
					solution = 2
					if l.Back().Value.(int) < l.Back().Prev().Prev().Value.(int) {
						solution++
					}
					l.Remove(l.Front())
					continue
				}
			case secondEl == thirdEl:
				log.Println("Wrong dataset (2 adiacent trees equals).")
				return -1
			case secondEl < thirdEl:
				switch {
				case thirdEl > fourthEl:
					// we are good, we can continue, we remove the first
					// element, because we have good conditions and we don't need it
					l.Remove(l.Front())
					continue
				case thirdEl == fourthEl:
					log.Println("Wrong dataset (2 adiacent trees equals).")
					return -1
				case thirdEl < fourthEl:
					// first we check if we did not made another cut until here
					if solution > 0 {
						log.Println("We have more than one cut. Wee need to exit.")
						return -1
					}
					// we need to remove one of the last 3 elemens, to have a good condition here
					// step 1: check if we can remove the first of the last 3 elements
					if firstEl > thirdEl {
						solution++
					}
					// step 2 it is always safe to cut the treee in the middle
					solution++

					// step 3: check if we can remove the last tree
					// we have 2 situations here:

					// step 3.1: this is the last elemen from A and it is safe to be cut
					if i == len(A)-1 {
						solution++
						return solution
					}

					// step 3.2: we need to load another element to cehck the condition
					checkIfSafeToRemoveTheLast = true
					direction = "up"
				}
			}
		}
	}

	return solution
}
