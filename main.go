package main

import "fmt"

const maxDiscountSum = 2000
const maxPercentOff = 30

func main() {
	var price, finalDiscount, percentOff float64

	fmt.Println("Введите цену товара:")
	fmt.Scan(&price)

	for {
		fmt.Println("Введите скидку:")
		fmt.Scan(&percentOff)

		if percentOff > maxPercentOff {
			fmt.Println("Скидка не может быть больше 30%")
			continue
		}
		finalDiscount = price / 100 * percentOff
		if finalDiscount > maxDiscountSum {
			fmt.Println("Скидка не может быть больше 2000 рублей")
			continue
		}
		break
	}

	fmt.Println("Скидка составила:", finalDiscount)
}
