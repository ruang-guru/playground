const { expect, describe, it } = require('@jest/globals');
const Song = require('./song')
const Playlist = require('./playlist')
const MusicPlayer = require('./music-player')

beforeEach(() => {
    let playlist = new Playlist([], false)
    const player = new MusicPlayer(playlist)
})

describe("Music Player", () => {
    describe("when add single song", () => {
        it("should add song to the playlist", () => {
            let playlist = new Playlist([], false)
            const player = new MusicPlayer(playlist)
            let song = new Song("Tulus", "Hati-Hati di Jalan")
            player.addSong(song)

            expect(player.playlist.songs.length).toEqual(1)
            expect(player.playlist.songs[0].singer).toEqual("Tulus")
            expect(player.playlist.songs[0].title).toEqual("Hati-Hati di Jalan")
        })
    })

    describe("when add multiple songs", () => {
        it("should add songs to the playlist", () => {
            let playlist = new Playlist([], false)
            const player = new MusicPlayer(playlist)
            player.addSong(new Song("Tulus", "Hati-Hati di Jalan"))
            player.addSong(new Song("NIKI", "Every Summertime"))

            expect(player.playlist.songs.length).toEqual(2)
            expect(player.playlist.songs[0].singer).toEqual("Tulus")
            expect(player.playlist.songs[0].title).toEqual("Hati-Hati di Jalan")
            expect(player.playlist.songs[1].singer).toEqual("NIKI")
            expect(player.playlist.songs[1].title).toEqual("Every Summertime")
        })
    })

    describe("when repeat playlist", () => {
        it("should set isRepeatable to be true", () => {
            let playlist = new Playlist([], false)
            const player = new MusicPlayer(playlist)

            player.playlist.repeat()
            expect(player.playlist.isRepeatable).toEqual(true)
        })
    })

    describe("when play song in empty playlist", () => {
        it("should return empty string", () => {
            let playlist = new Playlist([], false)
            const player = new MusicPlayer(playlist)

            let output = player.play()
            
            expect(output).toEqual("")
        })
    })
    describe("when play song with one song in non-repeatable condition", () => {
        it("should play the first song and playlist become empty", () => {
            let playlist = new Playlist([], false)
            const player = new MusicPlayer(playlist)
            player.addSong(new Song("Tulus", "Hati-Hati di Jalan"))

            let output = player.play()
            
            expect(player.playlist.songs.length).toEqual(0)
            expect(output).toEqual("Now Playing Tulus - Hati-Hati di Jalan")
        })
    })
    describe("when play song with one song in repeatable condition", () => {
        it("should play the first song and played song exists in playlist", () => {
            let playlist = new Playlist([], false)
            const player = new MusicPlayer(playlist)
            player.addSong(new Song("Tulus", "Hati-Hati di Jalan"))
            player.playlist.repeat()

            let output = player.play()
            
            expect(player.playlist.songs.length).toEqual(1)
            expect(output).toEqual("Now Playing Tulus - Hati-Hati di Jalan")
            expect(player.playlist.songs[0].singer).toEqual("Tulus")
            expect(player.playlist.songs[0].title).toEqual("Hati-Hati di Jalan")
        })
    })

    describe("when play song with multiple song in non-repeatable condition", () => {
        describe("play the first song", () => {
            it("should play the first song and played song not exists in playlist", () => {
                let playlist = new Playlist([], false)
                const player = new MusicPlayer(playlist)
                player.addSong(new Song("Tulus", "Hati-Hati di Jalan"))
                player.addSong(new Song("TREASURE", "DARARI"))
                player.addSong(new Song("NIKI", "Every Summertime"))
                player.addSong(new Song("Pamungkas", "To The Bone"))
    
                let output = player.play()
                
                expect(player.playlist.songs.length).toEqual(3)
                expect(output).toEqual("Now Playing Tulus - Hati-Hati di Jalan")
            })
        })
        describe("play the second song", () => {
            it("should play the first song and played song not exists in playlist", () => {
                let playlist = new Playlist([], false)
                const player = new MusicPlayer(playlist)
                player.addSong(new Song("Tulus", "Hati-Hati di Jalan"))
                player.addSong(new Song("TREASURE", "DARARI"))
                player.addSong(new Song("NIKI", "Every Summertime"))
                player.addSong(new Song("Pamungkas", "To The Bone"))
    
                player.play()
                let output = player.play()
                
                expect(player.playlist.songs.length).toEqual(2)
                expect(output).toEqual("Now Playing TREASURE - DARARI")
            })
        })
        
    })
    describe("when play song with multiple song in repeatable condition", () => {
        describe("play the first song", () => {
            it("should play the first song and played song become the last song in the playlist", () => {
                let playlist = new Playlist([], false)
                const player = new MusicPlayer(playlist)
                player.addSong(new Song("Tulus", "Hati-Hati di Jalan"))
                player.addSong(new Song("TREASURE", "DARARI"))
                player.addSong(new Song("NIKI", "Every Summertime"))
                player.addSong(new Song("Pamungkas", "To The Bone"))
    
                player.playlist.repeat()
                let output = player.play()
                
                expect(player.playlist.songs.length).toEqual(4)
                expect(output).toEqual("Now Playing Tulus - Hati-Hati di Jalan")
                expect(player.playlist.songs[3].singer).toEqual("Tulus")
                expect(player.playlist.songs[3].title).toEqual("Hati-Hati di Jalan")
            })
        })
        describe("play the second song", () => {
            it("play the second song and played song become the last song in the playlist", () => {
                let playlist = new Playlist([], false)
                const player = new MusicPlayer(playlist)
                player.addSong(new Song("Tulus", "Hati-Hati di Jalan"))
                player.addSong(new Song("TREASURE", "DARARI"))
                player.addSong(new Song("NIKI", "Every Summertime"))
                player.addSong(new Song("Pamungkas", "To The Bone"))
                player.playlist.repeat()
    
                player.play()
                let output = player.play()
                
                expect(player.playlist.songs.length).toEqual(4)
                expect(output).toEqual("Now Playing TREASURE - DARARI")
                expect(player.playlist.songs[3].singer).toEqual("TREASURE")
                expect(player.playlist.songs[3].title).toEqual("DARARI")
            })
        })
        
    })
})
