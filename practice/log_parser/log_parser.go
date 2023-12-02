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
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		//Not going to get the info back from the function unless this override is in place.
		p = update(p, parsed)
	}

	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	fmt.Printf("%-30s %10d\n", "TOTAL", p.total)
	// Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
