import express from "express";

import { getBooks, createBook, getBook } from "../controllers/books.js";

const router = express.Router();

router.get("/", getBooks);

router.post("/", createBook);

router.get("/:id", getBook);

export default router;
