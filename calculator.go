package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strconv"
)


func main() {		
var string_enter string

fmt.Fscan( os.Stdin, &string_enter) 



fmt.Println(string_enter)



if strings.Contains(string_enter, "+"){
	nums := strings.Split(string_enter, "+")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	fmt.Println(num1 + num2)
} else if strings.Contains(string_enter, "-"){
	nums := strings.Split(string_enter, "+")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	fmt.Println(num1 - num2)
} else if strings.Contains(string_enter, "*"){
	nums := strings.Split(string_enter, "*")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	fmt.Println(num1 * num2)
} else if strings.Contains(string_enter, "/"){
	nums := strings.Split(string_enter, "/")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	fmt.Println(num1 / num2)
}
}