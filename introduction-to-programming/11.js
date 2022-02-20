const CELL_SIZE = 20;
const SNAKE_1_COLOR = "purple";
const SNAKE_2_COLOR = "blue";
const CANVAS_SIZE = 400;
const REDRAW_INTERVAL = 100;
const WIDTH = CANVAS_SIZE / CELL_SIZE;
const HEIGHT = CANVAS_SIZE / CELL_SIZE;

let snake1 = {
    positionX: Math.floor(Math.random() * WIDTH),
    positionY: Math.floor(Math.random() * HEIGHT),
}
let snake2 = {
    positionX: Math.floor(Math.random() * WIDTH),
    positionY: Math.floor(Math.random() * HEIGHT),
}

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let ctx = snakeCanvas.getContext("2d");

        ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
        
        ctx.fillStyle = SNAKE_1_COLOR;
        ctx.fillRect(snake1.positionX * CELL_SIZE, snake1.positionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);

        ctx.fillStyle = SNAKE_2_COLOR;
        ctx.fillRect(snake2.positionX * CELL_SIZE, snake2.positionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    }, REDRAW_INTERVAL);
}

function teleport(snake) {
    if (snake.positionX < 0) {
        snake.positionX = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake.positionX >= WIDTH) {
        snake.positionX = 0;
    }
    if (snake.positionY < 0) {
        snake.positionY = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake.positionY >= HEIGHT) {
        snake.positionY = 0;
    }
}

function moveLeft(snake) {
    snake.positionX--;
    teleport(snake);
}

function moveRight(snake) {
    snake.positionX++;
    teleport(snake);
}

function moveDown(snake) {
    snake.positionY++;
    teleport(snake);
}

function moveUp(snake) {
    snake.positionY--;
    teleport(snake);
}

document.addEventListener("keydown", function (event) {
    if (event.key === "ArrowLeft") {
        moveLeft(snake1); 
    } else if (event.key === "ArrowRight") {
        moveRight(snake1); 
    } else if (event.key === "ArrowUp") {
        moveUp(snake1); 
    } else if (event.key === "ArrowDown") {
        moveDown(snake1); 
    }

    if (event.key === "a") {
        moveLeft(snake2); 
    } else if (event.key === "d") {
        moveRight(snake2); 
    } else if (event.key === "w") {
        moveUp(snake2); 
    } else if (event.key === "s") {
        moveDown(snake2); 
    }
})

