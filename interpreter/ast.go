package interpreter

import "fmt"

type Node interface {
	Node()
	String() string
}

type Program struct {
	nodes []Node
}

func (*Program) Node() {}
func (s *Program) String() string {
	str := ""

	for _, node := range s.nodes {
		str += node.String()
	}

	return str
}

type Subshell struct {
	nodes []Node
}

func (*Subshell) Node() {}
func (s *Subshell) String() string {
	str := ""

	for _, node := range s.nodes {
		str += node.String()
	}

	return str
}

type Sequence struct {
	lhs       Node
	rhs       Node
	separator string
}

func (*Sequence) Node() {}
func (s *Sequence) String() string {
	return fmt.Sprintf("(sequence %s %s %s)", s.separator, s.lhs, s.rhs)
}

type Conditional struct {
	lhs             Node
	rhs             Node
	conditionalType string
}

func (*Conditional) Node() {}
func (c *Conditional) String() string {
	return fmt.Sprintf("(conditional %s %s %s)", c.conditionalType, c.lhs, c.rhs)
}

type Pipe struct {
	lhs Node
	rhs Node
}

func (*Pipe) Node() {}
func (p *Pipe) String() string {
	return fmt.Sprintf("(pipe %s %s)", p.lhs, p.rhs)
}

type SimpleCommand struct {
	name        string
	params      []string
	redirection Node
}

func (*SimpleCommand) Node() {}
func (s *SimpleCommand) String() string {
	return fmt.Sprintf("(simpleCommand %s %v %s)", s.name, s.params, s.redirection)
}

type Redirection struct {
	ioNumber        int
	redirectionType string
	file            string
}

func (*Redirection) Node() {}
func (r *Redirection) String() string {
	return fmt.Sprintf("(redirection %d %s %s)", r.ioNumber, r.redirectionType, r.file)
}
