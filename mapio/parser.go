package mapio

import (
	"github.com/avallario/gohlmap/maptree"
	"math/big"
	"strconv"
)

type parseError struct {
	msg string
}

func (p parseError) Error() string {
	return "Parsing error: " + p.msg
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type parser struct {
	input         *scanner
	current_token Token
}

func NewParser(s *scanner) *parser {
	p := new(parser)
	p.input = s
	return p
}

func (p *parser) Parse() *maptree.HLMap {
	p.nextToken()
	hlmap := new(maptree.HLMap)

	for p.current_token.Kind == LBRACE {
		hlmap.Entitylist = append(hlmap.Entitylist, p.parseEntity())
	}

	return hlmap
}

func (p *parser) parseEntity() *maptree.Entity {
	entity := new(maptree.Entity)
	entity.Properties = make(map[string]string)

	p.accept(LBRACE)

	for p.current_token.Kind == STRING {
		key := p.current_token.Spelling
		p.nextToken()

		value := p.current_token.Spelling
		p.accept(STRING)

		entity.Properties[key] = value
	}

	for p.current_token.Kind == LBRACE {
		entity.Brushlist = append(entity.Brushlist, p.parseBrush())
	}

	p.accept(RBRACE)

	return entity
}

func (p *parser) parseBrush() *maptree.Brush {
	brush := new(maptree.Brush)

	p.accept(LBRACE)

	for p.current_token.Kind == LPAREN {
		brush.Facelist = append(brush.Facelist, p.parseFace())
	}

	p.accept(RBRACE)

	return brush
}

func (p *parser) parseFace() *maptree.Face {
	face := new(maptree.Face)

	p.accept(LPAREN)
	face.X1 = p.parseInt()
	face.Y1 = p.parseInt()
	face.Z1 = p.parseInt()
	p.accept(RPAREN)

	p.accept(LPAREN)
	face.X2 = p.parseInt()
	face.Y2 = p.parseInt()
	face.Z2 = p.parseInt()
	p.accept(RPAREN)

	p.accept(LPAREN)
	face.X3 = p.parseInt()
	face.Y3 = p.parseInt()
	face.Z3 = p.parseInt()
	p.accept(RPAREN)

	face.Texname = p.current_token.Spelling
	p.accept(STRING)

	p.accept(LBRACK)
	face.TX1 = p.parseRat()
	face.TY1 = p.parseRat()
	face.TZ1 = p.parseRat()
	face.TOffset1 = p.parseInt()
	p.accept(RBRACK)

	p.accept(LBRACK)
	face.TX2 = p.parseRat()
	face.TY2 = p.parseRat()
	face.TZ2 = p.parseRat()
	face.TOffset2 = p.parseInt()
	p.accept(RBRACK)

	face.Rot = p.parseRat()

	face.ScaleX = p.parseRat()
	face.ScaleY = p.parseRat()

	return face
}

func (p *parser) parseInt() int {
	var val int

	if p.current_token.Kind == NUMINT {
		val64, err := strconv.ParseInt(p.current_token.Spelling, 10, 64)
		check(err)
		val = int(val64)
		p.nextToken()
	} else if p.current_token.Kind == NUMFLOAT {
		flt, err := strconv.ParseFloat(p.current_token.Spelling, 64)
		check(err)
		val = int(flt)
		p.nextToken()
	} else {
		panic(parseError{"Expected NUM but found " + KindName(p.current_token.Kind)})
	}

	return val
}

func (p *parser) parseRat() *big.Rat {
	rat := big.NewRat(1, 1)

	if p.current_token.Kind == NUMINT {
		val, err := strconv.ParseInt(p.current_token.Spelling, 10, 64)
		check(err)
		rat.SetInt64(val)
		p.nextToken()
	} else if p.current_token.Kind == NUMFLOAT {
		val, err := strconv.ParseFloat(p.current_token.Spelling, 64)
		check(err)
		rat.SetFloat64(val)
		p.nextToken()
	} else {
		panic(parseError{"Expected NUM but found " + KindName(p.current_token.Kind)})
	}

	return rat
}

func (p *parser) nextToken() {
	var err error
	p.current_token, err = p.input.Scan()
	check(err)
}

func (p *parser) accept(token_kind int) {
	if p.current_token.Kind == token_kind {
		p.nextToken()
	} else {
		panic(parseError{"Expected " + KindName(token_kind) + " but found " + KindName(p.current_token.Kind)})
	}
}
