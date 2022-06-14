import React, { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import BackButton from "../components/BackButton";

const PeopleDetail = () => {
  const [detail, setDetail] = useState(null);
  const [loading, setLoading] = useState(false);
  const { id } = useParams();
  const controller = new AbortController();

  const loadDetail = async () => {
    setLoading(true);
    try {
      const url = "https://swapi.dev/api/people/" + id;
      const { data } = await axios.get(url, {
        signal: controller.signal,
      });
  
      setDetail(data);
  
    } catch (error) {
      console.log(error);
    }
    setLoading(false);
  };

  useEffect(() => {

    loadDetail();

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
            <h2>Gender</h2>
            <p>{detail?.gender}</p>
          </div>

          <div>
            <h2>Birth Year</h2>
            <p>{detail?.birth_year}</p>
          </div>

          <div>
            <h2>Mass</h2>
            <p>{detail?.mass} kg</p>
          </div>

          <div>
            <h2>Height</h2>
            <p>{detail?.height} cm</p>
          </div>

          <div>
            <h2>Eye Color</h2>
            <p>{detail?.eye_color}</p>
          </div>
        </div>
      ) : (
        <h2>Loading...</h2>
      )}
    </div>
  );
};

export default PeopleDetail;
