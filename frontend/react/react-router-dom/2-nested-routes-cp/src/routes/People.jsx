import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import peopleThumbnail from "../assets/people.jpeg";
import BackButton from "../components/BackButton";

const People = () => {
  const [people, setPeople] = useState([]);
  const [loading, setLoading] = useState(false);
  const controller = new AbortController();

  const loadPeople = async () => {
    setLoading(true);
    try {
      const { data } = await axios.get("https://swapi.dev/api/people", {
        signal: controller.signal,
      });
      setPeople(data.results);
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
    loadPeople();
    return () => {
      controller.abort();
    };
  }, []);

  return (
    <div className="page">
      <header>
        <BackButton />
        <img
          src={peopleThumbnail}
          width={40}
          height={40}
          alt="People Thumbnail"
        />
        <h1>People</h1>
      </header>

      {!loading ? (
        <div className="grid">
          {people.length &&
            people.map((person, index) => (
              <Link to={`/star-wars/people/${getID(person.url)}`} key={index}>
                <div className="card">
                  <p>{person.name}</p>
                </div>
              </Link>
            ))}
        </div>
      ) : (
        <h2>Loading...</h2>
      )}
    </div>
  );
};

export default People;
