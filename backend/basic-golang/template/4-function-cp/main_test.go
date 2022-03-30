package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	var leaderboardObject Leaderboard
	var expectedOutput string
	BeforeEach(func() {
		users := []*UserRank{
			{
				Name:  "Roger",
				Rank:  1,
				Score: 1000,
			},
			{
				Name:  "Tony",
				Rank:  2,
				Score: 900,
			},
			{
				Name:  "Bruce",
				Rank:  3,
				Score: 800,
			},
			{
				Name:  "Natasha",
				Rank:  4,
				Score: 700,
			},
			{
				Name:  "Clint",
				Rank:  5,
				Score: 600,
			},
		}

		leaderboardObject = Leaderboard{
			Users: users,
		}
		expectedOutput = `Roger: 1000Tony: 900Bruce: 800Natasha: 700Clint: 600Total Score: 4000`
	})
	Describe("ExecuteToByteBuffer", func() {
		It("returns slice of bytes with correct wording", func() {
			b, err := ExecuteToByteBuffer(leaderboardObject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(b)).To(Equal(expectedOutput))
		})
	})
})
