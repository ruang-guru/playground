package product

type SamsungS22 struct{}

func (l *SamsungS22) Price() float64 {
	return 17999999
}

func (l *SamsungS22) Is5G() bool {
	return true
}

func (l *SamsungS22) IsFoldable() bool {
	return false
}
