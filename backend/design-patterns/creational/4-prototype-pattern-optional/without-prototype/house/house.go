package house

// Bayangkan kalian mempunyai sebuah object House yang memiliki beberapa properti seperti ini.
// Nah kalian ingin membuat Copy dari House tersebut.
// Tapi apakah bisa semua nya kita copy ?
// Kalian bisa lihat disitu ada 2 properti yang private. Yang tidak bisa kita copy.
// Prototpye pattern menyelesaikan permasalah seperti ini.
type House struct {
	NumOfWindows int
	NumOfDoors   int
	typeOfJamban string
	typeOfLamp   string
}
