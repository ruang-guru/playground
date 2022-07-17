import React, { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import BackButton from "../components/BackButton";

const MovieDetail = () => {
  const [detail, setDetail] = useState(null);
  const [loading, setLoading] = useState(false);
  // TODO: answer here
  const controller = new AbortController();

  const loadDetail = async () => {
    setLoading(true);
    try {
      const url = /* beginanswer */ "https://swapi.dev/api/films/" + id; /* endanswer "" */
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
        <h1>{detail?.title}</h1>
      </header>

      {!loading ? (
        <div className="flex-col">
          <div>
            <h2>Release Date</h2>
            <p>{detail?.release_date}</p>
          </div>

          <div>
            <h2>Produced By</h2>
            <p>{detail?.producer}</p>
          </div>

          <div>
            <h2>Directed By</h2>
            <p>{detail?.director}</p>
          </div>

          <div>
            <h2>Summary</h2>
            <p>{detail?.opening_crawl}</p>
          </div>
        </div>
      ) : (
        <h2>Loading...</h2>
      )}
    </div>
  );
};

export default MovieDetail;
