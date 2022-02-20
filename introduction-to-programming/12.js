const CELL_SIZE = 20;
const CANVAS_SIZE = 400;
const REDRAW_INTERVAL = 100;
const WIDTH = CANVAS_SIZE / CELL_SIZE;
const HEIGHT = CANVAS_SIZE / CELL_SIZE;

let snake1 = {
    //this
    color: "purple",
    position: {
        x: Math.floor(Math.random() * WIDTH),
        y: Math.floor(Math.random() * HEIGHT),
    }
}
let snake2 = {
    color: "blue",
    position: {
        x: Math.floor(Math.random() * WIDTH),
        y: Math.floor(Math.random() * HEIGHT),
    }
}

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let ctx = snakeCanvas.getContext("2d");

        ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
        
        ctx.fillStyle = snake1.color;
        ctx.fillRect(snake1.position.x * CELL_SIZE, snake1.position.y * CELL_SIZE, CELL_SIZE, CELL_SIZE);

        ctx.fillStyle = snake2.color;
        ctx.fillRect(snake2.position.x * CELL_SIZE, snake2.position.y * CELL_SIZE, CELL_SIZE, CELL_SIZE);
    }, REDRAW_INTERVAL);
}

function teleport(snake) {
    if (snake.position.x < 0) {
        snake.position.x = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake.position.x >= WIDTH) {
        snake.position.x = 0;
    }
    if (snake.position.y < 0) {
        snake.position.y = CANVAS_SIZE / CELL_SIZE - 1;
    }
    if (snake.position.y >= HEIGHT) {
        snake.position.y = 0;
    }
}

function moveLeft(snake) {
    snake.position.x--;
    teleport(snake);
}

function moveRight(snake) {
    snake.position.x++;
    teleport(snake);
}

function moveDown(snake) {
    snake.position.y++;
    teleport(snake);
}

function moveUp(snake) {
    snake.position.y--;
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

