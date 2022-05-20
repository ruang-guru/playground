import React, { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import BackButton from "../components/BackButton";

const PlanetDetail = () => {
  const [detail, setDetail] = useState(null);
  const [loading, setLoading] = useState(false);
  // TODO: answer here
  const controller = new AbortController();

  const loadDetail = async () => {
    setLoading(true);
    try {
      const url = /* beginanswer */ "https://swapi.dev/api/planets/" + id; /* endanswer "" */
      const { data } = await axios.get(url, {
        signal: controller.signal,
      });
      // TODO: answer here
    } catch (error) {
      console.log(error);
    }
    setLoading(false);
  };

  useEffect(() => {
    // TODO: answer here
    return () => {
      controller.abort();
    };
  }, []);

  return (
    <div className="page">
      <header>
        <BackButton />
        <h1>{detail?.name}</h1>
      </header>

      {!loading ? (
        <div className="flex-col">
          <div>
            <h2>Rotation Period</h2>
            <p>{detail?.rotation_period}</p>
          </div>

          <div>
            <h2>Orbital Period</h2>
            <p>{detail?.orbital_period}</p>
          </div>

          <div>
            <h2>Terrain</h2>
            <p>{detail?.terrain}</p>
          </div>

          <div>
            <h2>Population</h2>
            <p>{detail?.population}</p>
          </div>

          <div>
            <h2>Diameter</h2>
            <p>{detail?.diameter}</p>
          </div>

          <div>
            <h2>Climate</h2>
            <p>{detail?.climate}</p>
          </div>
        </div>
      ) : (
        <h2>Loading...</h2>
      )}
    </div>
  );
};

export default PlanetDetail;
