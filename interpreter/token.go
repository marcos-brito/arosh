package interpreter

type tokenType string

type Token struct {
	t      tokenType
	lexeme string
}

const (
	EOF  = "EOF"
	WORD = "WORD"
	NAME = "NAME"

	IO_NUMBER = "IO_NUMBER"

	// Operators

	AND    = "&"
	PIPE   = "|"
	DOTDOT = ".."
	DOLLAR = "$"

	DAND  = "&&"
	DPIPE = "||"

	SEMI = ";"

	DSEMI    = ";;"
	DLESS    = "<<"
	LESS     = "<"
	GREAT    = ">"
	DGREAT   = ">>"
	LESSAND  = "<&"
	GREATAND = ">&"

	LESSGREAT = "<>"
	DLESSDASH = "<<-"
	CLOBBER   = ">|"

	// Reserved words

	IF   = "IF"
	THEN = "THEN"
	ELSE = "ELSE"
	ELIF = "ELIF"
	FI   = "FI"

	DO   = "DO"
	DONE = "DONE"

	CASE = "CASE"
	ESAC = "ESAC"

	WHILE = "WHILE"
	UNTIL = "UNTIL"
	FOR   = "FOR"

	LBRACE = "{"
	RBRACE = "}"
	BANG   = "!"
	IN     = "IN"
)

var reservedWords = map[string]string{
	"do":    DO,
	"done":  DONE,
	"if":    IF,
	"then":  THEN,
	"else":  ELSE,
	"elif":  ELIF,
	"fi":    FI,
	"case":  CASE,
	"esac":  ESAC,
	"while": WHILE,
	"until": UNTIL,
	"for":   FOR,
	"in":    IN,
}

func LookupReservedWord(word string) (string, bool) {
	word, ok := reservedWords[word]

	if ok {
		return word, true
	}

	return "", false
}
