package design

var chiguiroInstance *chiguiro

func NewChiguiroInstance() *chiguiro {
	if chiguiroInstance == nil {
		chiguiroInstance = &chiguiro{}
	}
	return chiguiroInstance
}

func (chiguiro *chiguiro) SetWeight(weight int) {
	chiguiro.weight = weight
}

func (chiguiro *chiguiro) GetWeight() int {
	return chiguiro.weight
}
