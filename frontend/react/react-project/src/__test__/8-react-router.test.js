// NOTE: React Router Dom v6 heavily based on window.location & window.history. This changes the behavior of the router, so we can't interact the history object directly with jest.

import { render, screen, waitFor } from "@testing-library/react"
import { API_URL } from "../api/config"
import App from "../App"

import { MemoryRouter } from "react-router-dom"

import axios from "axios"

jest.mock("axios")
jest.mock("react-router-dom", () => {
  const originalModule = jest.requireActual("react-router-dom")

  return {
    __esModule: true,
    ...originalModule,
    BrowserRouter: ({ children }) => <div>{children}</div>,
  }
})

describe("Router", () => {
  beforeEach(() => {
    axios.get.mockImplementation((url, _) => {
      if (url === `${API_URL}/post/list`) {
        return Promise.resolve({
          status: 200,
          data: {
            message: "success",
            data: [
              {
                id: "cl2ebcto20807s8pzpqqs16g0",
                title: null,
                content: "Test Caption 1111",
                image: "https://example.org/image.png",
                createdAt: "2022-04-25T06:03:31.249Z",
                author: {
                  id: "cl2eb0giz0669s8pzbwdombr3",
                  name: "User Testing 1",
                  image: "https://example.org/image.png",
                },
                liked: false,
                disliked: false,
                likeCount: 11,
                dislikeCount: 22,
              },
              {
                id: "cl2ebcto20807s8pzpqqs16g1",
                title: null,
                content: "Test Caption 2222",
                image: "https://example.org/image.png",
                createdAt: "2022-04-25T06:03:31.249Z",
                author: {
                  id: "cl2eb0giz0669s8pzbwdombr4",
                  name: "User Testing 2",
                  image: "https://example.org/image.png",
                },
                liked: false,
                disliked: false,
                likeCount: 33,
                dislikeCount: 44,
              },
            ],
          },
        })
      } else if (url === `${API_URL}/auth/session`) {
        return Promise.resolve({
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
      } else if (url === `${API_URL}/profile/cl2eb0giz0669s8pzbwdombr4`) {
        return Promise.resolve({
          status: 200,
          data: {
            message: "success",
            data: {
              posts: [
                {
                  id: "cl2ebcto20807s8pzpqqs16g1",
                  title: null,
                  content: "Test Caption 2222",
                  image: "https://example.org/image.png",
                  createdAt: "2022-04-25T06:03:31.249Z",
                  liked: false,
                  disliked: false,
                  likeCount: 33,
                  dislikeCount: 44,
                },
                {
                  id: "cl2ebcto20807s8pzpqqs16g2",
                  title: null,
                  content: "Test Caption 3333",
                  image: "https://example.org/image.png",
                  createdAt: "2022-04-25T06:03:31.249Z",
                  liked: false,
                  disliked: false,
                  likeCount: 55,
                  dislikeCount: 66,
                },
              ],
              profile: {
                id: "cl2eb0giz0669s8pzbwdombr4",
                name: "User Testing 2",
                image: "https://example.org/image.png",
              },
            },
          },
        })
      } else {
        return Promise.resolve({
          status: 404,
          data: {},
        })
      }
    })
  })

  it("should render the index page", async () => {
    render(
      <MemoryRouter initialEntries={["/"]}>
        <App />
      </MemoryRouter>,
    )
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/post/list`, {
      withCredentials: true,
    })
    await waitFor(() => {
      expect(screen.getByText("Test Caption 1111")).toBeInTheDocument()
    })
  })

  it("should render /profile/<UserId> route", async () => {
    render(
      <MemoryRouter initialEntries={["/profile/cl2eb0giz0669s8pzbwdombr4"]}>
        <App />
      </MemoryRouter>,
    )
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/auth/session`, {
      withCredentials: true,
    })
    expect(axios.get).toHaveBeenCalledWith(
      `${API_URL}/profile/cl2eb0giz0669s8pzbwdombr4`,
      { withCredentials: true },
    )

    await waitFor(() => {
      const profile = screen.getAllByText("User Testing 2")
      expect(profile).toHaveLength(3)
    })

    const posts = await screen.findAllByLabelText("Post Card")
    expect(posts).toHaveLength(2)

    expect(posts[0]).toHaveTextContent("Test Caption 2222")
    expect(posts[1]).toHaveTextContent("Test Caption 3333")
  })
})
