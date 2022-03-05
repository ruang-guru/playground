const CELL_SIZE = 20;
const CANVAS_SIZE = 400;
const MAX_X = CANVAS_SIZE / CELL_SIZE; // 20
const MAX_Y = CANVAS_SIZE / CELL_SIZE; // 20
const DIRECTION = {
    LEFT: 0,
    UP: 1,
    RIGHT: 2,
    DOWN: 3
}
const MOVE_INTERVAL = 120;

function initPosition() {
    return {
        x: Math.floor(Math.random() * MAX_X),
        y: Math.floor(Math.random() * MAX_Y),
    }
}
function initDirection() {
    return Math.floor(Math.random() * 4);
}

let snake1 = {
    color: "purple",
    position: initPosition(),
    score: 0,
    direction: initDirection(),
}
let snake2 = {
    color: "blue",
    position: initPosition(),
    score: 0,
    direction: initDirection(),
}
let apple = {
    color: "red",
    position: initPosition(),
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

        //clear canvas
        ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);

        //draw snake1
        ctx.fillStyle = snake1.color;
        ctx.fillRect(snake1.position.x * CELL_SIZE, snake1.position.y * CELL_SIZE, CELL_SIZE, CELL_SIZE);

        //draw snake2
        ctx.fillStyle = snake2.color;
        ctx.fillRect(snake2.position.x * CELL_SIZE, snake2.position.y * CELL_SIZE, CELL_SIZE, CELL_SIZE);

        //draw apple
        ctx.fillStyle = apple.color;
        ctx.fillRect(apple.position.x * CELL_SIZE, apple.position.y * CELL_SIZE, CELL_SIZE, CELL_SIZE);

        //draw score
        drawScore(snake1);
        drawScore(snake2);

    }, 5);
}

function teleport(snake) {
    if (snake.position.y == -1) {
        snake.position.y = MAX_Y - 1;
    }
    if (snake.position.y == MAX_Y) {
        snake.position.y = 0;
    }
    if (snake.position.x == -1) {
        snake.position.x = MAX_X - 1;
    }
    if (snake.position.x == MAX_X) {
        snake.position.x = 0;
    }
}

function moveUp(snake) {
    snake.position.y = snake.position.y - 1;
    teleport(snake);
    eat(snake, apple);
}

function moveDown(snake) {
    snake.position.y = snake.position.y + 1;
    teleport(snake);
    eat(snake, apple);
}

function moveLeft(snake) {
    snake.position.x = snake.position.x - 1;
    teleport(snake);
    eat(snake, apple);
}

function moveRight(snake) {
    snake.position.x = snake.position.x + 1;
    teleport(snake);
    eat(snake, apple);
}

function move(snake) {
    setInterval(function() {
        if (snake.direction == DIRECTION.LEFT) {
            moveLeft(snake);
        } else if (snake.direction == DIRECTION.RIGHT) {
            moveRight(snake);
        } else if (snake.direction == DIRECTION.DOWN) {
            moveDown(snake);
        } else if (snake.direction == DIRECTION.UP) {
            moveUp(snake);
        }
    }, MOVE_INTERVAL);
}

function eat(snake, apple) {
    if (snake.position.x == apple.position.x && snake.position.y == apple.position.y) {
        //posisi sama -> sedang makan
        snake.score = snake.score + 1;
        console.log(snake.color, snake.score);
        apple.position = initPosition();
    }
}

document.addEventListener("keydown", function(event) {
    if (event.key == "ArrowUp") {
        snake1.direction = DIRECTION.UP;
    } else if (event.key == "ArrowDown") {
        snake1.direction = DIRECTION.DOWN;
    } else if (event.key == "ArrowLeft") {
        snake1.direction = DIRECTION.LEFT;
    } else if (event.key == "ArrowRight") {
        snake1.direction = DIRECTION.RIGHT;
    }
        
    if (event.key == "w") {
        snake2.direction = DIRECTION.UP;
    } else if (event.key == "s") {
        snake2.direction = DIRECTION.DOWN;
    } else if (event.key == "a") {
        snake2.direction = DIRECTION.LEFT;
    } else if (event.key == "d") {
        snake2.direction = DIRECTION.RIGHT;
    }
})

move(snake1);
move(snake2);