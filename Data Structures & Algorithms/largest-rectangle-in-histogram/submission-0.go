func largestRectangleArea(heights []int) int {
    maxArea := 0
    stack := make([][2]int, 0)

    for i, h := range heights {
        start := i
        for len(stack) > 0 && stack[len(stack)-1][1] > h {
            index := stack[len(stack)-1][0]
            height := stack[len(stack)-1][1]
            stack = stack[:len(stack)-1]
            area := height * (i - index)
            if area > maxArea {
                maxArea = area
            }
            start = index
        }
        stack = append(stack, [2]int{start, h})
    }

    n := len(heights)
    for _, pair := range stack {
        area := pair[1] * (n - pair[0])
        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}