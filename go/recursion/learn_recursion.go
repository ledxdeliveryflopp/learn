package recursion

func FirstRecursionFactorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * FirstRecursionFactorial(n-1) // грубо говоря на месте вызова рекурсии при выходе появляется результат работы функции с вершины стека,
		// n в нашел случае хранится в этом стеке
	}
}

//func ()  {
//
//}
