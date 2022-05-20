import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

const Pokemons = () => {
  const [pokemons, setPokemons] = useState([]);

  const loadPokemons = async () => {
    try {
      const { data } = await axios.get(
        "https://pokeapi.co/api/v2/pokemon?limit=90&offset=151"
      );
      setPokemons(data.results);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    loadPokemons();
  }, []);

  return (
    <div>
      <h1>Pokemons</h1>
      <div className="poke-container">
        {pokemons.map((pokemon, index) => (
          <Link key={index} to={`/pokemons/${pokemon.name}`}>
            <div className="poke-card">
              <p>{pokemon.name}</p>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default Pokemons;
