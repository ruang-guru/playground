package product

type IPhone13 struct{}

func (l *IPhone13) Price() float64 {
	return 17999999
}

func (l *IPhone13) Is5G() bool {
	return true
}

func (l *IPhone13) IsFoldable() bool {
	return false
}
