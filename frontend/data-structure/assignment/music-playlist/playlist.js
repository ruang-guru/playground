const Song = require('./song')

module.exports = class Playlist {
    constructor(songs, repeatable) {
        this.songs = songs
        this.isRepeatable = repeatable
    }

    isEmpty() {
        return this.songs.length == 0
    }

    repeat() {
        this.isRepeatable = true
    }
}