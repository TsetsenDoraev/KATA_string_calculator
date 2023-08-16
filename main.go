package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	// Приветствие
	Greeting := "\n\nДобро пожаловать в строчный калькулятор"
	fmt.Println(Greeting)
	Delimiter := strings.Repeat("—", utf8.RuneCountInString(Greeting)-2)
	fmt.Println(Delimiter)

	// Запрос выражения у пользователя
	fmt.Println("Введите выражение: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\r\n")
	//cleanInput := strings.Fields(input)
	//FinalInput := strings.Join(cleanInput, "")

	// Определение оператора в выражении для Split и Switch
	var operator string
	operators := [...]string{" + ", " - ", " * ", " / "}
	for _, op := range operators {
		if strings.Contains(input, op) {
			operator = op
			break
		} else if op == operators[len(operators)-1] {
			fmt.Println("Оператор не найден, строка не является математической операцией")
			return
		}
	}

	operands := strings.Split(input, operator)
	if len(operands) > 2 {
		fmt.Println("Неверный формат выражения, не более двух операндов в выражении")
		return
	}
	if len(operands) < 2 {
		fmt.Println("Cтрока не является математической операцией")
		return
	}

	var typeOperand_1, typeOperand_2 string

	if operands[0][0] == '"' && operands[0][len(operands[0])-1] == '"' {
		if len(operands[0]) < 11 {
			typeOperand_2 = "str"
		} else {
			fmt.Println("Длинна строки не более 10 символов, проблема в первом операнде")
			return
		}
		//fmt.Printf("Тип первой переменной: %T\n", operands[0])
	} else {
		fmt.Println("Первый операнд может быть только строкой, используйте кавычки")
	}

	if operands[1][0] == '"' && operands[1][len(operands[1])-1] == '"' {
		operands[1] = operands[1][1 : len(operands[1])-1]
		if len(operands[1]) < 11 {
			typeOperand_2 = "str"
		} else {
			fmt.Println("Длинна строки не более 10 символов, проблема в втором операнде")
			return
		}
		//fmt.Printf("Тип второй переменной: %T\n", operands[1])

	} else if matched, _ := regexp.MatchString("\\b([1-9]|10)\\b", operands[1]); matched {
		typeOperand_2 = "int"

	} else {
		fmt.Println("Второй операнд должен быть числом от 1 до 10 или строкой")
	}

	//fmt.Println(operands[0], operator, operands[1]) // Удалить позже

	if typeOperand_1 == "str" && typeOperand_2 == "str" {
		var result string
		switch operator {
		case " + ":
			result = StrStrPlus(operands[0], operands[1])
			fmt.Println("\"" + result + "\"")
		case " - ":
			if !strings.Contains(operands[0], operands[1]) {
				result = operands[0]
				fmt.Println("\"" + result + "\"")
			} else {
				result = StrStrMinus(operands[0], operands[1])
				fmt.Println("\"" + result + "\"")
			}
		default:
			fmt.Println("Две строки можно только сложить или вычесть")
		}
	} else if typeOperand_1 == "str" && typeOperand_2 == "int" {
		operandsInt, err := strconv.Atoi(operands[1])
		if err != nil {
		}
		var result string
		switch operator {
		case " * ":
			result = StrMultiply(operands[0], operandsInt)
			if len(result) > 40 {
				result = result[:40]
				fmt.Println("\"" + result + "..." + "\"")
			} else {
				fmt.Println("\"" + result + "\"")
			}
		case " / ":
			result = strIntDivide(operands[0], operandsInt)
			if len(result) > 40 {
				result = result[:40]
				fmt.Println("\"" + result + "..." + "\"")
			} else {
				fmt.Println("\"" + result + "\"")
			}
		default:
			fmt.Println("Строку и число можно только умножить или разделить")
		}

	}

}

func StrStrPlus(num1, num2 string) string {
	result := num1 + num2
	return result
}

func StrStrMinus(num1, num2 string) string {
	if !strings.Contains(num1, num2) {
		return num1
	}

	result := strings.Replace(num1, num2, "", -1)
	return result
}

func StrMultiply(num1 string, num2 int) string {
	result := strings.Repeat(num1, num2)
	return result
}

func strIntDivide(num1 string, num2 int) string {
	length := len(num1)
	index := length / num2
	result := num1[:index]
	return result
}
