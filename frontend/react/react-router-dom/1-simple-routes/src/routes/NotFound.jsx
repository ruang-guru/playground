import React from "react";
import { useNavigate } from "react-router-dom";

const NotFound = () => {
  const navigate = useNavigate();

  return (
    <main>
      <h1 className="grey">Not Found </h1>
      <button onClick={() => navigate(-1)}>Take me back</button>
    </main>
  );
};

export default NotFound;
