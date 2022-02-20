//play with this
const CELL_SIZE = 20;
const SNAKE_COLOR = "purple";

function draw() {
    let snakeCanvas = document.getElementById("snakeBoard");
    let ctx = snakeCanvas.getContext("2d");

    ctx.fillStyle = SNAKE_COLOR;
    ctx.fillRect(3 * CELL_SIZE, 5 * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    ctx.fillRect(1 * CELL_SIZE, 2 * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    ctx.fillRect(0 * CELL_SIZE, 0 * CELL_SIZE, CELL_SIZE, CELL_SIZE);
}
