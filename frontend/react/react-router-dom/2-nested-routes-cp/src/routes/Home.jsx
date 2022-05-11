import React from "react";
import { Link } from "react-router-dom";
import people from "../assets/people.jpeg";
import planets from "../assets/planets.jpeg";
import movies from "../assets/movies.jpeg";

const Home = () => {
  return (
    <main className="main-menu">
      <Link to="/star-wars/people">
        <div>
          <h1>People</h1>
          <img src={people} width={240} alt="People Thumbnail" />
        </div>
      </Link>

      <Link to="/star-wars/planets">
        <div>
          <h1>Planets</h1>
          <img src={planets} width={240} alt="Planets Thumbnail" />
        </div>
      </Link>

      <Link to="/star-wars/movies">
        <div>
          <h1>Movies</h1>
          <img src={movies} width={240} alt="Movies Thumbnail" />
        </div>
      </Link>
    </main>
  );
};

export default Home;
