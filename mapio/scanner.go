package mapio

import (
	"bufio"
	"io"
	"os"
)

type scanError struct {
	msg string
}

func (s scanError) Error() string {
	return "Scanning error: " + s.msg
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '\r'
}

func isDigit(r rune) bool {
	return r >= 48 && r <= 57
}

type scanner struct {
	fstream      *bufio.Reader
	current_char rune
	paren_count  int
	eof          bool
}

func NewScanner(file *os.File) *scanner {
	s := new(scanner)
	s.current_char = ' '
	s.paren_count = 0
	s.eof = false
	s.fstream = bufio.NewReader(file)
	return s
}

func (s *scanner) Scan() (Token, error) {
	var err error = nil

	for isWhiteSpace(s.current_char) {
		s.nextChar()
	}

	// Scan in EOF

	if s.eof {
		return Token{EOF, "EOF"}, err
	}

	// Every third RPAREN must be followed by a TEXNAME
	// Scanner must take this into account because a TEXNAME may contain any character

	if s.paren_count == 3 {
		s.paren_count = 0
		spelling := ""

		for !isWhiteSpace(s.current_char) {
			spelling += string(s.current_char)
			s.nextChar()
		}

		return Token{STRING, spelling}, err
	}

	// Scan in symbol
	switch s.current_char {
	case '{':
		s.nextChar()
		return Token{LBRACE, "{"}, err
	case '}':
		s.nextChar()
		return Token{RBRACE, "}"}, err
	case '(':
		s.nextChar()
		return Token{LPAREN, "("}, err
	case ')':
		s.paren_count++
		s.nextChar()
		return Token{RPAREN, ")"}, err
	case '[':
		s.nextChar()
		return Token{LBRACK, "["}, err
	case ']':
		s.nextChar()
		return Token{RBRACK, "]"}, err
	case ',':
		s.nextChar()
		return Token{COMMA, ","}, err
	}

	// Scan in string
	if s.current_char == '"' {
		// Accumulate characters until next "
		spelling := ""
		s.nextChar()

		for s.current_char != '"' {
			spelling += string(s.current_char)
			s.nextChar()
		}
		s.nextChar()

		return Token{STRING, spelling}, err
	}

	// Scan in number (must be number at this point, so return ERROR if it doesn't fit NUM format)

	spelling := ""

	if s.current_char == '-' {
		spelling += "-"
		s.nextChar()
	}

	if !isDigit(s.current_char) {
		err = scanError{"Expected digit but found: '" + string(s.current_char) + "'"}
		return Token{ERROR, "ERROR"}, err
	}

	for isDigit(s.current_char) {
		spelling += string(s.current_char)
		s.nextChar()
	}

	if s.current_char == '.' {
		dec_spelling := "."
		found_nonzero := false
		s.nextChar()

		if !isDigit(s.current_char) {
			err = scanError{"Expected digit but found: '" + string(s.current_char) + "'"}
			return Token{ERROR, "ERROR"}, err
		}

		for isDigit(s.current_char) {
			if s.current_char != '0' {
				found_nonzero = true
			}
			dec_spelling += string(s.current_char)
			s.nextChar()
		}

		if s.current_char == 'e' {
			found_nonzero = true
			dec_spelling += "e"
			s.nextChar()

			if s.current_char == '-' {
				dec_spelling += "-"
				s.nextChar()
			}

			if !isDigit(s.current_char) {
				err = scanError{"Expected digit but found: '" + string(s.current_char) + "'"}
				return Token{ERROR, "ERROR"}, err
			}

			for isDigit(s.current_char) {
				dec_spelling += string(s.current_char)
				s.nextChar()
			}
		}

		// If there are only trailing zeroes, read as NUMINT
		if found_nonzero {
			return Token{NUMFLOAT, spelling + dec_spelling}, err
		} else {
			return Token{NUMINT, spelling}, err
		}

	} else {
		return Token{NUMINT, spelling}, err
	}
}

func (s *scanner) nextChar() {
	var err error = nil
	s.current_char, _, err = s.fstream.ReadRune()

	if err == io.EOF {
		s.eof = true
	} else if err != nil {
		panic(err)
	}
}
