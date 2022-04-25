class Stack {
    constructor() {
        this.data = []
        this.top = -1
    }

    push(elemen) {
        this.top += 1
        return this.data.push(elemen)
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

let word = "hello world"

var stack = new Stack()
for (let i = 0; i < word.length; i++) {
  stack.push(word.charAt(i))
}

let reversed = ""
while (!stack.isEmpty()) {
  let ch = stack.pop()
  reversed += ch
}

console.log(reversed)
