import { fireEvent, render, screen} from "@testing-library/react"
import { API_URL } from "../api/config"
import UploadForm from "../components/UploadForm"
import axios from "axios"
import { act } from "react-dom/test-utils"

jest.mock("axios")

describe("Upload Form", () => {
  it("should render the components", () => {
    render(<UploadForm />)
  })

  it("should render the form", () => {
    render(<UploadForm />)
    const form = screen.getByLabelText("Upload Form")
    expect(form).toBeInTheDocument()
  })

  it("should render the post caption input", () => {
    render(<UploadForm />)
    const captionInput = screen.getByLabelText("Caption Input")
    expect(captionInput).toBeInTheDocument()
  })

  it("should render the post image input", () => {
    render(<UploadForm />)
    const fileinput = screen.getByLabelText("Image Input")
    expect(fileinput).toBeInTheDocument()
  })

  it("should accept only some image file format", () => {
    render(<UploadForm />)
    const fileinput = screen.getByLabelText("Image Input")
    expect(fileinput).toHaveAttribute("type", "file")
    expect(fileinput).toHaveAttribute(
      "accept",
      "image/png, image/jpg, image/gif",
    )
  })

  it("should render the button", () => {
    render(<UploadForm />)
    const button = screen.getByLabelText("Submit Button")
    expect(button).toBeInTheDocument()
  })

  it("should change the caption when caption input changed", () => {
    render(<UploadForm />)
    const captionInput = screen.getByLabelText("Caption Input")
    fireEvent.change(captionInput, { target: { value: "Hello World" } })
    expect(captionInput).toHaveValue("Hello World")
  })

  it("should change the image when image input changed", () => {
    const mockFile = new File(["(⌐□_□)"], "gambar_dari_folder_tugas.png", {
      type: "image/png",
    })
    render(<UploadForm />)
    const fileinput = screen.getByLabelText("Image Input")
    fireEvent.change(fileinput, { target: { files: [mockFile] } })
    expect(fileinput.files.length).toBe(1)
    expect(fileinput.files[0].name).toBe("gambar_dari_folder_tugas.png")
  })

  it("should post to the api using form data when the form is submitted", async () => {
    const mockFile = new File(["(⌐□_□)"], "gambar_dari_folder_tugas.png", {
      type: "image/png",
    })
    const mockCaption = "ini gambar dari folder 'tugas'"
    const mockFormData = new FormData()

    mockFormData.append("content", mockCaption)
    mockFormData.append("image", mockFile)

    axios.post.mockResolvedValue({
      message: "success",
      data: {
        id: "cas2df35r34",
        content: mockCaption,
        image: "https://example.com/chucknorris.png",
        authorId: "c83i10ew3a1s",
        createdAt: "2020-01-01T00:00:00.000Z",
      },
    })

    render(<UploadForm onSubmit={()=>{}} />)
    const fileinput = screen.getByLabelText("Image Input")
    const captioninput = screen.getByLabelText("Caption Input")
    const button = screen.getByLabelText("Submit Button")

    fireEvent.change(fileinput, { target: { files: [mockFile] } })
    fireEvent.change(captioninput, { target: { value: mockCaption } })

    await act(async () => {
      button.click()
    })

    expect(axios.post).toHaveBeenCalledWith(
      `${API_URL}/post/create`,
      mockFormData,
      {
        withCredentials: true,
      },
    )
  })
})
