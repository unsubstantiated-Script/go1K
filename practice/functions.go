package practice

import (
	"fmt"
	"strconv"
)

func FunctionStuff() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("main -> local = %d\n", local)

	local = incr(local)
	show(local)

	local = incrBy(local, 5)
	show(local)

	_, err := incrByStr(local, "TWO")
	if err != nil {
		fmt.Printf("err -> %s\n", err)
	}

	local, err = incrByStr(local, "2")
	if err != nil {
		fmt.Printf("err -> %s\n", err)
	}
	show(local)

	//This is only returning a copy of the value not actually changing the value of "local"
	show(incrBy(local, 2))
	show(local)

	////This is successfully sanitized and will return zero because the string won't convert
	//local = sanitize(incrByStr(local, "NOPE"))
	//show(local)

	//This is successfully sanitized and will return zero because the string won't convert
	local = sanitize(incrByStr(local, "2"))
	show(local)

	local = limit(incrBy(local, 5), 500)
	show(local)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, factor int) int {
	return n * factor
}

func incrByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)
	return n, err
}

// Deals with aligning multiple arguments and returns.
func sanitize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

// Passing a  return like this instantiates its value (0) and type.
func limit(n, lim int) (m int) {
	m = n
	if m >= lim {
		m = lim
	}
	//This will go ahead and return our stuffs.
	return
}
