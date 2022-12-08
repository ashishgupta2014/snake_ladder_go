package helper

type Point struct {
	x int
	y int
}

func BuildPointerArray(number int, raw_array [][2]int) []Point {
	var output = make([]Point, number)
	for i := 0; i < number; i++ {
		var x = raw_array[i][0]
		var y = raw_array[i][1]
		output[i] = Point{x: x, y: y}
	}
	return output
}

func (point Point) GetX() int {
	return point.x
}

func (point Point) GetY() int {
	return point.y
}
