package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface { //プログラム内の文を表現するインターフェイス
	Node
	statementNode()
}

type Expression interface { //プログラム内の式を表現するインターフェイス
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Indentifier
	Value Expression
}

func (ls *LetStatement) expressionNode()      {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Indentifier struct {
	Token token.Token
	Value string
}

func (i *Indentifier) expressionNode()      {}
func (i *Indentifier) TokenLiteral() string { return i.Token.Literal }
