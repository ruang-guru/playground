package pintu

type Opener interface {
	Open()
}

type Pintu struct {
}

func (p Pintu) Open() {
	println("Pintu dibuka")
}
