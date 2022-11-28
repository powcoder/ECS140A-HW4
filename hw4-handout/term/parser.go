https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package term

import (
	"errors"
	"strconv"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <term>     ::= ATOM | NUM | VAR | <compound>
// <compound> ::= <functor> LPAR <args> RPAR
// <functor>  ::= ATOM
// <args>     ::= <term> | <term> COMMA <args>
//

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

var MyError = errors.New("invalid syntax for grammer G")

type DagParser struct {
	lex       *lexer
	peeking   bool
	peekToken *Token
	terms map[string]*Term
	termID map[*Term]int
	termIDGenerator int
}

func (d *DagParser) next() (*Token, error) {
	var token *Token
	if d.peeking {
		token = d.peekToken
		d.peeking = false
	} else {
		if tok, err := d.lex.next(); err != nil {
			return nil, err
		} else {
			token = tok
		}
	}
	return token, nil
}

// Eqivalent LL(1) Grammer
// <term> := ATOM <pars> | NUM | VAR
// <pars> := LPAR <args> RPAR | \epsilon
// <args> := <term> <otherArgs>
// <otherArgs> := COMMA <args> | \epsilon

// FIRST(<term>) = {ATOM, NUM, VAR}
// FIRST(<pars>) = {LPAR, \epsilon}
// FIRST(<args>) = {ATOM, NUM, VAR}
// FIRST(<otherArgs>) = {COMMA, \epsilon}

// FOLLOW(<term>) = {COMMA, RPAR, $}
// FOLLOW(<pars>) = {COMMA, RPAR, $}
// FOLLOW(<args>) = {RPAR}
// FOLLOW(<otherArgs>) = {RPAR}

// Parsing Table
// 				| ATOM				          | NUM/VAR			            | LPAR						 | RPAR  			  	   | COMMA              		 | $
// <term>       | <term> -> ATOM <pars>       | <term> -> NUM/VAR 			| ERROR 					 | ERROR 			  	   | ERROR              		 | ERROR
// <pars>		| ERROR				          | ERROR						| <pars> -> LPAR <args> RPAR | <pars> -> \epsilon 	   | <pars> -> \epsilon 		 | <pars> -> \epsilon
// <args>		| <args> -> <term><otherArgs> |	<args> -> <term><otherArgs> | ERROR						 | ERROR			       | ERROR			   			 | ERROR
// <otherArgs> 	| ERROR					      | ERROR						| ERROR						 | <otherArgs> -> \epsilon | <otherArgs> -> COMMA <args> | ERROR


func (p *DagParser) termNT() (*Term, error) {
	tok, err := p.next()
	if err != nil {
		return nil, err
	}

	switch tok.typ {
	// <term> -> NUM/VAR
	case tokenNumber:
		return p.mkTerminal(TermNumber, tok.literal), nil
	case tokenVariable:
		return p.mkTerminal(TermVariable, tok.literal), nil

	// <term> -> ATOM <pars>
	case tokenAtom:
		atom := p.mkTerminal(TermAtom, tok.literal)
		args, err := p.parsNT()
		if err != nil {
			return nil, err
		}
		if args != nil {
			return p.mkCompound(atom, args), nil
		}
		return atom, nil

	default:
		return nil, ErrParser
	}
}

func (p *DagParser) parsNT() ([]*Term, error) {
	tok, err := p.next()
	if err != nil {
		return nil, err
	}

	switch tok.typ {
	// <pars> -> \epsilon
	case tokenEOF, tokenComma, tokenRpar:
		p.peeking = true
		p.peekToken = tok
		return nil, nil

	// <pars> -> LPAR <args> RPAR
	case tokenLpar:
		args, err := p.argsNT()

		if err != nil {
			return nil, err
		}

		_, _ = p.next()
		return args, nil

	default:
		return nil, ErrParser
	}
}

func (p *DagParser) argsNT() ([]*Term, error) {
	term, err := p.termNT()
	if err != nil {
		return nil, err
	}

	otherArgs, err := p.otherArgsNT()
	if err != nil {
		return nil, err
	}

	return append([]*Term{term}, otherArgs...), nil
}

func (p *DagParser) otherArgsNT() ([]*Term, error) {
	tok, err := p.next()
	if err != nil {
		return nil, err
	}

	switch tok.typ {
	//	<otherArgs> -> \epsilon
	case tokenRpar:
		p.peeking = true
		p.peekToken = tok
		return nil, nil

	// <otherArgs> -> COMMA <args>
	case tokenComma:
		return p.argsNT()

	default:
		return nil, ErrParser
	}
}

func (p *DagParser) mkTerminal(typ TermType, lit string) *Term {
	if term, ok := p.terms[lit]; ok {
		return term
	}
	term := &Term{Typ: typ, Literal: lit}
	p.insertTerm(lit, term)
	return term
}

func (p *DagParser) mkCompound(functor *Term, args []*Term) *Term {
	key := strconv.Itoa(p.termID[functor])
	for _, arg := range args {
		key += strconv.Itoa(p.termID[arg])
	}

	if term, ok := p.terms[key]; ok {
		return term
	}

	term := &Term{Typ: TermCompound, Functor: functor, Args: args}
	p.insertTerm(key, term)
	return term
}

func (p *DagParser) insertTerm(key string, term *Term) {
	p.terms[key] = term
	p.termID[term] = p.termIDGenerator
	p.termIDGenerator++
}

func (d *DagParser) Parse(s string) (*Term, error) {
	if len(s) == 0 {
		return nil, nil
	}
	d.lex = newLexer(s)

	if term, err := d.termNT(); err != nil {
		return nil, err
	} else {
		if tok, err := d.next(); err != nil {
			return nil, err
		} else if tok.typ != tokenEOF{
			return nil, MyError
		} else {
			return term, nil
		}
	}
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &DagParser{peeking: false, terms: make(map[string]*Term), termID: make(map[*Term]int), termIDGenerator: 0}
}
