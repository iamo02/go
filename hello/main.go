package main

import "hello/customer"

func main() {
	//x := 12
	//a, s := 1, 2
	//fmt.Printf("%d%d%d", x, a, s)

	//	if x <= 2 {
	//		print(x)
	//	}

	//	var z = []int{10, 20, 30, 40, 50}
	//	z = append(z, 9)
	//	fmt.Printf("%#v", z[2:4])

	// var c map[string]string = map[string]string{}
	// c["th"] = "Thailand"
	// c["en"] = "United state"

	// sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += i
	// }
	// println(sum)

	// x := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(x); i++ {
	// 	sum += x[i]
	// }

	// println(sum)

	// for _, v := range x {
	// 	sum += v
	// }

	// println(sum)

	// var x int = 10
	// var y *int = &x
	// *y = 20
	// println(&x)
	// println(*y)
	// println(x)

	//	println(hello("O"))
	println(customer.Hello())
}

func hello(name string) string {

	return "hello " + name

}
