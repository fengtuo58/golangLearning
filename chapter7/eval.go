type Var string

type literal float64

type unary struct {
	op rune
	x Expr
}

type binary struct {
	op rune
	x Expr
}

type call struct {
	fn string
	args[]Expr
}

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64 
}

func TestEval(t *testing.T) {
	tests := []struct{
		expr string
		env Env
		want string
	} {
		{"sqrt(A/pi)", Env{"A":87616, "pi":math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x":12, "y":1}, "1729"},
		{"5/9*(F-32)", Env{"F":212}, "100"},	
	}
}

var preExpr string
for _, test := range tests {
	if test.expr != preExpr {
		fmt.Printf("\n%s\n", test.expr)
		preExpr = test.Expr
	}
	
}
