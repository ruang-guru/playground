import { render, screen, waitFor } from "@testing-library/react"
import Navbar from "../components/Navbar"
import axios from "axios"
import React from "react"

jest.mock("axios")
describe("Navbar Component", () => {
  axios.get.mockResolvedValue({ status: 200, data: {} })

  it("should render the components", () => {
    render(<Navbar />)
    const navbar = screen.getByLabelText("Navbar")
    expect(navbar).toBeInTheDocument()
  })

  it("should render the logo", async() => {
    render(<Navbar />)
    await waitFor(() => {
      const logo = screen.getByLabelText("App Logo")
      expect(logo).toBeInTheDocument()
    })
  })

  it("should render the title", async() => {
    render(<Navbar />)
    await waitFor(() => {
        const title = screen.getByLabelText("App Title")
        expect(title).toBeInTheDocument()
    })
  })

  it("should href link to / in the title", async() => {
      render(<Navbar />)
      await waitFor(() => {
        const link = screen.getByLabelText("App Title")
        expect(link).toHaveAttribute("href", "/")
      })
  })
})
