import { render, screen } from "@testing-library/react"
import Button from "../components/LikeDislikeButton"
import { API_URL } from "../api/config"

import axios from "axios"
import { act } from "react-dom/test-utils"

jest.mock("axios")
describe("Like & Dislike API", () => {
  beforeEach(() => {
    axios.get.mockResolvedValue({
      status: 200,
      data: {
        message: "success",
        data: {
          id: "123",
          userId: "456",
          postId: "789",
          createdAt: "2020-01-01",
          type: "LIKE",
        },
      },
    })
  })
  it("should call api when like", async () => {
    render(
      <Button
        id={123}
        likeCount={0}
        dislikeCount={0}
        isLiked={false}
        isDisliked={true}
      />,
    )
    const likeButton = screen.getByLabelText("Like Button")
    await act(async () => {
      likeButton.click()
    })
    expect(axios.get).toHaveBeenCalledTimes(1)
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/post/123/like`, {
      withCredentials: true,
    })
  })

  it("should call api when dislike", async () => {
    render(
      <Button
        id={123}
        likeCount={0}
        dislikeCount={0}
        isLiked={false}
        isDisliked={false}
      />,
    )
    const dislikeButton = screen.getByLabelText("Dislike Button")
    await act(async () => {
      dislikeButton.click()
    })
    expect(axios.get).toHaveBeenCalledTimes(1)
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/post/123/dislike`, {
      withCredentials: true,
    })
  })

  it("should call api when unlike", async () => {
    render(
      <Button
        id={123}
        likeCount={1}
        dislikeCount={0}
        isLiked={true}
        isDisliked={false}
      />,
    )
    const likeButton = screen.getByLabelText("Like Button")
    await act(async () => {
      likeButton.click()
    })
    expect(axios.get).toHaveBeenCalledTimes(1)
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/post/123/unlike`, {
      withCredentials: true,
    })
  })

  it("should call api when undislike", async () => {
    render(
      <Button
        id={123}
        likeCount={1}
        dislikeCount={0}
        isLiked={false}
        isDisliked={true}
      />,
    )
    const dislikeButton = screen.getByLabelText("Dislike Button")
    await act(async () => {
      dislikeButton.click()
    })
    expect(axios.get).toHaveBeenCalledTimes(1)
    expect(axios.get).toHaveBeenCalledWith(`${API_URL}/post/123/undislike`, {
      withCredentials: true,
    })
  })
})
