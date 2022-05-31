import request from "supertest";
import server from "../index.js";

describe("Sample Test", () => {
  it("should test that true === true", () => {
    expect(true).toBe(true);
  });
});

describe("Update and Delete Books ", () => {
  afterAll((done) => {
    server.close()
    done()
  })
  test("It should add new book", async () => {
    const newBook = {
      title: "Ruangguru Book",
      author: "Ruangguru Writer"
    };

    const response = await request(server).post("/books").send(newBook);
    expect(response.statusCode).toEqual(201);
    expect(response.body.success).toEqual(true);
    expect(response.body.message).toEqual(
      `Book [${newBook.title}] added to the database.`
    );
  });

  test("It should update title of a book", async () => {
    const getBookId = await request(server).get("/books");
    const bookId = getBookId.body.data[1].id;

    const updateABook = {
      title: "Updated A Book",
      author: "Updated A Writer"
    };

    const response = await request(server)
      .patch(`/books/${bookId}`)
      .send(updateABook);
    expect(response.statusCode).toEqual(200);
    expect(response.body.success).toEqual(true);
    expect(response.body.message).toEqual(
      `title has been updated to ${updateABook.title}.author has been updated to ${updateABook.author}`
    );
  });

  test("It should delete a book", async () => {
    const getBookId = await request(server).get("/books");
    const bookId = getBookId.body.data[1].id;

    const response = await request(server).delete(`/books/${bookId}`);
    expect(response.statusCode).toEqual(200);
    expect(response.body.success).toEqual(true);
    expect(response.body.message).toEqual(
      `book with id ${bookId} has been deleted`
    );

    const deletedBook = await request(server).delete(`/books/${1234567890}`);
    expect(deletedBook.statusCode).toEqual(404);
    expect(deletedBook.body.message).toEqual(
      `book with id ${1234567890} not found`
    );
  });
});
