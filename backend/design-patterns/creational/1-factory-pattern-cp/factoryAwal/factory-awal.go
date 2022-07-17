package factoryAwal

import "fmt"

// Nah disini kita akan belajar menggunakan konsep factory
// Disini ktia akan membuat contoh penggunaan factory pattern di Go]

// Referensi nya disini ya https://refactoring.guru/design-patterns/factory-method

// Misalnya nih kita mau buat sebuah Content
// Kita bisa mulai dari interface nya, nah setiap kontent dapat di Play() nih

type Content interface {
	Play()
}

// ini adalah Concrete implementation seperti pada contoh diagram pada refrensi ConcreteProductA
type AnimeContent struct {
}

func (a AnimeContent) Play() {
	fmt.Println("Anime content is playing")
}

// ConcreteProductB
type KoreanDramaContent struct {
}

func (k KoreanDramaContent) Play() {
	fmt.Println("Korean Drama content is playing")
}

// Ini adalah Creater,
type ContentCreator interface {
	Produce() Content
}

// ConcreteCreatorA
type Mappa struct {
}

func (m *Mappa) Produce() Content {
	return &AnimeContent{}
}

// ConcreteCreatorB
type NetflixKorea struct {
}

func (n *NetflixKorea) Produce() Content {
	return &KoreanDramaContent{}
}
