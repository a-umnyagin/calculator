package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите математическую операцию c двумя операндами и одним оператором (+, -, /, *).")
	var math [4]string
	var a, b, c, aR, bR, cR int
	Roman := [101]string{"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV",
		"XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII",
		"XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI",
		"XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI",
		"LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII",
		"LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV",
		"XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
	}

	fmt.Scanln(&math[0], &math[1], &math[2], &math[3]) //Ввод с консоли

	if math[0] == "" || math[1] == "" || math[2] == "" { //Проверка на наличие операции
		fmt.Println("Строка не является математической операцией")
		return
	}

	if math[3] != "" { //Проверка на формат операции
		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	for i := 1; i <= 10; i++ { //Конвертация римских цифр в арабские
		if math[0] == Roman[i] {
			aR = i
		}
		if math[2] == Roman[i] {
			bR = i
		}
	}
	if aR != 0 && bR != 0 { //Проверка римских цифр на соответствие условиям задания
		if math[1] == "+" { //Определение оператора и проверка корректности
			cR = aR + bR
		} else if math[1] == "-" {
			cR = aR - bR
			if cR == 0 {
				fmt.Println("Недопустимая операция, в римской системе счисления нет нуля")
				return
			}
			if cR < 0 {
				fmt.Println("Недопустимая операция, в римской системе счисления нет отрицательных чисел")
				return
			}
		} else if math[1] == "*" {
			cR = aR * bR
		} else if math[1] == "/" {
			cR = aR / bR
		} else {
			fmt.Println("Недопустимое значение оператора")
			return
		}
		fmt.Println(Roman[cR]) //Ответ
		return
	}

	a, err := strconv.Atoi(math[0]) //Проверка корректности переменной
	if (err != nil || a < 1 || a > 10) && aR == 0 {
		fmt.Println("Недопустимое значение переменной")
		return
	}

	b, erro := strconv.Atoi(math[2]) //Проверка корректности переменной
	if (erro != nil || b < 1 || b > 10) && bR == 0 {
		fmt.Println("Недопустимое значение переменной")
		return
	}

	if (aR == 0 && bR != 0) || (aR != 0 && bR == 0) { //Проверка на одинаковую систему счисления
		fmt.Println("Недопустимо использовать разные системы счисления")
		return
	}

	if math[1] == "+" { //Определение оператора и проверка корректности
		c = a + b
	} else if math[1] == "-" {
		c = a - b
	} else if math[1] == "*" {
		c = a * b
	} else if math[1] == "/" {
		c = a / b
	} else {
		fmt.Println("Недопустимое значение оператора")
		return
	}

	fmt.Print(c) //Ответ
}
