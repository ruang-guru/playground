/**
 * @jest-environment jsdom
 */
import React from "react"
import { render, screen } from "@testing-library/react"
import Navbar from "../components/Navbar"
import axios from "axios"
import { act } from "react-dom/test-utils"
import { API_URL, BASE_URL } from "../api/config"

jest.mock("axios")

describe("Login", () => {
  it("should render login button", () => {
    render(<Navbar />)
    const button = screen.getByLabelText(/Login/i)
    expect(button).toBeInTheDocument()
  })

  it("should redirect to oauth page when login button is clicked", async () => {
    delete window.location
    window.location = { assign: jest.fn() }
    render(<Navbar />)
    const button = screen.getByLabelText(/Login/i)
    await act(async () => {
      button.click()
    })
    expect(window.location.assign).toHaveBeenCalledTimes(1)
    expect(window.location.assign).toHaveBeenCalledWith(
      expect.stringContaining(`${BASE_URL}/auth?redirect=`),
    )
  })

  it("should render profile name if logged in", async () => {
    axios.get.mockResolvedValue({
      status: 200,
      data: {
        message: "success",
        user: {
          id: "123",
          name: "John Doe",
          image: "https://example.com/image.png",
          email: "john doe",
        },
      },
    })
    render(<Navbar />)
    expect(axios.get).toHaveBeenCalledTimes(1)
    expect(axios.get).toHaveBeenCalledWith(
      `${API_URL}/auth/session`,
      { withCredentials: true },
    )

    const username = await screen.findByLabelText("Profile")
    expect(username).toBeInTheDocument()
    expect(username).toHaveTextContent("John Doe")
  })
})
