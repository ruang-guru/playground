import { render, screen, fireEvent, act } from '@testing-library/react';
import "@testing-library/jest-dom";
import App, {randomEmoji} from '../App';
import React from "react"

describe("Change emoji with button", () => {

  it("should render the Acak emoji Button", () => {
    render(<App />)
    const emojiButton = screen.getByText("Acak emoji")
    expect(emojiButton).toBeInTheDocument()
  })

  it("should get emoji after random emoji", async () => {
    render(
      <App />,
    )
    let dataEmoji = randomEmoji()
    expect(dataEmoji).toBe(dataEmoji);
  })

  it("should change the emoji after Acak emoji button clicked", async () => {
    render(
      <App />,
    )
    const emojiButton = screen.getByText("Acak emoji")
    const emojiDisplay = screen.getByTestId("emoji")

    await act(async () => {
      emojiButton.click()
    })

    expect(emojiDisplay.textContent).not.toBe("");
  })
})
