package main

import (
  "fmt"
  // "strconv"
)

type numbers struct {
  number string
  index  int
}

func main() {
  str := "123456789"
  number := 2
  input := numbers{str, number}
  numbersArr := input.addToSlice()
  numberOfDigits := len(str) - number
  fmt.Println("gluin", gluin(numberOfDigits, numbersArr))
//   a := numbersArr[0:3]
// //   firstNumbers := a[0]+a[1]+a[2]
// //   fmt.Println(firstNumbers)
}

func (n numbers) addToSlice() []string {
  fmt.Println("Input number", n.number)
  numbersArr := make([]string, 0)
  for i := 0; i < len(n.number); i++ {
    stringValue := string(n.number[i])
    numbersArr = append(numbersArr, string(stringValue))
  }
  fmt.Println(numbersArr)

  return numbersArr
}
func  gluin(numberOfDigits int, a []string) string { //метод який складає слайс в один рядок
  var s string
  for i := 0; i < numberOfDigits; i++ {
    s+=a[i]
}
return s
}