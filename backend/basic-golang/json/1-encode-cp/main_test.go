package main

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	var leaderboard Leaderboard
	var expected []byte
	BeforeEach(func() {
		users := []*UserRank{
			{
				Name:  "Roger",
				Email: "roger@ruangguru.com",
				Rank:  1,
			},
			{
				Name:  "Tony",
				Email: "tony@ruangguru.com",
				Rank:  2,
			},
			{
				Name:  "Bruce",
				Email: "bruce@ruangguru.com",
				Rank:  3,
			},
			{
				Name:  "Natasha",
				Email: "natasha@ruangguru.com",
				Rank:  4,
			},
			{
				Name:  "Clint",
				Email: "clint@ruangguru.com",
				Rank:  5,
			},
		}

		leaderboard = Leaderboard{
			Users: users,
		}
		expected, _ = json.Marshal(leaderboard)
	})
	Describe("EncodeToJson", func() {
		It("returns json", func() {
			jsonByte, err := EncodeToJson(leaderboard)
			jsonString := string(jsonByte)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(jsonByte).To(Equal(expected))
			Expect(jsonString).To(ContainSubstring("name"))
			Expect(jsonString).To(ContainSubstring("rank"))
			Expect(jsonString).ToNot(ContainSubstring("email"))
		})
	})
})
