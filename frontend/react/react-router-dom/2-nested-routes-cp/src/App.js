import React from "react";
import { Routes, Route } from "react-router-dom";
import Home from "./routes/Home";
import Main from "./layouts/Main";
import People from "./routes/People";
import PeopleDetail from "./routes/PeopleDetail";
// TODO: answer here
import NotFound from "./routes/NotFound";
import "./App.css";

const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="star-wars" element={<Main />}>
        <Route path="people">
          <Route index element={<People />} />
          <Route path=":id" element={<PeopleDetail />} />
        </Route>
        {/* TODO: answer here */}
        <Route path="*" element={<NotFound />} />
      </Route>
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
};

export default App;
