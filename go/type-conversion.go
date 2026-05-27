package main

import "fmt"

func main() {
	var tasksDone int = 7
	var totalTasks int = 10

	progress := float64(tasksDone) / float64(totalTasks) * 100

	fmt.Printf("Tasks done: %d\n", tasksDone)
	fmt.Printf("Total tasks: %d\n", totalTasks)
	fmt.Printf("Progress: %.1f%%\n", progress)

	wrongProgress := tasksDone / totalTasks * 100
	fmt.Printf("Wrong progress: %d%%\n", wrongProgress)

}
