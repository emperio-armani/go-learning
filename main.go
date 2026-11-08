package main

import "fmt"

func main() {
	var month int
	fmt.Scan(&month)
	switch month {
	case 1:
		fmt.Println("Январь")
	case 2:
		fmt.Println("Февраль")
	case 3:
		fmt.Println("Март")
	case 4:
		fmt.Println("Апрель")
	case 5:
		fmt.Println("Май")
	case 6:
		fmt.Print("Июнь")
	case 7:
		fmt.Print("Июль")
	case 8:
		fmt.Print("Август")
	case 9:
		fmt.Print("Сентябрь")
	case 10:
		fmt.Print("Октябрь")
	case 11:
		fmt.Print("Ноябрь")
	case 12:
		fmt.Print("Декабрь")
	default:
		fmt.Print("Некорректный номер месяца")

	}
}
