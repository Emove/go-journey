/**
 * @author Emove
 * @date 2021/2/4
 */
package time

import (
	"fmt"
	"time"
)

func GetYesterdayTime() {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05.000"))
	time := time.Now().Add(-24 * time.Hour)
	fmt.Println(time.Format("2006-01-02 03:04:05.000"))
}
