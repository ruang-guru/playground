import React, { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";

const PokemonDetail = () => {
  /**
   * 7. Sesuai param yang telah didefinisikan pada Route (:name),
   * maka apabila bisa ditangkap sebagai param.name
   */
  const param = useParams(); 
  const [detail, setDetail] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  const loadPokemonDetail = async () => {
    setIsLoading(true);
    try {
      /**
       * 8. param.name lalu digunakan untuk fetching data dari API
       * endpoint detail sebagai URL parameter
       */
      const { data } = await axios.get(
        "https://pokeapi.co/api/v2/pokemon/" + param.name
      );
      setDetail(data);
    } catch (error) {
      console.log(error);
    }
    setIsLoading(false);
  };

  useEffect(() => {
    loadPokemonDetail();
  }, []);

  return (
    <div className="poke-detail">
      {isLoading ? (
        <h1>Loading...</h1>
      ) : detail ? (
        <>
          <h1>{param.name}</h1>

          <img
            src={detail && detail.sprites.front_default}
            width={200}
            height={200}
          />

          <div>
            <h2>types</h2>
            <div className="poke-types">
              {detail &&
                detail.types.map((type, index) => (
                  <span className="poke-card" key={index}>
                    {type.type.name}
                  </span>
                ))}
            </div>
          </div>

          <div>
            <h2>moves</h2>
            <div className="poke-moves">
              {detail &&
                detail.moves.map((move, index) => (
                  <span className="poke-card" key={index}>
                    {move.move.name}
                  </span>
                ))}
            </div>
          </div>
        </>
      ) : (
        <h1>No Pokemon Found</h1>
      )}
    </div>
  );
};

export default PokemonDetail;
