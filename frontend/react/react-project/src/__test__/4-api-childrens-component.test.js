import { render, screen, waitFor, fireEvent } from "@testing-library/react"
import { act } from "react-dom/test-utils"
import { API_URL } from "../api/config"
import App from "../App"
import axios from "axios"
jest.mock("axios")

describe("App Childrens Components", () => {
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
      } else {
        return Promise.resolve({
          status: 404,
          data: {},
        })
      }
    })
  })

  it("should render navbar", async () => {
    render(<App />)
    await waitFor(async () => {
      const navbar = await screen.findByLabelText("Navbar")
      expect(navbar).toBeInTheDocument()
    })
  })

  it("should request session and render navbar(profile) with correct data", async () => {
    render(<App />)
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/auth/session`, {
      withCredentials: true,
    })

    const profile = await screen.findByLabelText("Profile")
    expect(profile).toHaveTextContent("John Doe")
  })
  it("should request post list & render the Post Card components", async () => {
    render(<App />)
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/post/list`, {
      withCredentials: true,
    })
    const cards = await screen.findAllByLabelText("Post Card")
    expect(cards).toHaveLength(2)
  })

  it("should render the Post Card components with correct data", async () => {
    render(<App />)
    const cards = await screen.findAllByLabelText("Post Card")
    expect(cards).toHaveLength(2)
    expect(cards[0]).toHaveTextContent("Test Caption 1111")
    expect(cards[0]).toHaveTextContent("User Testing 1")
    expect(cards[1]).toHaveTextContent("Test Caption 2222")
    expect(cards[1]).toHaveTextContent("User Testing 2")
  })

  it("should render the upload form", async () => {
    render(<App />)
    await waitFor(async () => {
      const uploadForm = screen.getByLabelText("Upload Form")
      expect(uploadForm).toBeInTheDocument()
    })
  })

  it("should lift the state up after submit (optimistic update)", async () => {
    axios.post.mockResolvedValue({
      status: 200,
      data: {
        id: "cas2df35r34",
        content: "Test New Caption 9999",
        image: "https://example.com/image.png",
        authorId: "c83i10ew3a1s",
        createdAt: "2020-01-01T00:00:00.000Z",
      },
    })
    const mockFile = new File(["(⌐□_□)"], "gambar_dari_folder_tugas.png", {
      type: "image/png",
    })
    const mockCaption = "ini gambar dari folder 'tugas'"
    render(<App />)
    const fileinput = screen.getByLabelText("Image Input")
    const captioninput = screen.getByLabelText("Caption Input")
    const button = screen.getByLabelText("Submit Button")
    fireEvent.change(fileinput, { target: { files: [mockFile] } })
    fireEvent.change(captioninput, { target: { value: mockCaption } })

    await act(async () => {
      button.click()
    })

    const post = await screen.findByText("Test New Caption 9999")
    expect(post).toBeInTheDocument()
  })
})
