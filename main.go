package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/niiket/Sample/numerus"
)

type Data struct {
	A        string
	B        string
	Operator string
}

func NewData(data ...string) *Data {
	return &Data{
		A:        data[0],
		B:        data[2],
		Operator: data[1],
	}
}

func (d *Data) GetCheck() (string, error) {
	var check int
	var res string
	_, err := strconv.Atoi(d.A)
	if err != nil {
		check++
	}

	_, err = strconv.Atoi(d.B)
	if err != nil {
		check++
	}

	if check == 0 {
		res = "Arabic"
	} else if check == 1 {
		return res, errors.New("Ошибка: Используются одновременно разные системы счисления.")
	} else {
		res = "Roman"
	}

	if d.Operator != "+" && d.Operator != "-" && d.Operator != "*" && d.Operator != "/" {
		return res, errors.New(fmt.Sprintf("Ошибка: Оператор %s не удовдетворяет математической операции.", d.Operator))
	}

	return res, nil
}

func (d *Data) GetСalc() (string, error) {
	var a, b, res int
	var result string

	numeric, err := d.GetCheck()
	if err != nil {
		return "", err
	}

	if numeric == "Arabic" {
		a, err = strconv.Atoi(d.A)
		if err != nil {
			return "", err
		}

		b, err = strconv.Atoi(d.B)
		if err != nil {
			return "", err
		}

	} else {
		a, err = parsePomanToArabicNumeric(d.A)
		if err != nil {
			return "", err
		}
		b, err = parsePomanToArabicNumeric(d.B)

	}

	switch d.Operator {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	}

	if numeric == "Roman" {
		if res <= 0 {
			return "", errors.New("Ошибка: Результ операции с Римскими цифрами может быть только положительным.")
		}

		result = numerus.Numeral(res).String()
	} else {
		result = strconv.Itoa(res)
	}

	return result, nil
}

func parsePomanToArabicNumeric(a string) (int, error) {
	var val int
	switch a {
	case "I":
		val = 1
	case "II":
		val = 2
	case "III":
		val = 3
	case "IV":
		val = 4
	case "V":
		val = 5
	case "VI":
		val = 6
	case "VII":
		val = 7
	case "VIII":
		val = 8
	case "IX":
		val = 9
	case "X":
		val = 10
	default:
		return 0, errors.New("Ошибка: Калькулятор не работает с данным диапозоном чисел !")
	}
	return val, nil
}

func main() {

	retriev := bufio.NewScanner(os.Stdin)
	retriev.Scan()
	temp := retriev.Text()
	temp = strings.TrimSpace(temp)
	example := strings.Split(temp, " ")

	if len(example) > 3 {
		fmt.Println("Ошибка: Формат математической операции не удовлетворяет заданию — два операнда и один оператор.")
		os.Exit(0)
	}

	data := NewData(example...)

	result, err := data.GetСalc()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(result)
}
