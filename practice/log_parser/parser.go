package log_parser

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result //total visits per domain
	domains []string          // unique domain names
	total   int               // total visits to all domains
	lines   int               // number of parsed lines (for the error messages)

}

func makeParser() parser {
	//This is a constructor pattern
	return parser{
		sum: make(map[string]result),
	}
}

// Declaring values in the return statement allows you to avoid having to declare them at each return within the function.
func parse(p parser, line string) (parsed result, err error) {

	fields := strings.Fields(line)

	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	parsed.domain = fields[0]

	//Sum the total visits per domain
	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d), fields[1], lines")
		return
	}

	return
}

func update(p parser, parsed result) parser {
	domain, visits := parsed.domain, parsed.visits

	// Collect the unique domains
	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	//Keep track of total and per domain visits
	p.total += visits

	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}

	return p
}
