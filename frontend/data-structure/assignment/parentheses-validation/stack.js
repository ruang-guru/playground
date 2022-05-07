module.exports = class Stack {
    constructor() {
        this.data = []
        this.top = -1
    }

    push(elemen) {
        if (this.data.length == this.size) {
            throw "stack overflow"
        } else {
            this.top += 1
            return this.data.push(elemen)
        }
    }

    pop() {
        if (this.top == -1) {
            throw "stack underflow"
        } else {
            let poppedValue = this.data.pop()
            this.top -= 1
            return poppedValue
        }
    }

    peek() {
        if (this.top == -1) {
            throw "stack is empty"
        } else {
            return this.data[this.top]
        }
    }
}
