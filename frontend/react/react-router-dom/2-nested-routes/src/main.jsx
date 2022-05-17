import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import MainMenu from "./routes/MainMenu";
import Pokemons from "./routes/Pokemons";
import PokemonDetail from "./routes/PokemonDetail";
import NotFound from "./routes/NotFound";
import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        {/* 1. localhost:3000 akan merender <App /> */}
        <Route path="/" element={<App />}>
          {/* 4. localhost:3000 akan merender <App /> dan <Outlet/> akan digantikan dengan <MainMenu /> */}
          <Route index element={<MainMenu />} />
          {/* 5. localhost:3000/pokemons akan merender <App /> dan <Outlet/> akan digantikan dengan <Pokemons /> */}
          <Route path="pokemons" element={<Outlet />}>
            <Route index element={<Pokemons />} />
            {/* 6. localhost:3000/pokemons/quilava, localhost:3000/pokemons/totodile, atau localhost:3000/pokemons/apapun akan merender <App /> dan <Outlet/> akan digantikan dengan <PokemonDetail /> */}
            <Route path=":name" element={<PokemonDetail />} />
          </Route>
        </Route>
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);
