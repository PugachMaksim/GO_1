package main

import (
	"fmt"
	// "os"
	"strconv"
	"strings"
	// "strconv"
)
var romeToArab = map[string]int{
	"I": 1,
	"II": 2,
	"III": 3,
	"IV": 4,
	"V": 5,
	"VI": 6,
	"VII": 7,
	"VIII": 8,
	"IX": 9,
	"X": 10,
}


var arabToRome = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
	10: "X",
	20: "XX",
	30: "XXX",
	40: "XL",
	50: "L",
	60: "LX",
	70: "LXX",
	80: "LXXX",
	90: "XC",
	100: "C",
}

func main() {		
	var string_enter string = "X+IX"
	// fmt.Fscan( os.Stdin, &string_enter) 

	// fmt.Println(string_enter)

	if strings.Contains(string_enter, "+"){
		nums := strings.Split(string_enter, "+")
		if value1, flag := romeToArab[nums[0]]; flag {
			if value2, flag := romeToArab[nums[1]]; flag {
				res := value1 + value2
				fmt.Println(Arab(res))
				fmt.Println(res)
			} else {
				panic("Неверный формат: Должны быть римские или арабские цифры")
			}
		} else {
				num1, err := strconv.Atoi(nums[0])
				if err != nil{
					panic("Введено не число...")
				}
				num2, err := strconv.Atoi(nums[1])
				if err != nil{
					panic("Введено не чило...")
				}
				fmt.Println(num1 + num2)
			} 
	} else if strings.Contains(string_enter, "-"){
		nums := strings.Split(string_enter, "+")
		if value1, flag := romeToArab[nums[0]]; flag {
			if value2, flag := romeToArab[nums[1]]; flag {
				fmt.Println(value1 - value2)
			} else {
				panic("Неверный формат: Должны быть римские или арабские цифры")
			}
		} else {
				num1, err := strconv.Atoi(nums[0])
				if err != nil{
					panic("Введено не число...")
				}
				num2, err := strconv.Atoi(nums[1])
				if err != nil{
					panic("Введено не чило...")
				}
				fmt.Println(num1 - num2)
			}
	} else if strings.Contains(string_enter, "*"){
		nums := strings.Split(string_enter, "*")
		if value1, flag := romeToArab[nums[0]]; flag {
			if value2, flag := romeToArab[nums[1]]; flag {
				fmt.Println(value1 * value2)
			} else {
				panic("Неверный формат: Должны быть римские или арабские цифры")
			}
		} else {
				num1, err := strconv.Atoi(nums[0])
				if err != nil{
					panic("Введено не число...")
				}
				num2, err := strconv.Atoi(nums[1])
				if err != nil{
					panic("Введено не чило...")
				}
				fmt.Println(num1 * num2)
			}
	} else if strings.Contains(string_enter, "/"){
		nums := strings.Split(string_enter, "/")
		if value1, flag := romeToArab[nums[0]]; flag {
			if value2, flag := romeToArab[nums[1]]; flag {
				fmt.Println(value1 / value2)
			} else {
				panic("Неверный формат: Должны быть римские или арабские цифры")
			}
		} else {
				num1, err := strconv.Atoi(nums[0])
				if err != nil{
					panic("Введено не число...")
				}
				num2, err := strconv.Atoi(nums[1])
				if err != nil{
					panic("Введено не чило...")
				}
				fmt.Println(num1 / num2)
			}
	} 

}

func Arab(res int) string {
	resRome := ""
	i := 1
	for ; res < 1; i++{
		otv := res % 10
		fmt.Println(arabToRome[20])
		switch i {
			case 1:
				resRome = arabToRome[otv] + resRome
				res = res / 10
			case 2:
				otv = otv * 10
				resRome = arabToRome[otv] + resRome
				res = res / 10
			case 3:
				otv = otv * 100
				resRome = arabToRome[otv] + resRome
				res = res / 10
			default:
				return "Hello"
		}
	}
	return resRome
}