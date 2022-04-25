class Stack {
    constructor() {
        this.data = []
        this.top = -1
    }

    isEmpty() {
        return this.top == -1
    }

    pop() {
        if (this.isEmpty()) {
            throw "stack underflow"
        } else {
            let poppedValue = this.data.pop()
            this.top -= 1
            return poppedValue
        }
    }
}
