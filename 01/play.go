package main

import (
	"fmt"
)

type TokenKind int

type Token interface {
	Kind() TokenKind
	Text() string
}

type TokenEle struct {
	kind TokenKind
	text string
}

func (t *TokenEle) Kind() TokenKind {
	return t.kind
}

func (t *TokenEle) Text() string {
	return t.text
}

var tokenArray = []Token{
	&TokenEle{Keyword, "function"},
	&TokenEle{Identifier, "sayHello"},
	&TokenEle{Seperator, "("},
	&TokenEle{Seperator, ")"},
	&TokenEle{Identifier, "println"},
	&TokenEle{Seperator, "("},
	&TokenEle{StringLiteral, "Hello world!"},
	&TokenEle{Seperator, ")"},
	&TokenEle{Seperator, ";"},
	&TokenEle{Seperator, "}"},
	&TokenEle{Identifier, "sayHello"},
	&TokenEle{Seperator, "("},
	&TokenEle{Seperator, ")"},
	&TokenEle{Seperator, ";"},
	&TokenEle{EOF, ""},
}

const (
	Keyword = iota
	Identifier
	StringLiteral
	Seperator
	Operator
	EOF
)

// simulate static oop
type VirtualStatic struct{}

var v = new(VirtualStatic)

func (v *VirtualStatic) Dump(string) {}
func isStatementNode(node interface{}) bool {
	if _, ok := node.(Statement); ok {
		return true
	}
	return false
}

func isFunctionBodyNode(node interface{}) bool {
	_, ok := node.(FunctionBody)
	return ok
}

func isFunctionCallNode(node interface{}) bool {
	_, ok := node.(FunctionCall)
	return ok
}

type AstNode interface {
	Dump(prefix string)
}

type Statement interface {
	AstNode
	State()
}

// 程序根节点
type Prog struct {
	stmts []Statement
}

func (p *Prog) Dump(prefix string) {
	fmt.Println(prefix + "Prog")
	for _, s := range p.stmts {
		s.Dump(prefix + "\t")
	}
}

func NewProg(stmts []Statement) *Prog {
	return &Prog{
		stmts: stmts,
	}
}

// 函数声明节点
type FunctionDecl struct {
	name string
	body *FunctionBody
}

func NewFunctionDecl(name string, body *FunctionBody) *FunctionDecl {
	return &FunctionDecl{
		name: name,
		body: body,
	}
}

func (fdcl *FunctionDecl) Dump(prefix string) {
	fmt.Println(prefix + "FunctionDecl " + fdcl.name)
	fdcl.body.Dump(prefix + "\t")
}

func (fdcl *FunctionDecl) State() {}

// 函数体
type FunctionBody struct {
	stmts []Statement
}

func NewFunctionBody(stmts []Statement) *FunctionBody {
	return &FunctionBody{
		stmts: stmts,
	}
}

func (fbody *FunctionBody) Dump(prefix string) {
	fmt.Println(prefix + "FunctionBody")
	for _, s := range fbody.stmts {
		s.Dump(prefix + "\t")
	}
}

// 函数调用
type FunctionCall struct {
	name       string
	parameters []string
	definition *FunctionDecl
}

func NewFunctionCall(name string, params []string) *FunctionCall {
	return &FunctionCall{
		name:       name,
		parameters: params,
	}
}

func (fc *FunctionCall) Dump(prefix string) {
	def := ", not resolved"
	if fc.definition != nil {
		def = ", resolved"
	}
	fmt.Println(prefix + "FunctionCall " + fc.name + def)
	for _, s := range fc.parameters {
		fmt.Println(prefix + "\t" + "Parameter: " + s)
	}
}

func (fc *FunctionCall) State() {}

type Tokenizer struct {
	tokens []Token
	pos    int
}

func NewTokenizer(tokens []Token) *Tokenizer {
	return &Tokenizer{
		tokens: tokens,
	}
}

func (t *Tokenizer) Next() Token {
	var res Token
	if t.pos <= len(t.tokens) {
		res = t.token[t.pos]
		t.pos++
		return res
	} else {
		return t.token[t.pos]
	}
}

func (t *Tokenizer) Position() int {
	return t.pos
}

func (t *Tokenizer) TraceBack(pos int) {
	t.pos = pos
}

// Parser
type Parser struct {
	tokenizer *Tokenizer
}

func NewParser(t *Tokenizer) *Parser {
	return &Tokenizer{
		tokenizer: t,
	}
}

// prog = (functionDecl | functionCall) *;
func (p *Parser) parseProg() *Prog {
	var stmt Statement
	var stmts = []Statement{}
	for {
		stmt = p.parseFunctionDecl()
		if isStatementNode(stmt) {
			stmts = append(stmts, stmt)
			continue
		}

		stmt = p.parseFuntionCall()
		if isStatementNode(stmt) {
			stmts = append(stmts, stmt)
			continue
		}

		// TODO: check
		if stmt == nil {
			break
		}
	}

	return NewProg(stmts)
}

// functionDecl: "function" Identifer "(" ")" functionBody
func (p *Parser) parseFunctionDecl() *FunctionDecl {
	oldPos = p.tokenizer.Position()
	t := p.Tokenizer.Next()
	if t.Kind() == Keyword && t.Text() == "function" {
		t = p.tokenizer.Next()
		if t.Kind() == Identifier {
			t1 := p.tokenizer.Next()
			if t1.Text() == "(" {
				t2 := p.tokenizer.Next()
				if t2.Text() == ")" {
					functionBody := p.parseFunctionBody
					if isFunctionBodyNode(functionBody) {
						return NewFuncionDecl(t.Text(), functionBody)
					}
				} else {
					fmt.Println("Expecting ')' in FunctionDecl, while we got a " + t2.Text())
					return
				}
			} else {
				fmt.Println("Expecting '(' in FunctionDecl, while we got a " + t1.Text())
				return
			}
		}

		p.tokenizer.TraceBack(oldPos)
		return nil
	}
}

func (p *Parser) parseFunctionBody() *FunctionBody {
	return nil
}

func (p *Parser) parseFunctionCall() *FunctionCall {
	return nil
}

func main() {
	fmt.Println("test")
}
