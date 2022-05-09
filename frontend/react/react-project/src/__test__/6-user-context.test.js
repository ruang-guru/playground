import { useContext } from "react"
import { SessionContext, SessionProvider } from "../context/SessionContext"
import React from "react"
import { render } from "@testing-library/react"

const ContextSetter = () => {
  const { setUser, setIsLoggedIn } = useContext(SessionContext)
  return (
    <>
      <button
        aria-label="Set User"
        onClick={() =>
          setUser({
            id: "1",
            name: "John Doe",
            image: "https://example.org/abc.png",
          })
        }
      >
        Set User
      </button>
      <button aria-label="Set IsLoggedIn" onClick={() => setIsLoggedIn(true)}>
        Set IsLoggedIn
      </button>
    </>
  )
}

const ContextGetter = () => {
  const { user, isLoggedIn } = useContext(SessionContext)
  return (
    <>
      <p>{`name: ${user.name}`}</p>
      <p>{`image: ${user.image}`}</p>
      <p>{`loggedIn: ${isLoggedIn ? "Yes" : "No"}`}</p>
    </>
  )
}

describe("User Context & Provider", () => {
  it("should set user on ContextSetter and render on ContextGetter", () => {
    render(
      <SessionProvider>
        <ContextSetter />
        <ContextGetter />
      </SessionProvider>,
    )
  })

  it("should set isLoggedIn on ContextSetter and render on ContextGetter", () => {
    render(
      <SessionProvider>
        <ContextSetter />
        <ContextGetter />
      </SessionProvider>,
    )
  })
})
