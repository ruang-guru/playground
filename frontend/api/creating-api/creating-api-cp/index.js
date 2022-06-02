import express from "express";
import bodyParser from "body-parser";

import booksRoutes from "./routes/books.js";

const app = express();
const PORT = process.env.PORT || 5000;

app.use(bodyParser.json());

app.use("/books", booksRoutes);
app.get("/", (req, res) => res.send("Welcome to the books API!"));
app.all("*", (req, res) =>
  res.send("You've tried reaching a route that doesn't exist.")
);

const server = app.listen(PORT, () =>
  console.log(`Server running on port: http://localhost:${PORT}`)
);

export default server;
