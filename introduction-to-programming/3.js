const CELL_SIZE = 20;
const SNAKE_COLOR = "purple";

let snakePositionX = 3;
let snakePositionY = 5;

function draw() {
    let snakeCanvas = document.getElementById("snakeBoard");
    let ctx = snakeCanvas.getContext("2d");

    ctx.fillStyle = SNAKE_COLOR;
    ctx.fillRect(snakePositionX * CELL_SIZE, snakePositionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);

    //play with this
    snakePositionX = snakePositionX + 1;
    snakePositionY = snakePositionY + 1;

    ctx.fillRect(snakePositionX * CELL_SIZE, snakePositionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);
}
