package factoryLanjutan

import (
	"fmt"
	"time"
)

var (
	Sunday  time.Time
	Monday  time.Time
	Tuesday time.Time
)

func init() {
	Sunday = time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC)
	Monday = time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC)
	Tuesday = time.Date(2022, time.March, 22, 0, 0, 0, 0, time.UTC)
}

// Nah disini kan kalian udah membuatkan boiler plate nya, tapi sebelumnya itu merupakan contoh
// yang sederhana.
// Disini kita coba bermain main pada Fungsional Produce nya, coba kita tambahkan parameter waktu
// Jadi setiap Creator itu menghasilkan content berbeda beda tiap hari

type Content interface {
	Play() string
}

// ini adalah Concrete implementation seperti pada contoh diagram pada refrensi ConcreteProductA
type AnimeContent struct {
	Name string
}

func (a AnimeContent) Play() string {
	fmt.Println("Anime content is playing: ", a.Name)
	return a.Name
}

// ConcreteProductB
type KoreanDramaContent struct {
	Name string
}

func (k KoreanDramaContent) Play() string {
	fmt.Println("Korean Drama content is playing: ", k.Name)
	return k.Name
}

// Ini adalah Creator,
// Disini kita menambahkan time.Time pada Produce() Content
type ContentCreator interface {
	Produce(time.Time) Content
}

// ConcreteCreatorA
type Mappa struct {
}

// Gunakan time.Weekday().String() untuk menentukan hari apa saja yang kita inginkan
// Nah disini kita menambahkan sebuah logic pada Produce ketika hari hari tertentu

// Mappa memproduce Anime berjudul "Attack on Titan" pada hari "Sunday" dan "Jujutsu Kaisen" pada hari "Monday"
func (m *Mappa) Produce(time time.Time) Content {
	return nil // TODO: replace this
}

// ConcreteCreatorB
type NetflixKorea struct {
}

// Mappa memproduce Anime berjudul "All of Us Are Dead" pada hari "Sunday" dan
// "Our Beloved Summer" pada hari "Monday"
// "Start Up" pada hari "Tuesday"
func (n *NetflixKorea) Produce(time time.Time) Content {
	return nil // TODO: replace this
}
