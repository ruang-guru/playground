package pintu

import "fmt"

type SmartGagang struct {
	opener Opener
}

func NewSmartGagang(opener Opener) *SmartGagang {
	return &SmartGagang{opener: opener}
}

func (s SmartGagang) Open() {
	fmt.Println("this is electronic key")
	ok := s.ConnectToWifi()
	if ok {
		s.opener.Open()
	}
}

func (s SmartGagang) ConnectToWifi() bool {
	fmt.Println("connecting to wifi")
	return true
}
