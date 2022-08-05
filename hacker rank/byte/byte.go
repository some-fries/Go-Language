package main

import "fmt"

func Power(num, numPow int32) int32 {
	newNum := int32(1)
	for i := 0; i < int(numPow); i++ {
		newNum = newNum * num
	}
	return newNum
}

func countBits(num uint32) int32 {
	counter := int32(0)
	newNum := int32(num)
	power := int32(0)
	val := int32(1)
	// otherVal := int32(1)

	for {
		if int32(num) < val {
			break
		}
		//newNum = newNum - val
		val = val * 2
		power += 1
	}
	newNum = int32(num)
	for {
		maxPower := Power(2, power-1)
		if newNum <= 0 {
			return counter
		} else { //if maxPower >= newNum {
			//fmt.Println(counter)
			newNum -= maxPower
			counter += 1
			power -= 1
		}
	}

}

func main() {
	/*num := 128
	power := 0
	val := 1
	for {
		if num < val {
			println(power)
			break
		}
		num = num - val
		val = val * 2
		power += 1

	}*/
	//fmt.Println(Power(2, 3))
	fmt.Println(countBits(uint32(25)))
}
