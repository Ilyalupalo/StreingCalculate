package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: \"Hello\" * \"5\"):")

	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	str1, operator, str2, err := parseExpression(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if len(str1) > 10 || len(str2) > 10 {
		fmt.Println("Ошибка: Строка должна быть длиной не более 10 символоы")
		return
	}

	result, err := calculate(str1, operator, str2)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	resultOutput := formatResult(result)
	fmt.Printf("\"%s\"\n", resultOutput)
}

func parseExpression(expression string) (string, string, string, error) {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("Некорректный ввод. Введите выражение в формате \"строка\" оператор \"строка\"")
	}
	if !isQuotes(parts[0]) {
		return "", "", "", fmt.Errorf("Первым аргументом выражения, подаваемым на вход, должна быть строка!")
	}
	str1 := strings.Replace(parts[0], "\"", "", -1)
	operator := parts[1]
	str2 := strings.Replace(parts[2], "\"", "", -1)

	return str1, operator, str2, nil
}

func calculate(str1, operator, str2 string) (string, error) {
	switch operator {
	case "+":
		if !isQuotes(str2) {
			return "", fmt.Errorf("Обе строки должны быть возведены в ковычки!")
		}
		return str1 + str2, nil
	case "-":
		if !isQuotes(str2) {
			return "", fmt.Errorf("Обе строки должны быть возведены в ковычки!")
		}
		return strings.Replace(str1, str2, "", -1), nil
	case "*":
		multiplier, err := strconv.Atoi(str2)
		if err != nil {
			fmt.Println("Ошибка: некорректное число для умножения")
		}
		if multiplier < 1 || multiplier > 10 {
			return "", fmt.Errorf("Число должно быть в диапазоне от 1 до 10.")
		}
		return strings.Repeat(str1, multiplier), nil
	case "/":
		divider, err := strconv.Atoi(str2)
		if err != nil {
			return "", fmt.Errorf("Некорректное число для деления")
		}
		if divider < 1 || divider > 10 {
			return "", fmt.Errorf("Число должно быть в диапазоне от 1 до 10.")
		}
		if divider == 0 {
			return "", fmt.Errorf("Делить на ноль нельзя")
		}
		if divider > len(str1) {
			return " ", nil
		}
		return str1[:len(str1)/divider], nil
	default:
		return "", fmt.Errorf("Некорректный ввод. Введите оператор +, -, /, *")
	}
}

func isQuotes(str1 string) bool {
	return len(str1) > 2 && str1[0] == '"' && str1[len(str1)-1] == '"'
}

func formatResult(result string) string {
	if len(result) > 40 {
		return result[:40] + "..."
	}
	return result
}
