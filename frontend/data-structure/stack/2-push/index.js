class Stack {
    constructor() {
        this.data = []
        this.top = -1
    }

    push(elemen) {
        this.top += 1
        return this.data.push(elemen)
    }
}
