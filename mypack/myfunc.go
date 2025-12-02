// пакет предоставляет функию для умножения двух чисел
package mypack

import (
	"golang.org/x/exp/constraints"
)

// функция Add принимает два целых числа,
// находит их сумму и возвращает результат
//
// подробности можно найти [тут]
//
// [тут]: https://www.mathsisfun.com/numbers/addition.html

type Number interface{
	constraints.Integer | constraints.Float
}
// func Add(a,b int) int{
// 	return a + b
// } 
func Add[T Number] (a, b T) T{
	return a + b
}

