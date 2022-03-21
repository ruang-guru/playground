const CELL_SIZE = 20;
const CANVAS_SIZE = 400;
const REDRAW_INTERVAL = 100;
const WIDTH = CANVAS_SIZE / CELL_SIZE;
const HEIGHT = CANVAS_SIZE / CELL_SIZE;

function initPosition() {
    return {
        x: Math.floor(Math.random() * WIDTH),
        y: Math.floor(Math.random() * HEIGHT),
    }
}

let snake1 = {
    color: "purple",
    position: initPosition(),
    score: 0,
}
let snake2 = {
    color: "blue",
    position: initPosition(),
    score: 0,
}
let apple = {
    color: "red",
    position: initPosition(),
}

function drawCell(ctx, x, y, color) {
    ctx.fillStyle = color;
    ctx.fillRect(x * CELL_SIZE, y * CELL_SIZE, CELL_SIZE, CELL_SIZE);
}

function drawScore(snake) {
    let scoreCanvas;
    if (snake.color == snake1.color) {
        scoreCanvas = document.getElementById("score1Board");
    } else {
        scoreCanvas = document.getElementById("score2Board");
    }
    let scoreCtx = scoreCanvas.getContext("2d");

    scoreCtx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
    scoreCtx.font = "30px Arial";
    scoreCtx.fillStyle = snake.color
    scoreCtx.fillText(snake.score, 10, scoreCanvas.scrollHeight / 2);
}

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let ctx = snakeCanvas.getContext("2d");

        ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
        
        drawCell(ctx, snake1.position.x, snake1.position.y, snake1.color);
        drawCell(ctx, snake2.position.x, snake2.position.y, snake2.color);
        drawCell(ctx, apple.position.x, apple.position.y, apple.color);

        drawScore(snake1);
        drawScore(snake2);
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

function eat(snake, apple) {
    if (snake.position.x == apple.position.x && snake.position.y == apple.position.y) {
        apple.position = initPosition();
        snake.score++;
    }
}

function moveLeft(snake) {
    snake.position.x--;
    teleport(snake);
    eat(snake, apple);
}

function moveRight(snake) {
    snake.position.x++;
    teleport(snake);
    eat(snake, apple);
}

function moveDown(snake) {
    snake.position.y++;
    teleport(snake);
    eat(snake, apple);
}

function moveUp(snake) {
    snake.position.y--;
    teleport(snake);
    eat(snake, apple);
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

