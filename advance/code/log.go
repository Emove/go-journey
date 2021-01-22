/**
 * @author Emove
 * @date 2021/1/21
 */
package code

import (
	"github.com/sirupsen/logrus"
)

type Person struct {
	name   string
	age    uint32
	gender uint32
}

func LogInfo() {
	person := Person{
		name:   "Emove",
		age:    20,
		gender: 1,
	}
	logrus.Infof("person info: %v", person)
}
