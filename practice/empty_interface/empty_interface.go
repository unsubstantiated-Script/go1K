package empty_interface

import "fmt"

func EmptyInterface() {
	//var any interface{}
	//
	////Can store any type in an empty interface...
	//
	//any = []int{1, 2, 3}
	//any = map[int]bool{1: true, 2: false}
	//any = "kak"
	//any = 3
	//// You can't directly use the dynamic value in operations
	////any = any * 3
	////Must extract via type assertion
	//any = any.(int) * 2
	//fmt.Println(any)

	nums := []int{1, 2, 3}

	var any interface{}
	//You can assign a slice to an empty interface value.
	any = nums

	//Need to extract the int slice value to get this out.
	_ = len(any.([]int))

	var many []interface{}
	//Can't directly assign to an empty interface slice to another empty interface.
	//many = nums
	//var words []string = nums

	//A loop will do it tho.
	for _, n := range nums {
		many = append(many, n)
	}

	//Many is a slice of interface{} values, so it can't store anything. Only an interface can story any type of value.
	fmt.Println(many)
}
