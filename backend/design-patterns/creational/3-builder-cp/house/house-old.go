package house

// Contoh ketika membuat House, disini kita tidak menerapkan design pattern builder
type HouseOld struct {
	numOfWindows    int
	numOfDoors      int
	hasGarage       bool
	hasSwimmingPool bool
}

// Nah misalnya kita mau buat object House baru nih. Apakah semua komponen yang kita butuhkan harus kita buat disini `NewHouse()`
// Bayangkan klo kalian mau buat NewHouse dengan properti baru. Nah kalian bakal nambah parameter di NewHouse() lagi kan ?
// Jadi bukan hanya itu saja yang dirubah tapi semua kode yang memakai NewHouse harus dirubah juga
// Belum lagi kalau ada logic ketika membuat sebuah parameter tersebut. Apakah mau taruh di NewHouse() juga ?
// Builder pattern menyelesaikan hal tersebut
func NewHouse(numOfWindows int, numOfDoors int, hasgarage bool, hasSwimmingPool bool) *HouseOld {
	return &HouseOld{
		numOfWindows:    numOfWindows,
		numOfDoors:      numOfDoors,
		hasGarage:       hasgarage,
		hasSwimmingPool: hasSwimmingPool,
	}
}
