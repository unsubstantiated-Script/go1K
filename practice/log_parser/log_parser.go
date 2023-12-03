package log_parser

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func LogParser() {

	p := makeParser()

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		parsed := parse(p, in.Text())
		//Not going to get the info back from the function unless this override is in place.
		update(p, parsed)
	}

	summarize(p)
	dumpErrs([]error{in.Err(), err(p)})
}

// No Pointer needed here because this function doesn't change anything, just prints it.
func summarize(p *parser) {

	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}

	fmt.Println()
	fmt.Printf("%-30s %10d\n", "TOTAL", p.total)
}

// dumpErrs simplifies handling multiple errors
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}
