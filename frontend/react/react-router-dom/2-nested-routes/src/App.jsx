import React from "react";
import { NavLink, Outlet } from "react-router-dom";
import "./App.css";

const App = () => (
  <>
    {/* 2. Semua route yang menjadi children route '/' akan merender element <nav> */}
    <nav>
      <NavLink
        style={({ isActive }) => ({
          fontWeight: isActive ? 600 : 400,
          color: isActive ? "firebrick" : "steelblue",
          paddingBlock: 8,
          paddingInline: 16,
          borderRadius: 8,
        })}
        to="/"
      >
        Main Menu
      </NavLink>
      <NavLink
        style={({ isActive }) => ({
          fontWeight: isActive ? 600 : 400,
          color: isActive ? "firebrick" : "steelblue",
          paddingBlock: 8,
          paddingInline: 16,
          borderRadius: 8,
        })}
        to="/pokemons"
      >
        Pokemons
      </NavLink>
    </nav>

    {/* 3. Semua route yang menjadi children route '/' akan me-replace komponen <Outlet /> */}
    <Outlet />
  </>
);

export default App;
