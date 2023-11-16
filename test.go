// Короттаев Николай
//
// тестовое задание - test.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	difSys := errors.New("Разные системы исчеслений")
	toMuch := errors.New("Ввод не удовлетваряет заданию - два операнда и один оператор(+, -, *, /)") //Создаем ошибки для завершения программы и вывода в консоль
	negRoman := errors.New("В римсской системе нет отрицательных чисел")
	notMath := errors.New("Не является математической операцией")
	numLim := errors.New("Входные данные должныбыть римскими или арабскими цифрами от 1 до 10")

	scanner := bufio.NewScanner(os.Stdin) //Сканируем стору с терминала
	scanner.Scan()
	line := scanner.Text()

	var midle int = 0 //Переменна я для обозначения операттора сттроки

	for i, c := range line {
		if c == '+' || c == '-' || c == '*' || c == '/' { // Вычисляем оператор и проверяем, что он один
			if midle >= 1 {
				fmt.Print(toMuch)
				return
			}
			midle = i
		}
	}

	first := strings.TrimSpace(line[:midle])
	second := strings.TrimSpace(line[(midle + 1):]) // делим строку не левую и правую части от оператора и убираем все пробелы

	if len(second) == 0 && len(first) == 0 { // проверяем что ни одна из подстрок не пустые
		fmt.Print(notMath)
		return
	}

	if _, err := strconv.Atoi(first); err == nil {
		if _, err := strconv.Atoi(second); err != nil { // проверяем чтто обе подстрооки одной системы исчеслений
			fmt.Print(difSys)
			return
		}
	}
	if _, err := strconv.Atoi(first); err != nil {
		if _, err := strconv.Atoi(second); err == nil { // проверяем чтто обе подстрооки одной системы исчеслений
			fmt.Print(difSys)
			return
		}
	}
	if _, err := strconv.Atoi(first); err != nil {
		if _, err := strconv.Atoi(second); err != nil { //Проверяем что обе подстроки римские цифры
			rfirst := romanToInt(first)
			rsecond := romanToInt(second)
			if rsecond > 11 || rsecond < 0 && rfirst < 0 || rfirst > 11 {
				fmt.Print(numLim)

			} else {
				result := calculator(rfirst, rsecond, midle, line) // С помощью функции переводим из римской системы в целочисленные и спомощью друугой функции выполняем математическую операцию
				if result > 0 {
					fmt.Println(intToRoman(result)) // Если число положительное выодим результтат если нет, ошибку
				} else if result < 0 {
					fmt.Print(negRoman)
					return
				} else {
					fmt.Print(numLim)
				}

			}
		}
	}
	if first, err := strconv.Atoi(first); err == nil {
		if second, err := strconv.Atoi(second); err == nil { // проверяем что обе подстрооки арабские числа переводим их в инт и при помощи функции вывоодим результат
			if first < 0 || first > 11 && second < 0 || second > 11 {
				fmt.Print(numLim)
			} else {
				fmt.Println(calculator(first, second, midle, line))
			}
		}
	}
}
func calculator(first, second, midle int, line string) int { // Задаем функцию высчитывающию результат
	switch operator := line[midle]; operator {
	case '-':
		return (first - second)
	case '+':
		return (first + second)
	case '*':
		return (first * second)
	case '/':
		return (first / second)
	}
	return 0
}

func romanToInt(s string) int { // Задаем функцию для перевода римских цифр в арабские
	rMap := map[string]int{"I": 1, "V": 5, "X": 10}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func intToRoman(num int) string { // задаем функцию для перевод результтата в римские цифры
	nums := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string

	for i := 0; i < len(nums); i++ {
		result += strings.Repeat(romans[i], num/nums[i])
		num %= nums[i]
	}

	return result
}
