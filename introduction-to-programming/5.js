const CELL_SIZE = 20;
const SNAKE_COLOR = "purple";
const CANVAS_SIZE = 400;

//play with this
const REDRAW_INTERVAL = 100;

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let ctx = snakeCanvas.getContext("2d");

        let snakePositionX = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);
        let snakePositionY = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);

        //uncomment this
        //ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);

        ctx.fillStyle = SNAKE_COLOR;
        ctx.fillRect(snakePositionX * CELL_SIZE, snakePositionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    }, REDRAW_INTERVAL);
}
