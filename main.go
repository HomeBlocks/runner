package runner

import "fmt"

type TaskStorage struct {
	TestField string
}

func main() {
	New()
}

func task1() {
	fmt.Println(1)
}

func task2() {
	fmt.Println(2)
}

func task3() {
	fmt.Println(3)
}

func task4() {
	fmt.Println(4)
}
