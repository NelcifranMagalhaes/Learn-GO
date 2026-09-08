package main

import "fmt"

func mathCalc(n1, n2 int) (soma int, subt int) {
	soma = n1 + n2
	subt = n1 - n2
	return
}

func main() {
	soma, subt := mathCalc(30, 10)
	fmt.Println(soma, subt)
}
