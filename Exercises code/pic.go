package main

//import "fmt"
//import "math"
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picSlice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		dxArray := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			dxArray[j] = uint8(2*i + j*j)
		}
		picSlice[i] = dxArray
	}
	return picSlice
}

func main() {
	pic.Show(Pic)
}
