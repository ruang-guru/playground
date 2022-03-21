const CELL_SIZE = 20;
const SNAKE_COLOR = "purple";
const CANVAS_SIZE = 400;

let snakePositionX = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);
let snakePositionY = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);

function draw() {
    let snakeCanvas = document.getElementById("snakeBoard");
    let ctx = snakeCanvas.getContext("2d");

    ctx.fillStyle = SNAKE_COLOR;
    ctx.fillRect(snakePositionX * CELL_SIZE, snakePositionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);
}
