package main

import (
	"fmt"
	"strconv"
	// "strconv"
)

type numbers struct {
	number string
	index  int
}

func main() {
	str := "1234"
	number := 1
	input := numbers{str, number}
	numbersArr := input.addToSlice()//записуємо рядок в слайс
	numberOfDigits := len(str) - number //розрядність найбільшого можливого числа
	resoultArr := gluinNumbers(numberOfDigits, numbersArr)
	fmt.Println(findMax(resoultArr))
}
func findMax(resoultArr []string) int {
	var maxValue int
	for _, k := range resoultArr {
		value, _ := strconv.Atoi(k)
		if maxValue < value {
			maxValue = value
		}
	}
	return maxValue
}

func (n numbers) addToSlice() []string { //метод який додає всі цифри до слайсу
	fmt.Println("Input number", n.number)
	numbersArr := make([]string, 0, 12)
	for i := 0; i < len(n.number); i++ {
		stringValue := string(n.number[i])
		numbersArr = append(numbersArr, string(stringValue))
	}
	fmt.Println(numbersArr)
	return numbersArr
}

func gluinNumbers(numberOfDigits int, numbersArr []string) []string { //функція яка складає слайс в один рядок і повертає масив найбільших значень
	var s string
	resoultArr := make([]string, 0, 4)
	forCounter := 0
	for j := 0; j < len(numbersArr); j++ {
		for i := forCounter; i < numberOfDigits; i++ {
			s += numbersArr[i]
		}
		if numberOfDigits < len(numbersArr) {
			resoultArr = append(resoultArr, s)
			s = ""
			numberOfDigits++
			forCounter++
		} else {
			break
		}
	}
	resoultArr = append(resoultArr, s)
	return resoultArr
}
