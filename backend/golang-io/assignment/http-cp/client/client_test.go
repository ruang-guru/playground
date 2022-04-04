package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	client "github.com/ruang-guru/playground/backend/golang-io/assignment/http-cp/client"
)

var _ = Describe("Client", func() {
	It("should call https://pokeapi.co/api/v2/pokemon/1 and returns the Pokemon data", func() {
		pokemon, err := client.GetPokemonData()
		Expect(err).To(Not(HaveOccurred()))

		Expect(pokemon.Name).To(Equal("bulbasaur"))
		Expect(pokemon.Abilities[0].Ability.Name).To(Equal("overgrow"))
		Expect(pokemon.Abilities[1].Ability.Name).To(Equal("chlorophyll"))
		Expect(pokemon.Species.Name).To(Equal("bulbasaur"))
	})
})
