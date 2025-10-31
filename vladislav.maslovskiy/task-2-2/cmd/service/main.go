package main

import (
	"container/heap"
	"fmt"

	intMaxHeap "github.com/vladislavmaslovskiy/task-2-2/internal/intmaxheap"
)

func findPreferredDish(dishes []int, targetPosition int) int {
	var selectedDish int

	dishesHeap := intMaxHeap.InitIntMaxHeap(dishes)

	for range targetPosition {
		dish, popped := heap.Pop(dishesHeap).(int)
		if popped {
			selectedDish = dish
		}
	}

	return selectedDish
}

func main() {
	var totalDishes int

	_, err := fmt.Scanln(&totalDishes)
	if err != nil || totalDishes < 1 || totalDishes > 10000 {
		fmt.Println("Invalid dishes count!")

		return
	}

	dishes := make([]int, totalDishes)

	for index := range totalDishes {
		_, err = fmt.Scan(&dishes[index])
		if err != nil {
			fmt.Println("Invalid dish description!")

			return
		}

		if dishes[index] < -10000 || dishes[index] > 10000 {
			fmt.Println("Unsupported dish description!")

			return
		}
	}

	var desiredPosition int

	_, err = fmt.Scanln(&desiredPosition)
	if err != nil || desiredPosition < 1 || desiredPosition > totalDishes {
		fmt.Println("Invalid number of a dish in buffet!")

		return
	}

	fmt.Println(findPreferredDish(dishes, desiredPosition))
}
