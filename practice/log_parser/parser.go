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
	lerr    error
}

func makeParser() *parser {
	//This is a constructor pattern
	return &parser{
		sum: make(map[string]result),
	}
}

// Declaring values in the return statement allows you to avoid having to declare them at each return within the function.
func parse(p *parser, line string) (r result) {

	if p.lerr != nil {
		return
	}

	p.lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	var err error

	r.domain = fields[0]
	//Sum the total visits per domain
	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}
	return
}

func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	// Collect the unique domains
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	//Keep track of total and per domain visits
	p.total += r.visits

	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}

	//Need to make a clone here to manipulate the data. Data in GoLang is immutable.
	clone := p.sum[r.domain]
	_ = &clone
}

func err(p *parser) error {
	return p.lerr
}
