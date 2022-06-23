import { render, screen } from "@testing-library/react"
import Button from "../components/LikeDislikeButton"
import { act } from "react-dom/test-utils"
import axios from "axios"

jest.mock("axios")
describe("Like & Dislike Button Component", () => {
  axios.get.mockResolvedValue({ status: 200, data: {} })
  it("should render the like and dislike buttons", () => {
    render(
      <Button
        likeCount={0}
        dislikeCount={0}
        isLiked={false}
        isDisliked={true}
      />,
    )
    const likeButton = screen.getByLabelText("Like Button")
    expect(likeButton).toBeInTheDocument()
    const dislikeButton = screen.getByLabelText("Dislike Button")
    expect(dislikeButton).toBeInTheDocument()
  })

  it("should render the like & dislike count", () => {
    render(<Button likeCount={456} dislikeCount={678} />)
    const likeCount = screen.getByLabelText("Like Count")
    expect(likeCount).toBeInTheDocument()
    expect(likeCount).toHaveTextContent("456")
    const dislikeCount = screen.getByLabelText("Dislike Count")
    expect(dislikeCount).toBeInTheDocument()
    expect(dislikeCount).toHaveTextContent("678")
  })

  it("should change the like count when button like clicked", async () => {
    render(
      <Button
        likeCount={476}
        dislikeCount={0}
        isLiked={false}
        isDisliked={true}
      />,
    )
    const likeButton = screen.getByLabelText("Like Button")
    await act(async () => {
      likeButton.click()
    })
    const likeCount = screen.getByLabelText("Like Count")
    expect(likeCount).toHaveTextContent("477")
  })

  it("should change the dislike count when button dislike clicked", async () => {
    render(
      <Button
        likeCount={0}
        dislikeCount={289}
        isLiked={false}
        isDisliked={false}
      />,
    )
    axios.get.mockResolvedValue({ status: 200, data: {} })
    const dislikeButton = screen.getByLabelText("Dislike Button")
    await act(async () => {
      dislikeButton.click()
    })
    const dislikeCount = screen.getByLabelText("Dislike Count")
    expect(dislikeCount).toHaveTextContent("290")
  })

  it("should change the like count when already liked then button like(unlike) clicked", async () => {
    render(
      <Button
        likeCount={476}
        dislikeCount={0}
        isLiked={true}
        isDisliked={false}
      />,
    )
    const likeButton = screen.getByLabelText("Like Button")
    await act(async () => {
      likeButton.click()
    })
    const likeCount = screen.getByLabelText("Like Count")
    expect(likeCount).toHaveTextContent("475")
  })
  it("should change the like count when already disliked then button like clicked", async () => {
    render(
      <Button
        likeCount={476}
        dislikeCount={256}
        isLiked={false}
        isDisliked={true}
      />,
    )
    const likeButton = screen.getByLabelText("Like Button")
    await act(async () => {
      likeButton.click()
    })
    const likeCount = screen.getByLabelText("Like Count")
    expect(likeCount).toHaveTextContent("477")
    const dislikeCount = screen.getByLabelText("Dislike Count")
    expect(dislikeCount).toHaveTextContent("255")
  })
  it("should change the dislike count when already disliked then button dislike(undislike) clicked", async () => {
    render(
      <Button
        likeCount={0}
        dislikeCount={289}
        isLiked={false}
        isDisliked={true}
      />,
    )
    const dislikeButton = screen.getByLabelText("Dislike Button")
    await act(async () => {
      dislikeButton.click()
    })
    const dislikeCount = screen.getByLabelText("Dislike Count")
    expect(dislikeCount).toHaveTextContent("288")
  })
  it("should change the dislike count when already liked then button dislike clicked", async () => {
    render(
      <Button
        likeCount={278}
        dislikeCount={289}
        isLiked={true}
        isDisliked={false}
      />,
    )
    const dislikeButton = screen.getByLabelText("Dislike Button")
    await act(async () => {
      dislikeButton.click()
    })
    const dislikeCount = screen.getByLabelText("Dislike Count")
    expect(dislikeCount).toHaveTextContent("290")
    const likeCount = screen.getByLabelText("Like Count")
    expect(likeCount).toHaveTextContent("277")
  })
})
