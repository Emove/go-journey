package code

import (
	"fmt"
)

func Contains() {
	fileName := "map_lab.yaml"
	fmt.Println(fileName[0:5])
	fmt.Println(fileName[len(fileName)-5:])
}

func JointSql() {
	trackId := 1
	xs := []int{1, 2, 3}
	y := 2
	condition := fmt.Sprintf("track_id = %d", trackId)
	if len(xs) > 0 {
		condition = condition + " and now_x in ("
		for i := range xs {
			condition = condition + fmt.Sprintf("%d, ", xs[i])
		}
		condition = condition[:len(condition)-2] + ") "
	}
	condition = condition + fmt.Sprintf(" and now_y = %d", y)
	fmt.Println(condition)
}
