import { render, screen } from "@testing-library/react"
import CardImage from "../components/PostCard"
import {BrowserRouter} from "react-router-dom"



describe("CardImage Component", () => {
    it ("should render the components", () => {
        render(
          <BrowserRouter>
            <CardImage />
          </BrowserRouter>,
        )
        const postCard = screen.getByLabelText("Post Card")
        expect(postCard).toBeInTheDocument()
    })

    it ("should render the image", () => {
        const random = 'https://example.com/' + (Math.random() + 1).toString(36).substring(2)
        render(
          <BrowserRouter>
            <CardImage image={random} />
          </BrowserRouter>,
        )
        const image = screen.getByLabelText("Post Image")
        expect(image).toHaveAttribute('src', random)
    })

    it ("should render the caption", () => {
        const random = (Math.random() + 1).toString(36).substring(2);
        render(
          <BrowserRouter>
            <CardImage caption={random} />
          </BrowserRouter>,
        )
        const caption = screen.getByLabelText("Post Caption")
        expect(caption).toHaveTextContent(random)
    })

    it ("should render the username", () => {
        const random = (Math.random() + 1).toString(36).substring(2)
        render(
          <BrowserRouter>
            <CardImage username={random} />
          </BrowserRouter>,
        )
        const username = screen.getByLabelText("Post User Name")
        expect(username).toHaveTextContent(random)
    })

    it ("should have button like and dislike inside", () => {
        render(
          <BrowserRouter>
            <CardImage />
          </BrowserRouter>,
        )
        const likeButton = screen.getByLabelText("Like Button")
        const dislikeButton = screen.getByLabelText("Dislike Button")
        expect(likeButton).toBeInTheDocument()
        expect(dislikeButton).toBeInTheDocument()
    })

    it("should have count for like and dislike inside", () => {
      render(
        <BrowserRouter>
          <CardImage />
        </BrowserRouter>,
      )
      const likeButton = screen.getByLabelText("Like Count")
      const dislikeButton = screen.getByLabelText("Dislike Count")
      expect(likeButton).toBeInTheDocument()
      expect(dislikeButton).toBeInTheDocument()
    })
})