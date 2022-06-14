import * as React from "react";
import { render, screen, waitFor } from "@testing-library/react";
import { MemoryRouter, Routes, Route } from "react-router-dom";
import axios from "axios";
import App from "../App";

const controller = new AbortController()

jest.mock("axios")

describe("routing", () => {
  beforeEach(() => {
    axios.get.mockImplementation((url, _) => {
      if (url == `https://swapi.dev/api/films`) {
        return Promise.resolve({
          status: 200,
          data: {
            results: [
              {
                "title": "A New Hope",
                "url": "https://swapi.dev/api/films/1/"
              },
              {
                "title": "The Empire Strikes Back",
                "url": "https://swapi.dev/api/films/2/"
              }
            ]
          }
        })
      } else if (url === `https://swapi.dev/api/films/2`) {
        return Promise.resolve({
          status: 200,
          data: {
            "opening_crawl": "It is a dark time for the\r\nRebellion. Although the Death\r\nStar has been destroyed,\r\nImperial troops have driven the\r\nRebel forces from their hidden\r\nbase and pursued them across\r\nthe galaxy.\r\n\r\nEvading the dreaded Imperial\r\nStarfleet, a group of freedom\r\nfighters led by Luke Skywalker\r\nhas established a new secret\r\nbase on the remote ice world\r\nof Hoth.\r\n\r\nThe evil lord Darth Vader,\r\nobsessed with finding young\r\nSkywalker, has dispatched\r\nthousands of remote probes into\r\nthe far reaches of space....",
            "director": "Irvin Kershner",
            "producer": "Gary Kurtz, Rick McCallum",
            "release_date": "1980-05-17"
          }
        })
      } else if (url === `https://swapi.dev/api/planets`) {
        return Promise.resolve({
          status: 200,
          data: {
            results: [{
              "name": "Tatooine",
              "url": "https://swapi.dev/api/planets/1/"
            }]
          }
        })
      } else if (url === `https://swapi.dev/api/planets/1`) {
        return Promise.resolve({
          status: 200,
          data: {
            "rotation_period": "23",
            "orbital_period": "304",
            "terrain": "desert",
            "population": "200000",
            "diameter": "10465",
            "climate": "arid"
          }
        })
      }
    })
  })


  it("renders index route correctly", () => {
    render(
      <MemoryRouter initialEntries={["/"]}>
        <App />
      </MemoryRouter>
    );

    expect(screen.getByText(/People/i)).toBeInTheDocument();
    expect(screen.getByText(/Planets/i)).toBeInTheDocument();
    expect(screen.getByText(/Movies/i)).toBeInTheDocument();
  });

  it("renders /star-wars/planets route correctly", async () => {
    render(
      <MemoryRouter initialEntries={["/star-wars/planets"]}>
        <App />
      </MemoryRouter>
    )

    expect(axios.get).toHaveBeenCalledWith(`https://swapi.dev/api/planets`, {
      signal: controller.signal,
    })

    await waitFor(() => {
      expect(screen.getByText("Tatooine")).toBeInTheDocument()
    })
  });

  it("renders /star-wars/planets/:id route correctly", async () => {
    render(
      <MemoryRouter initialEntries={["/star-wars/planets/1"]}>
        <App />
      </MemoryRouter>
    )

    expect(axios.get).toHaveBeenCalledWith(`https://swapi.dev/api/planets/1`, {
      signal: controller.signal,
    })

    await waitFor(() => {
      expect(screen.getByText("23")).toBeInTheDocument()
      expect(screen.getByText("304")).toBeInTheDocument()
      expect(screen.getByText("desert")).toBeInTheDocument()
      expect(screen.getByText("200000")).toBeInTheDocument()
      expect(screen.getByText("10465")).toBeInTheDocument()
      expect(screen.getByText("arid")).toBeInTheDocument()
    })
  });

  it("renders /star-wars/movies route correctly", async () => {
    render(
      <MemoryRouter initialEntries={["/star-wars/movies"]}>
        <App />
      </MemoryRouter>
    )

    expect(axios.get).toHaveBeenCalledWith(`https://swapi.dev/api/films`, {
      signal: controller.signal,
    })

    await waitFor(() => {
      expect(screen.getByText("A New Hope")).toBeInTheDocument()
    })
  });

  it("renders /star-wars/movies/:id route correctly", async () => {
    render(
      <MemoryRouter initialEntries={["/star-wars/movies/2"]}>
        <App />
      </MemoryRouter>
    )

    expect(axios.get).toHaveBeenCalledWith(`https://swapi.dev/api/films/2`, {
      signal: controller.signal,
    })

    await waitFor(() => {
      expect(screen.getByText("1980-05-17")).toBeInTheDocument()
      expect(screen.getByText("Gary Kurtz, Rick McCallum")).toBeInTheDocument()
      expect(screen.getByText("Irvin Kershner")).toBeInTheDocument()
      expect(screen.getByText("It is a dark time for the Rebellion. Although the Death Star has been destroyed, Imperial troops have driven the Rebel forces from their hidden base and pursued them across the galaxy. Evading the dreaded Imperial Starfleet, a group of freedom fighters led by Luke Skywalker has established a new secret base on the remote ice world of Hoth. The evil lord Darth Vader, obsessed with finding young Skywalker, has dispatched thousands of remote probes into the far reaches of space....")).toBeInTheDocument()
    })
  });
});
