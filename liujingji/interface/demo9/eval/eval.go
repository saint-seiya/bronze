package eval

import (
	"fmt"
	"math"
)

type Expr interface {
	// Eval返回表达式在env上下文下的值
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Var string

type literal float64

// 一元操作符
type unary struct {
	// + - * /中的一个
	op rune
	x  Expr
}

// 二元操作符
type binary struct {
	// + - * /中的一个
	op   rune
	x, y Expr
}

// 函数调用表达式, sin(x)
type call struct {
	// pow sin sqrt中的一个
	fn   string
	args []Expr
}

// 对包含变量的表达式进行求值, 需要一个上下文(environment)来把变量映射到数值
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsuppored binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
