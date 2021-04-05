package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var filename string = "maze.txt"

func main() {
	//寫迷宮時 上左下右 逆時針順序探索 ，同時搜索最短路線，不走已經走過的位置
	//廣度優先算法，把所有近距離路線都探索過(先上下左右才往深處)

	fmt.Println(readMaze())
	//start = 0 0
	maze := readMaze()
	startP := point{0, 0}
	endP := point{5, 6}
	steps := walk(maze, startP, endP)
	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%d\t", steps[i][j])
		}
		fmt.Println()
	}
	// fmt.Println(steps)
}

type point struct {
	x, y int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) Add(pp point) point {
	p.x = p.x + pp.x
	p.y = p.y + pp.y
	return point{p.x, p.y}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.x < 0 || p.x >= len(grid) {
		return 0, false
	}
	if p.y < 0 || p.y >= len(grid[p.x]) {
		return 0, false
	}
	return grid[p.x][p.y], true
}

func walk(maze [][]int, startP, endP point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{startP} //剩餘需要移動的位置
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:] //取第一個值用
		for _, dir := range dirs {
			next := cur.Add(dir)

			if cur == endP { //到終點就結束
				break
			} else if val, ok := next.at(maze); !ok || val == 1 { // maze at next is 0 = 不可探索的空間
				continue
			} else if val, ok := next.at(steps); !ok || val != 0 { // and steps at next is 0 = 探索過的地區
				continue
			} else if next == startP { // next != start 不重覆起點
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.x][next.y] = curSteps + 1
			fmt.Printf("step %d: [%d %d]\n", curSteps+1, next.x, next.y)
			Q = append(Q, next)
		}
	}

	return steps
}

func readMaze() [][]int {
	file, err := ioutil.ReadFile(filename) //內部已有open / close 不需要再close
	HandleErr(err)
	fmt.Println(file)
	fileStr := string(file)
	var res [][]int
	// var rows, cols int
	for i, v := range strings.Split(fileStr, "\n") {
		// 第一行 設定檔
		rowData := strings.Split(strings.TrimSpace(v), " ")
		if i == 0 {
			if len(rowData) < 2 {
				return res
			}
			// else {
			// cols, err = strconv.Atoi(rowData[0])
			// HandleErr(err)
			// rows, err = strconv.Atoi(rowData[1])
			// HandleErr(err)
			// }
		} else {
			var row []int
			for _, innerV := range rowData {
				data, err := strconv.Atoi(innerV)
				HandleErr(err)
				row = append(row, data)
			}
			res = append(res, row)
		}
	}
	return res
}

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}
