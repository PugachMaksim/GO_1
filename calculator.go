package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strconv"
)

var romeToArab = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabToRome = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	20:  "XX",
	30:  "XXX",
	40:  "XL",
	50:  "L",
	60:  "LX",
	70:  "LXX",
	80:  "LXXX",
	90:  "XC",
	100: "C",
}

func main() {
	var string_enter string
	fmt.Fscan(os.Stdin, &string_enter)
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,/%,-%,+", r)
	}
	nummms := strings.FieldsFunc(string_enter, splitFunc)
	if len(nummms) > 2 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if len(nummms) < 2 {
		panic("строка не является математической операцией.")
	} else {
		num1, err := strconv.Atoi(nummms[0])
		if err != nil {
			if value1, flag := romeToArab[nummms[0]]; flag {
				if value2, flag := romeToArab[nummms[1]]; flag {
					if strings.Contains(string_enter, "+") {
						res := value1 + value2
						fmt.Println(Arab(res))
					} else if strings.Contains(string_enter, "-") {
						if value1 < value2 {
							panic("в римской системе нет отрицательных чисел")
						}
						res := (value1 - value2)
						if res == 0 {
							panic("Ноль, ошибка")
						}
						fmt.Println(Arab(res))
					} else if strings.Contains(string_enter, "*") {
						res := (value1 * value2)
						fmt.Println(Arab(res))
					} else if strings.Contains(string_enter, "/") {
						res := (value1 / value2)
						if value1 % value2 == 0 {
							fmt.Println(Arab(res))
						} else {
							panic("Не целое число!!!")
						}
					}
				} else {
					value2, err := strconv.Atoi(nummms[1])
					if err == nil {
						panic("используются одновременно разные системы счисления.")
					}
					_ = value2
					panic("До десяти!! 1")
				}
			} else {
				panic("До десяти!! 2")
			}

		} else {
			num2, err := strconv.Atoi(nummms[1])
			if err != nil {
				if num_2, ok := romeToArab[nummms[1]]; ok {
					_ = num_2
					panic("Разные системы счисления")
				}
				panic("Введено не целое число 2")
			} else {
				if num1 != 0 && num2 != 0 {
					if strings.Contains(string_enter, "+") {
						fmt.Println(num1 + num2)
					} else if strings.Contains(string_enter, "-") {
						fmt.Println(num1 - num2)
					} else if strings.Contains(string_enter, "*") {
						fmt.Println(num1 * num2)
					} else if strings.Contains(string_enter, "/") {
						fmt.Println(num1 / num2)
					}
				} else {
					panic("Не должно быть нуля, почему то")
				}
			}
		}
	}
}

func Arab(res int) string {
	/*
		конвертация Араб -> Рим
	*/
	resRome := ""
	count := 1
	for res > 0 {
		otv := res % 10
		switch count {
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
		count++
	}
	return resRome
}
