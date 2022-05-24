import { render } from "@testing-library/react"
import App from "../App"

class LocalStorageMock {
  constructor() {
    this.store = {}
  }

  clear() {
    this.store = {}
  }

  getItem(key) {
    return this.store[key] || null
  }

  setItem(key, value) {
    this.store[key] = String(value)
  }

  removeItem(key) {
    delete this.store[key]
  }
}

global.localStorage = new LocalStorageMock()

describe("Persistent data test", () => {

    it("should save items to local storage", () => {
        render(<App />)
        
        expect(localStorage.getItem("items")).not.toBeNull()
    })

    it("should save cart to local storage", () => {
        render(<App />)
        
        expect(localStorage.getItem("cart")).not.toBeNull()
    })
})
