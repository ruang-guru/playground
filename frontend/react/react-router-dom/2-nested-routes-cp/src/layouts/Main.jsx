import React from "react";
import { NavLink, Link, Outlet } from "react-router-dom";
import logo from "../assets/star-wars-logo.png";

const style = (isActive) => ({
  fontWeight: isActive ? 600 : 400,
  color: isActive ? "gold" : "white",
  paddingBlock: 8,
  paddingInline: 16,
  borderRadius: 8,
});

const PATHS = [
  { to: "/star-wars/people", label: "People" },
  { to: "/star-wars/planets", label: "Planets" },
  { to: "/star-wars/movies", label: "Movies" },
];

const Main = () => {
  return (
    <>
      <nav>
        <Link to="/">
          <img src={logo} width={64} alt="Logo" />
        </Link>
        <div>
          {PATHS.map((path) => (
            <NavLink
              key={path.to}
              to={path.to}
              style={({ isActive }) => style(isActive)}
            >
              {path.label}
            </NavLink>
          ))}
        </div>
      </nav>
      <Outlet />
    </>
  );
};

export default Main;
