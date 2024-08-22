package design

type chiguiroBuilder struct {
	chiguiro *chiguiro
}

func NewChiguiroBuilder() *chiguiroBuilder {
	return &chiguiroBuilder{}

}

func (ch *chiguiroBuilder) weight(weight int) *chiguiroBuilder {
	ch.chiguiro.weight = weight
	return ch
}

func (ch *chiguiroBuilder) name(name string) *chiguiroBuilder {
	ch.chiguiro.name = name
	return ch
}
func (ch *chiguiroBuilder) age(age int) *chiguiroBuilder {
	ch.chiguiro.age = age
	return ch
}

func (ch *chiguiroBuilder) build() *chiguiro {
	return ch.chiguiro
}
