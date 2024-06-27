package problem6

import "fmt"

func main() {
	maze := [][]int{
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 0, 0},
	}
	fmt.Println(hasExit(maze))
	maze = [][]int{
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 1, 0},
	}
	fmt.Println(hasExit(maze))
}

func hasExit(maze [][]int) bool {
	xpos, ypos := 0, 0
	traversed := make([][]bool, len(maze)) // initialize a slice of dy slices
	for i := 0; i < len(maze); i++ {
		traversed[i] = make([]bool, len(maze[0])) // initialize a slice of dx unit8 in each of dy slices
	}
	return exitHelper(xpos, ypos, maze, traversed)
}

func exitHelper(xpos, ypos int, maze [][]int, traversed [][]bool) bool {
	// if either index is out of bounds return false
	if ypos < 0 || xpos < 0 || ypos >= len(maze) || xpos >= len(maze[0]) {
		return false
	}
	// if current position = 1 return false
	if maze[ypos][xpos] == 1 {
		return false
	}
	// if currentn position is the end return true
	if ypos == len(maze)-1 && xpos == len(maze[0])-1 {
		return true
	}
	// if we have already visited this point, return false
	if traversed[ypos][xpos] {
		return false
	}

	//after all quickfail cases are passed, mark the current point as checked and go to the next point
	traversed[ypos][xpos] = true
	return exitHelper(xpos, ypos+1, maze, traversed) || exitHelper(xpos+1, ypos, maze, traversed) || exitHelper(xpos, ypos-1, maze, traversed) || exitHelper(xpos-1, ypos, maze, traversed)
}
