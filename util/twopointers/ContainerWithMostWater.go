package twopointers

import "fmt"

func ContainerWithMostWater(heights []int) int {
	return containerWithMostWaterOptimizedMe(heights)
}

func containerWithMostWater1(height []int) int {
	finalArea := 0

	if len(height) <= 1 {
		return finalArea
	}

	left := 0
	right := 1

	for right < len(height) {
		if height[right] > height[left] {
			left = right
		} else {
			area := height[right] * (right - left)
			fmt.Println(left, right, area)
			if area > finalArea {
				finalArea = area
			}
		}
		right += 1
	}
	return finalArea
}

func containerWithMostWater2(height []int) int {
	finalResult := 0

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			var tempArea int
			if height[j] > height[i] {
				tempArea = height[i] * (j - i)
			} else {
				tempArea = height[j] * (j - i)
			}
			if tempArea >= finalResult {
				finalResult = tempArea
			}
		}
	}

	return finalResult
}

func containerWithMostWaterOptimizedMe(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		fmt.Println(height[left], height[right])
		area := findMin(height[left], height[right]) * (right - left)
		if area >= maxArea {
			maxArea = area
		}
		if height[left] <= height[right] {
			left = left + 1
		} else {
			right = right - 1
		}
	}

	return maxArea
}

func findMin(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}
