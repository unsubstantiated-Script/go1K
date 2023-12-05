package practice

type computer struct {
	brand string
}

type computer2 struct {
	brand *string
}

func MakeDaPointer() {

	//var (
	//	counter int
	//	V       int
	//	P       *int
	//)
	//
	//counter = 100
	//P = &counter
	//
	//V = *P

	//if P == nil {
	//	fmt.Printf("P is nil and its address is %p\n", P)
	//}
	//
	//if P == &counter {
	//	fmt.Printf("P is nil and its address is %p == %p\n", P, &counter)
	//}

	//fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	//fmt.Printf("P      : %-13p addr: %-13p  *P: %-13d\n", P, &P, *P)
	//fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)
	//
	//fmt.Println("\n....... change counter")
	//
	//counter = 10
	//fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	//counter = passVal(counter)
	//fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	//
	//fmt.Println("\n....... change counter")
	//
	//passPtrVal(&counter)
	//fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)

	//var null *computer
	//
	//if null == nil {
	//	fmt.Println("null is nil")
	//}
	//
	//apple := &computer{brand: "apple"}
	//
	//newApple := apple
	//
	//if apple == newApple {
	//	fmt.Printf("these equal each other. apple: %p and newApple: %p\n", apple, newApple)
	//}
	//
	//sony := &computer{brand: "sony"}
	//
	//if sony != apple {
	//	fmt.Printf("these don't equal each other. sony: %p and apple: %p\n", sony, apple)
	//}
	//
	//// put apple's value into a new ordinary variable
	//appleVal := *apple
	//
	//// print apple pointer variable's address, and the address it points to
	//// and, print the new variable's addresses as well
	//fmt.Printf("apple                     : %p %p\n", &apple, apple)
	//fmt.Printf("appleVal                  : %p\n", &appleVal)
	//
	//if *apple == appleVal {
	//	fmt.Println("Deez are equal!")
	//
	//	fmt.Printf("*apple: %+v   appleVal: %+v\n", *apple, appleVal)
	//}
	//
	//change(sony, "hp")
	//
	//fmt.Printf("sony is now %s\n", sony.brand)
	//
	//fmt.Printf("sony : %s\n", returnVal(apple))
	//
	//fmt.Printf("sony : %s\n", newComputer("dell"))
	//fmt.Printf("sony : %s\n", newComputer("lenovo"))
	//fmt.Printf("sony : %s\n", newComputer("asus"))

	//a, b := 3.14, 6.28
	//swap(&a, &b)
	//fmt.Printf("a : %g — b : %g\n", a, b)
	//
	//pa, pb := &a, &b
	//swapAdd(pa, pb)
	//fmt.Printf("pa: %p — pb: %p\n", pa, pb)
	//fmt.Printf("pa: %g — pb: %g\n", *pa, *pb)

	//c := &computer2{}
	//change(c, "apple")
	//fmt.Printf("brand: %s\n", *c.brand)
}

//func change(c *computer2, brand string) {
////	c.brand = &brand
////}

//
//func swap(a, b *float64) {
//	*a, *b = *b, *a
//}
//
//func swapAdd(a, b *float64) (*float64, *float64) {
//	return b, a
//}

//func change(c *computer, brand string) {
//	c.brand = brand
//}
//
//func returnVal(c *computer) computer {
//	return *c
//}
//
//func newComputer(brand string) *computer {
//	return &computer{brand: brand}
//}

//func passVal(n int) int {
//	n = 50
//	fmt.Printf("n     : %-13d addr: %-13p\n", n, &n)
//
//	return n
//}
//
//func passPtrVal(pn *int) {
//	fmt.Printf("pn      : %-13p addr: %-13p  *pn: %-13d\n", pn, &pn, *pn)
//	*pn++
//	fmt.Printf("pn      : %-13p addr: %-13p  *pn: %-13d\n", pn, &pn, *pn)
//}
