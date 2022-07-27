/*
 * @author: Emove
 * @date: Do not edit
 */
package code

import "fmt"

func Test() {
	grade := "B"
	marks := 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70, 60, 50:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀！")
	case grade == "B":
		fmt.Println("良好")
	default:
		fmt.Println("差")
	}
}
