package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	var leaderboardObject Leaderboard
	BeforeEach(func() {
		users := []*UserRank{
			{
				Name: "Roger",
				Rank: 1,
			},
			{
				Name: "Tony",
				Rank: 2,
			},
			{
				Name: "Bruce",
				Rank: 3,
			},
			{
				Name: "Natasha",
				Rank: 4,
			},
			{
				Name: "Clint",
				Rank: 5,
			},
		}

		leaderboardObject = Leaderboard{
			Users: users,
		}
	})
	Describe("DecodeToLeaderboard", func() {
		It("returns leaderboard object", func() {
			jsonData := `{
				"Users":[
				   {
					  "Name":"Roger",
					  "Email":"roger@ruangguru.com",
					  "Rank":1
				   },
				   {
					  "Name":"Tony",
					  "Email":"tony@ruangguru.com",
					  "Rank":2
				   },
				   {
					  "Name":"Bruce",
					  "Email":"bruce@ruangguru.com",
					  "Rank":3
				   },
				   {
					  "Name":"Natasha",
					  "Email":"natasha@ruangguru.com",
					  "Rank":4
				   },
				   {
					  "Name":"Clint",
					  "Email":"clint@ruangguru.com",
					  "Rank":5
				   }
				]
			 }`
			leaderboard, err := DecodeToLeaderboard([]byte(jsonData))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(leaderboard.Users).To(HaveLen(5))
			for idx, l := range leaderboard.Users {
				Expect(l.Name).To(Equal(leaderboardObject.Users[idx].Name))
				Expect(l.Rank).To(Equal(leaderboardObject.Users[idx].Rank))
				Expect(l.Email).To(BeEmpty())
			}
		})
		It("returns error when given invalid json", func() {
			jsonData := `{
				"Users":[
				   {
					  "Name":"Roger",
					  "Email":"roger@ruangguru.com",
					  "Rank":1
				   },
				   {
					  "Name":"Tony",
					  "Email":"tony@ruangguru.com",
					  "Rank":2
				   },
				   {
					  "Name":"Bruce",
					  "Email":"bruce@ruangguru.com",
					  "Rank":3
				   },
				   {
					  "Name":"Natasha",
					  "Email":"natasha@ruangguru.com",
					  "Rank":4
				   },
				   {
					  "Name":"Clint",
					  "Email":"clint@ruangguru.com",
					  "Rank":5
				   }
				]
			 `
			leaderboard, err := DecodeToLeaderboard([]byte(jsonData))
			Expect(err).Should(HaveOccurred())
			Expect(leaderboard.Users).To(BeNil())
		})
	})
})
