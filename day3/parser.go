package day3

import (
	"fmt"
	"strconv"
)

type Parser struct {
	currentExpression                     string
	allowedExpressions                    map[string]bool
	parsingMul, parsingLeft, parsingRight bool
	left, right                           int
	parsedSingle                          bool
	results                               []Result
	mulMode                               bool
	parsingMulMode                        bool
}

type Result struct {
	val     int
	enabled bool
}

func NewParser() *Parser {
	parser := Parser{}
	parser.currentExpression = ""
	parser.allowedExpressions = make(map[string]bool)
	parser.allowedExpressions["mul"] = true
	parser.mulMode = true
	return &parser
}

func (p *Parser) Parse(expression string) {
	for _, c := range expression {
		p.AddCharacter(c)
	}
}

func (p *Parser) AddCharacter(char rune) {
	validChar := p.EvalCharacter(char)

	if !validChar {
		p.currentExpression = ""
		p.parsingMul = false
		p.parsingLeft = false
		p.parsingRight = false
		p.left = 0
		p.right = 0
		return
	}

	p.currentExpression = p.currentExpression + string(char)

	if p.parsingMul {
		p.parsingMul = false
		p.parsingLeft = true
		p.currentExpression = ""
	}

	if p.parsingLeft && p.left != 0 {
		p.parsingLeft = false
		p.parsingRight = true
		p.currentExpression = ""
	}

	if p.currentExpression == "mul" {
		p.parsingMul = true
		p.currentExpression = ""
	}

	if p.parsedSingle {
		p.parsingRight = false
		p.parsingLeft = false
		p.parsingMul = false
		p.parsedSingle = false
		p.results = append(p.results, Result{val: p.left * p.right, enabled: p.mulMode})
		p.left = 0
		p.right = 0
		p.currentExpression = ""
	}
}

func (p *Parser) EvalCharacter(c rune) bool {
	p.EvalMulMode(c)

	if p.parsingMulMode {
		return true
	}

	if p.parsingMul {
		if c == '(' {
			return true
		}
	}

	if p.parsingLeft {
		if p.currentExpression == "" && c == ',' {
			return false
		}
		if c == ',' {
			num, err := strconv.Atoi(p.currentExpression)
			if err != nil {
				panic(fmt.Sprintf("Cannot convert num: %s", p.currentExpression))
			}

			p.left = num
			return true
		}
		return c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' || c == '0'
	}

	if p.parsingRight {
		if p.currentExpression == "" && c == ')' {
			return false
		}
		if c == ')' {
			num, err := strconv.Atoi(p.currentExpression)
			if err != nil {
				panic(fmt.Sprintf("Cannot convert num: %s", p.currentExpression))
			}

			p.right = num
			p.parsedSingle = true
			return true
		}
		return c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' || c == '0'
	}

	return p.EvalAny(c)
}

func (p *Parser) EvalMulMode(c rune) {
	p.parsingMulMode = false

	if c == 'd' {
		p.parsingMulMode = true
		return
	}

	if p.currentExpression == "d" && c == 'o' {
		p.parsingMulMode = true
		return
	}

	if p.currentExpression == "do" && c == '(' {
		p.parsingMulMode = true
		return
	}

	if p.currentExpression == "do(" && c == ')' {
		p.mulMode = true
		p.currentExpression = ""
		return
	}

	if p.currentExpression == "do" && c == 'n' {
		p.parsingMulMode = true
		return
	}

	if p.currentExpression == "don" && c == '\'' {
		p.parsingMulMode = true
		return
	}

	if p.currentExpression == "don'" && c == 't' {
		p.parsingMulMode = true
		return
	}

	if p.currentExpression == "don't" && c == '(' {
		p.parsingMulMode = true
		return
	}

	if p.currentExpression == "don't(" && c == ')' {
		p.mulMode = false
		p.currentExpression = ""
		return
	}
}

func (p *Parser) EvalAny(c rune) bool {
	if c == 'm' {
		return true
	}

	if p.currentExpression == "m" && c == 'u' {
		return true
	}

	if p.currentExpression == "mu" && c == 'l' {
		return true
	}

	return false
}
