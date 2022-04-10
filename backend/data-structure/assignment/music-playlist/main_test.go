package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Music Player", func() {
	var player MusicPlayer
	BeforeEach(func() {
		player = NewMusicPlayer()
	})
	Describe("Add Song", func() {
		Describe("Add single song", func() {
			It("Should add song to the playlist", func() {
				player.AddSong(Song{
					Singer: "Tulus",
					Title:  "Hati-Hati di Jalan",
				})
				Expect(player.Playlist.Songs).To(HaveLen(1))
				Expect(player.Playlist.Songs[0].Singer).To(Equal("Tulus"))
				Expect(player.Playlist.Songs[0].Title).To(Equal("Hati-Hati di Jalan"))
			})
		})
		Describe("Add multiple songs", func() {
			It("Should add songs to the playlist", func() {
				player.AddSong(Song{
					Singer: "Tulus",
					Title:  "Hati-Hati di Jalan",
				})
				player.AddSong(Song{
					Singer: "NIKI",
					Title:  "Every Summertime",
				})
				Expect(player.Playlist.Songs).To(HaveLen(2))
				Expect(player.Playlist.Songs[0].Singer).To(Equal("Tulus"))
				Expect(player.Playlist.Songs[0].Title).To(Equal("Hati-Hati di Jalan"))
				Expect(player.Playlist.Songs[1].Singer).To(Equal("NIKI"))
				Expect(player.Playlist.Songs[1].Title).To(Equal("Every Summertime"))
			})
		})
	})
	Describe("Repeat", func() {
		It("Should set IsRepeatable to be true", func() {
			player.Playlist.Repeat()
			Expect(player.Playlist.IsRepeatable).To(BeTrue())
		})
	})
	Describe("Play Song", func() {
		Describe("Play song in the empty playlist", func() {
			It("Should return nothing", func() {
				output := player.Play()
				Expect(output).To(Equal(""))
			})
		})
		Describe("Play playlist with one song", func() {
			Describe("Play playlist with non-repeatable condition", func() {
				It("Should play the first song and playlist become empty", func() {
					player.AddSong(Song{
						Singer: "Tulus",
						Title:  "Hati-Hati di Jalan",
					})
					output := player.Play()
					Expect(player.Playlist.Songs).To(HaveLen(0))
					Expect(output).To(Equal("Now playing Tulus - Hati-Hati di Jalan"))
				})
			})
			Describe("Play playlist with repeatable condition", func() {
				It("Should play the first song and played song exists in playlist", func() {
					player.AddSong(Song{
						Singer: "Tulus",
						Title:  "Hati-Hati di Jalan",
					})
					player.Playlist.Repeat()
					output := player.Play()
					Expect(player.Playlist.Songs).To(HaveLen(1))
					Expect(output).To(Equal("Now playing Tulus - Hati-Hati di Jalan"))
				})
			})
		})
		Describe("Play playlist with multiple songs", func() {
			Describe("Play playlist with non-repeatable condition", func() {
				Describe("Play the first song", func() {
					It("Should play the first song and played song not exists in playlist", func() {
						player.AddSong(Song{
							Singer: "Tulus",
							Title:  "Hati-Hati di Jalan",
						})
						player.AddSong(Song{
							Singer: "TREASURE",
							Title:  "DARARI",
						})
						player.AddSong(Song{
							Singer: "NIKI",
							Title:  "Every Summertime",
						})
						player.AddSong(Song{
							Singer: "Pamungkas",
							Title:  "To The Bone",
						})
						output := player.Play()
						Expect(player.Playlist.Songs).To(HaveLen(3))
						Expect(output).To(Equal("Now playing Tulus - Hati-Hati di Jalan"))
					})
				})
				Describe("Play the second song", func() {
					It("Should play the second song and played song not exists in playlist", func() {
						player.AddSong(Song{
							Singer: "Tulus",
							Title:  "Hati-Hati di Jalan",
						})
						player.AddSong(Song{
							Singer: "TREASURE",
							Title:  "DARARI",
						})
						player.AddSong(Song{
							Singer: "NIKI",
							Title:  "Every Summertime",
						})
						player.AddSong(Song{
							Singer: "Pamungkas",
							Title:  "To The Bone",
						})
						player.Play()
						output := player.Play()
						Expect(player.Playlist.Songs).To(HaveLen(2))
						Expect(output).To(Equal("Now playing TREASURE - DARARI"))
					})
				})

			})
			Describe("Play playlist with repeatable condition", func() {
				Describe("Play the first song", func() {
					It("Should play the first song and played song become the last song in the playlist", func() {
						player.AddSong(Song{
							Singer: "Tulus",
							Title:  "Hati-Hati di Jalan",
						})
						player.AddSong(Song{
							Singer: "TREASURE",
							Title:  "DARARI",
						})
						player.AddSong(Song{
							Singer: "NIKI",
							Title:  "Every Summertime",
						})
						player.AddSong(Song{
							Singer: "Pamungkas",
							Title:  "To The Bone",
						})
						player.Playlist.Repeat()
						output := player.Play()
						Expect(player.Playlist.Songs).To(HaveLen(4))
						Expect(output).To(Equal("Now playing Tulus - Hati-Hati di Jalan"))
						Expect(player.Playlist.Songs[3].Singer).To(Equal("Tulus"))
						Expect(player.Playlist.Songs[3].Title).To(Equal("Hati-Hati di Jalan"))
					})
				})
				Describe("Play the second song", func() {
					It("Should play the second song and played song become the last song in the playlist", func() {
						player.AddSong(Song{
							Singer: "Tulus",
							Title:  "Hati-Hati di Jalan",
						})
						player.AddSong(Song{
							Singer: "TREASURE",
							Title:  "DARARI",
						})
						player.AddSong(Song{
							Singer: "NIKI",
							Title:  "Every Summertime",
						})
						player.AddSong(Song{
							Singer: "Pamungkas",
							Title:  "To The Bone",
						})
						player.Playlist.Repeat()
						player.Play()
						output := player.Play()
						Expect(player.Playlist.Songs).To(HaveLen(4))
						Expect(output).To(Equal("Now playing TREASURE - DARARI"))
						Expect(player.Playlist.Songs[3].Singer).To(Equal("TREASURE"))
						Expect(player.Playlist.Songs[3].Title).To(Equal("DARARI"))
					})
				})
			})
		})
	})
})
