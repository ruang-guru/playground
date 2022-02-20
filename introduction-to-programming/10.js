const CELL_SIZE = 20;
const SNAKE_1_COLOR = "purple";
const SNAKE_2_COLOR = "blue";
const CANVAS_SIZE = 400;
const REDRAW_INTERVAL = 100;
const WIDTH = CANVAS_SIZE / CELL_SIZE;
const HEIGHT = CANVAS_SIZE / CELL_SIZE;

let snake1PositionX = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);
let snake1PositionY = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);
let snake2PositionX = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);
let snake2PositionY = Math.floor(Math.random() * CANVAS_SIZE / CELL_SIZE);

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let ctx = snakeCanvas.getContext("2d");

        ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
        ctx.fillStyle = SNAKE_1_COLOR;
        ctx.fillRect(snake1PositionX * CELL_SIZE, snake1PositionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);

        ctx.fillStyle = SNAKE_2_COLOR;
        ctx.fillRect(snake2PositionX * CELL_SIZE, snake2PositionY * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    }, REDRAW_INTERVAL);
}

function teleportSnake1() {
    if (snake1PositionX < 0) {
        snake1PositionX = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake1PositionX >= WIDTH) {
        snake1PositionX = 0;
    }
    if (snake1PositionY < 0) {
        snake1PositionY = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake1PositionY >= HEIGHT) {
        snake1PositionY = 0;
    }
}

function moveSnake1Left() {
    snake1PositionX--;
    teleportSnake1();
}

function moveSnake1Right() {
    snake1PositionX++;
    teleportSnake1();
}

function moveSnake1Down() {
    snake1PositionY++;
    teleportSnake1();
}

function moveSnake1Up() {
    snake1PositionY--;
    teleportSnake1();
}

function teleportSnake2() {
    if (snake2PositionX < 0) {
        snake2PositionX = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake2PositionX >= WIDTH) {
        snake2PositionX = 0;
    }
    if (snake2PositionY < 0) {
        snake2PositionY = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake2PositionY >= HEIGHT) {
        snake2PositionY = 0;
    }
}

function moveSnake2Left() {
    snake2PositionX--;
    teleportSnake2();
}

function moveSnake2Right() {
    snake2PositionX++;
    teleportSnake2();
}

function moveSnake2Down() {
    snake2PositionY++;
    teleportSnake2();
}

function moveSnake2Up() {
    snake2PositionY--;
    teleportSnake2();
}

document.addEventListener("keydown", function (event) {
    if (event.key === "ArrowLeft") {
        moveSnake1Left(); 
    } else if (event.key === "ArrowRight") {
        moveSnake1Right();
    } else if (event.key === "ArrowUp") {
        moveSnake1Up(); 
    } else if (event.key === "ArrowDown") {
        moveSnake1Down();
    }

    if (event.key === "a") {
        moveSnake2Left();
    } else if (event.key === "d") {
        moveSnake2Right();
    } else if (event.key === "w") {
        moveSnake2Up();
    } else if (event.key === "s") {
        moveSnake2Down();
    }
})

