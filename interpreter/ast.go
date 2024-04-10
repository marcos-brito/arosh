package interpreter

type Node interface {
	Node()
}

type Program struct {
	nodes []Node
}

func (*Program) Node() {}

type Subshell struct {
	nodes []Node
}

func (*Subshell) Node() {}

type Sequence struct {
	lhs       Node
	rhs       Node
	separator string
}

func (*Sequence) Node() {}

type Conditional struct {
	lhs             Node
	rhs             Node
	conditionalType string
}

func (*Conditional) Node() {}

type Pipe struct {
	lhs Node
	rhs Node
}

func (*Pipe) Node() {}

type SimpleCommand struct {
	name        string
	params      []string
	redirection Redirection
}

func (*SimpleCommand) Node() {}

type Redirection struct {
	ioNumber        int
	redirectionType string
	file            string
}

func (*Redirection) Node() {}
