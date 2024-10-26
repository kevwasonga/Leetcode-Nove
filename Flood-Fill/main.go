//  733. Flood Fill   https://leetcode.com/problems/flood-fill/description/

package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	sr, sc, newColor := 1, 1, 2
	result := floodFill(grid, sr, sc, newColor)
	fmt.Println("Grid after flood fill:")
	for _, row := range result {
		fmt.Println(row)
	}
}

// floodFill applies a new color to all connected cells of the starting cell's color
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	targetColor := image[sr][sc]
	if targetColor == newColor {
		return image 
	}
	fill(image, sr, sc, targetColor, newColor)
	return image
}

// fill applies the new color recursively to adjacent cells with the target color
func fill(image [][]int, row int, col int, targetColor int, newColor int) {
	if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) || image[row][col] != targetColor {
		return
	}

	image[row][col] = newColor

	directions := []Direction{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}

	for _, dir := range directions {
		fill(image, row+dir.dy, col+dir.dx, targetColor, newColor)
	}
}

// Direction struct 
type Direction struct {
	dy int
	dx int
}
