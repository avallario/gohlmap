package mapio

const (
	LBRACE   = iota
	RBRACE   = iota
	LPAREN   = iota
	RPAREN   = iota
	LBRACK   = iota
	RBRACK   = iota
	COMMA    = iota
	NUMINT   = iota
	NUMFLOAT = iota
	STRING   = iota
	EOF      = iota
	ERROR    = iota
)

func KindName(kind int) string {
	switch kind {
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACK:
		return "LBRACK"
	case RBRACK:
		return "RBRACK"
	case COMMA:
		return "COMMA"
	case NUMINT:
		return "INT"
	case NUMFLOAT:
		return "FLOAT"
	case STRING:
		return "STRING"
	case EOF:
		return "EOF"
	case ERROR:
		return "ERROR"
	default:
		return "UNDEFINED KIND"
	}
}

type Token struct {
	Kind     int
	Spelling string
}
