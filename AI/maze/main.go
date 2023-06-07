package main

import "fmt"

func main() {
	// maze is a 2D array representing the maze
	// 0 = empty space, 1 = obstacle, R = rat, T = target
	var maze = [][]byte{
		{'R', '0', '1', '1', '0', '1'},
		{'1', '0', '0', '1', '0', '0'},
		{'0', '0', '1', '0', '0', '0'},
		{'1', '0', '1', '0', '1', '0'},
		{'1', '0', '0', '0', 'T', '0'},
	}

	// visited tracks the cells that have been visited
	var visited = make([][]bool, len(maze))

	// queue is used for breadth-first search
	var queue [][]int

	// startRow and startCol represent the starting position of the rat
	var startRow, startCol int

	// targetRow and targetCol represent the target position
	var targetRow, targetCol int

	// Initialize visited and queue
	for i := 0; i < len(maze); i++ {
		visited[i] = make([]bool, len(maze[0]))
		queue = append(queue, []int{i, 0})
	}

	// Find the starting and target positions
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[0]); j++ {
			if maze[i][j] == 'R' {
				startRow = i
				startCol = j
			} else if maze[i][j] == 'T' {
				targetRow = i
				targetCol = j
			}
		}
	}

	// Perform breadth-first search
	queue = append(queue, []int{startRow, startCol})
	visited[startRow][startCol] = true

	// distances will store the number of steps required to reach each cell
	var distances = make([][]int, len(maze))
	for i := 0; i < len(maze); i++ {
		distances[i] = make([]int, len(maze[0]))
	}

	// directions is a list of possible movements
	var directions = [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	for len(queue) > 0 {
		// Get the next cell to visit
		current := queue[0]
		queue = queue[1:]

		// Check if we've reached the target
		if current[0] == targetRow && current[1] == targetCol {
			break
		}

		// Visit all adjacent cells
		for _, direction := range directions {
			row := current[0] + direction[0]
			col := current[1] + direction[1]

			// Check if the adjacent cell is valid and unvisited
			if row >= 0 && row < len(maze) && col >= 0 && col < len(maze[0]) && !visited[row][col] && maze[row][col] != '1' {
				queue = append(queue, []int{row, col})
				visited[row][col] = true
				distances[row][col] = distances[current[0]][current[1]] + 1
			}
		}
	}
	// Print the distances to each cell
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[0]); j++ {
			fmt.Printf("%3d", distances[i][j])
		}
		fmt.Println()
	}
}
