import { useState } from "react";

import "./App.css";
import MoviePages from "./screens/MoviePages";
import RegisterPage from "./screens/RegisterPage";
import Header from "./components/Header";

const defineHeaderText = (page, userData) => {
  if (page == "register") {
    return "Welcome to Marvel Movie Page !";
  }
  return `Welcome, ${userData.userName}`;
};

const App = () => {
  const [page, setPage] = useState("register");
  const [userData, setUserData] = useState({});

  return (
    <div id="root" className="App">
      <Header headerText={defineHeaderText(page, userData)} />
      {page == "register" ? (
        <RegisterPage
          userData={userData}
          setUserData={setUserData}
          setPage={setPage}
        />
      ) : (
        <MoviePages userData={userData} />
      )}
    </div>
  );
};

export default App;
