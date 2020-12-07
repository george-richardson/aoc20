package day7_handy_haversacks

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) scanFor(token Token) (literal string, err error) {
	tok, lit := p.scan()
	if tok != token {
		return "", fmt.Errorf("unexpected token %v found looking for %v at %v", tok, token, lit)
	}
	return lit, nil
}

func (p *Parser) scanForWS() (err error) {
	_, err = p.scanFor(WS)
	return
}

func (p *Parser) scanForSequence(tokens *[]Token) (literal string, err error) {
	tokensDeref := *tokens
	builder := strings.Builder{}
	for i := 0; i < len(tokensDeref); i++ {
		lit, err := p.scanFor(tokensDeref[i])
		if err != nil {
			p.unscan()
			return "", err
		}
		builder.WriteString(lit)
	}
	return builder.String(), nil
}

func (p *Parser) ParseLine() (*bagRule, error) {
	name, err := p.scanForSequence(&[]Token{WORD, WS, WORD})
	if err != nil {
		return nil, err
	}

	rule := bagRule{
		name: name,
		contains: []struct {
			name  string
			count int
		}{},
	}

	_, err = p.scanForSequence(&[]Token{WS, BAG, WS})
	if err != nil {
		return nil, err
	}

	if _, err = p.scanForSequence(&[]Token{CONTAIN, WS}); err != nil {
		return nil, err
	}
	_, err = p.scanForSequence(&[]Token{NO, WS, OTHER, WS, BAG, PERIOD})
	if err == nil {
		return &rule, nil
	}
OUT:
	for true {
		tok, lit := p.scan()
		switch tok {
		case COMMA:
			if err := p.scanForWS(); err != nil {
				return nil, err
			}
			tok = NUMBER
			lit, err = p.scanFor(NUMBER)
			if err != nil {
				return nil, err
			}
			fallthrough
		case NUMBER:
			value, err := strconv.Atoi(lit)
			if err != nil {
				return nil, err
			}

			if err := p.scanForWS(); err != nil {
				return nil, err
			}

			containedName, err := p.scanForSequence(&[]Token{WORD, WS, WORD})

			_, err = p.scanForSequence(&[]Token{WS, BAG})
			if err != nil {
				return nil, err
			}
			rule.contains = append(rule.contains, struct {
				name  string
				count int
			}{name: containedName, count: value})
		case PERIOD:
			break OUT
		default:
			return nil, fmt.Errorf("unexpected token %v at %v", tok, lit)
		}
	}
	return &rule, nil
}

type bagRule struct {
	name     string
	contains []struct {
		name  string
		count int
	}
}
