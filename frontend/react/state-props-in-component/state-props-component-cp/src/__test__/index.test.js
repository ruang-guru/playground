import {render, screen, act} from "@testing-library/react"
import { Item } from "../App"
import App from "../App"

global.alert = jest.fn()

describe("State and Props Component", () => {
  it("should add value when '+' button is clicked", async () => {
    render(
      <App />
    )

    const addButton = screen.getByLabelText('add-button-1')
    const numberCounter = screen.getByLabelText("item-1")

    await act(async () => {
      addButton.click()
    })

    expect(numberCounter.value).toBe("1")
  })

  it("should decrease the value when the '-' button is clicked", async () => {
    render(
      <App />
    )

    const addButton = screen.getByLabelText('add-button-1')
    const minusButton = screen.getByLabelText('minus-button-1')
    const numberCounter = screen.getByLabelText("item-1")

    await act(async () => {
      addButton.click()
    })

    await act(async () => {
      minusButton.click()
    })

    expect(numberCounter.value).toBe("0")
  })

  it("should show alert if the value added more than 10", async()=> {
    render(
      <App />
    )

    const addButton = screen.getByLabelText('add-button-1')
    const numberCounter = screen.getByLabelText("item-1")

    for (let i = 0; i < 11; i++) {
      await act(async () => {
        addButton.click()
      })
    }

    expect(numberCounter.value).toBe("10")
    expect(global.alert).toHaveBeenCalledTimes(1)
  })



  it("should count the total value", async () => {
    let totalItem = [0,0,0]
    const cases = [['add-button-1', 'item-1', 8], ['add-button-2', 'item-2', 7], ['add-button-3', 'item-3', 3], ['minus-button-1', 'item-1', 4], ['minus-button-2', 'item-2', 2], ['minus-button-3', 'item-3', 1]]

    render(<App/>)
    const totalCounter = screen.getByLabelText("total-item")

    for (let i = 0; i < cases.length; i++) {
      const [buttonLabel, countLabel, expected] = cases[i]
      const button = screen.getByLabelText(buttonLabel)
      const id = ~~countLabel.split('-')[1] - 1

      for (let i = 0; i < expected; i++) {
        await act(async () => {
          button.click()
        })
      }

      const buttonType = buttonLabel.split("-")[0]
      if (buttonType === 'add') {
        totalItem[id] += expected
      } else {
        totalItem[id] -= expected
      }
      const numberCounter = screen.getByLabelText(countLabel)
      
      expect(numberCounter.value).toBe(totalItem[id].toString())
      expect(totalCounter).toHaveTextContent(totalItem.reduce((a,b) => a+b).toString())
    }
  })
})