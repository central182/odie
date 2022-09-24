package entry

type LexicalCategory interface {
	Name() string
}

type Adjective struct{}

func (a Adjective) Name() string {
	return "adjective"
}

type Adverb struct{}

func (a Adverb) Name() string {
	return "adverb"
}

type CombiningForm struct{}

func (c CombiningForm) Name() string {
	return "combining form"
}

type Conjunction struct{}

func (c Conjunction) Name() string {
	return "conjunction"
}

type Contraction struct{}

func (c Contraction) Name() string {
	return "contraction"
}

type Determiner struct{}

func (d Determiner) Name() string {
	return "determiner"
}

type Idiomatic struct{}

func (i Idiomatic) Name() string {
	return "idiomatic"
}

type Interjection struct{}

func (i Interjection) Name() string {
	return "interjection"
}

type Noun struct{}

func (n Noun) Name() string {
	return "noun"
}

type Numeral struct{}

func (n Numeral) Name() string {
	return "numeral"
}

type Other struct{}

func (o Other) Name() string {
	return "other"
}

type Particle struct{}

func (p Particle) Name() string {
	return "particle"
}

type Predeterminer struct{}

func (p Predeterminer) Name() string {
	return "predeterminer"
}

type Prefix struct{}

func (p Prefix) Name() string {
	return "prefix"
}

type Preposition struct{}

func (p Preposition) Name() string {
	return "preposition"
}

type Pronoun struct{}

func (p Pronoun) Name() string {
	return "pronoun"
}

type Residual struct{}

func (r Residual) Name() string {
	return "residual"
}

type Suffix struct{}

func (s Suffix) Name() string {
	return "suffix"
}

type Verb struct{}

func (v Verb) Name() string {
	return "verb"
}

func newLexicalCategory(name string) (LexicalCategory, NewLexicalCategoryError) {
	for _, lc := range []LexicalCategory{
		Adjective{},
		Adverb{},
		CombiningForm{},
		Conjunction{},
		Contraction{},
		Determiner{},
		Idiomatic{},
		Interjection{},
		Noun{},
		Numeral{},
		Other{},
		Particle{},
		Predeterminer{},
		Prefix{},
		Preposition{},
		Pronoun{},
		Residual{},
		Suffix{},
		Verb{},
	} {
		if lc.Name() == name {
			return lc, nil
		}
	}

	err := &newLexicalCategoryError{}
	err.setHasUnknownLexicalCategory()
	return nil, err
}
