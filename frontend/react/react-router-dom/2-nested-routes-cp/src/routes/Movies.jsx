import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import moviesThumbnail from "../assets/movies.jpeg";
import BackButton from "../components/BackButton";

const Movies = () => {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(false);
  const controller = new AbortController();

  const loadMovies = async () => {
    setLoading(true);
    try {
      const { data } = await axios.get("https://swapi.dev/api/films", {
        signal: controller.signal,
      });
      // TODO: answer here
    } catch (error) {
      console.log(error);
    }
    setLoading(false);
  };

  const getID = (url) => {
    const parsed = url?.split("/");
    return parsed[parsed.length - 2];
  };

  useEffect(() => {
    loadMovies();
    return () => {
      controller.abort();
    };
  }, []);

  return (
    <div className="page">
      <header>
        <BackButton />
        <img
          src={moviesThumbnail}
          width={40}
          height={40}
          alt="Movies Thumbnail"
        />
        <h1>Movies</h1>
      </header>

      {!loading ? (
        <div className="grid">
          {movies.length &&
            movies.map((movie, index) => (
              // TODO: answer here
            ))}
        </div>
      ) : (
        <h2>Loading...</h2>
      )}
    </div>
  );
};

export default Movies;
