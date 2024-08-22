package design

type chiguiro struct {
}

var chiguiroInstance *chiguiro

func GetChiguiroInstance() *chiguiro {
	if chiguiroInstance == nil {
		chiguiroInstance = &chiguiro{}
	}
	return chiguiroInstance
}
