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
	fmt.Println("Введите выражение (например: \"100\" + \"500\"):")

	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	operator, err := splitOperator(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	str1, str2, err := parseExpression(expression, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if len(str1) > 10 || len(str2) > 10 {
		fmt.Println("Ошибка: Строка должна быть длиной не более 10 символоы")
		return
	}

	result, err := calculate(operator, str1, str2)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	resultOutput := formatResult(result)
	fmt.Printf("\"%s\"\n", resultOutput)
}

func splitOperator(expression string) (string, error) {
	for _, value := range []string{" - ", " + ", " * ", " / "} {
		if strings.Contains(expression, value) {
			return value, nil
		}
	}
	return "", fmt.Errorf("Некорректный ввод. Введите оператор +, -, /, *")
}

func parseExpression(expression, operator string) (string, string, error) {
	parts := strings.Split(expression, operator)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Некорректный ввод. Введите выражение в формате \"строка\" оператор \"строка\"")
	}
	if !isQuotes(parts[0]) {
		return "", "", fmt.Errorf("Первым аргументом выражения, подаваемым на вход, должна быть строка!")
	}
	if operator == " + " || operator == " - " {
		if !isQuotes(parts[1]) {
			return "", "", fmt.Errorf("Обе строки должны быть возведены в ковычки!")
		}
	}
	if operator == " / " || operator == " * " {
		if isQuotes(parts[1]) {
			return "", "", fmt.Errorf("Второй аргумент должен быть числом")
		}
	}
	str1 := strings.TrimSpace(strings.Replace(parts[0], "\"", "", -1))
	str2 := strings.TrimSpace(strings.Replace(parts[1], "\"", "", -1))
	fmt.Println(str1)

	return str1, str2, nil
}

func calculate(operator, str1, str2 string) (string, error) {
	switch operator {
	case " + ":
		return str1 + str2, nil
	case " - ":
		return strings.Replace(str1, str2, "", -1), nil
	case " * ":
		multiplier, err := strconv.Atoi(str2)
		if err != nil {
			fmt.Println("Ошибка: некорректное число для умножения")
		}
		if multiplier < 1 || multiplier > 10 {
			return "", fmt.Errorf("Число должно быть в диапазоне от 1 до 10.")
		}
		return strings.Repeat(str1, multiplier), nil
	case " / ":
		divider, err := strconv.Atoi(str2)
		if err != nil {
			return "", fmt.Errorf("Некорректное число для деления")
		}
		if isQuotes(str2) {
			return "", fmt.Errorf("Делить можно только на число")
		}
		if divider == 0 {
			return "", fmt.Errorf("Делить на ноль нельзя")
		}
		if divider < 1 || divider > 10 {
			return "", fmt.Errorf("Число должно быть в диапазоне от 1 до 10.")
		}
		if divider > len(str1) {
			return " ", nil
		}
		return str1[:len(str1)/divider], nil
	default:
		return "", fmt.Errorf("Некорректный ввод. Введите оператор +, -, /, *")
	}
}

func isQuotes(str string) bool {
	return len(str) > 2 && str[0] == '"' && str[len(str)-1] == '"'
}

func formatResult(result string) string {
	if len(result) > 40 {
		return result[:40] + "..."
	}
	return result
}
