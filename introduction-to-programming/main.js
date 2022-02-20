const MOVE_INTERVAL = 120;
const DRAW_INTERVAL = 50;

const DIRECTION = {
    UP: 0,
    RIGHT: 1,
    DOWN: 2,
    LEFT: 3
}

function move(snake) {
    switch (snake.direction) {
        case DIRECTION.UP:
            moveUp(snake);
            break;
        case DIRECTION.RIGHT:
            moveRight(snake);
            break;
        case DIRECTION.DOWN:
            moveDown(snake);
            break;
        case DIRECTION.LEFT:
            moveLeft(snake);
            break;
    }

    teleport(snake);
    eat(snake);
    moveBody(snake);
    checkCollision(snake);
    setTimeout(function() {
        move(snake)
    }, MOVE_INTERVAL);
}
function moveDown(snake) { snake.head.y++; }
function moveUp(snake) { snake.head.y--; }
function moveLeft(snake) { snake.head.x--; }
function moveRight(snake) { snake.head.x++; }

function checkCollision(snake) {
    let isCollide = false;
    for (let i = 1; i < snake.body.length; i++) {
        if (snake.body[i].x == snake.head.x && snake.body[i].y == snake.head.y) {
            isCollide = true;
            alert("Game over")
        }
    }
    if (isCollide) {
        initSnake();
        score = 0;
    }
}

function moveBody(snake) {
    snake.body.unshift({ x: snake.head.x, y: snake.head.y });
    snake.body.pop();
}

function turn(snake, direction) {
    let oppositeDirections = {
        [DIRECTION.UP]: DIRECTION.DOWN,
        [DIRECTION.RIGHT]: DIRECTION.LEFT,
        [DIRECTION.DOWN]: DIRECTION.UP,
        [DIRECTION.LEFT]: DIRECTION.RIGHT
    }

    if (direction !== oppositeDirections[snake.direction]) {
        snake.direction = direction;
    }
};

function eat(snake) {
    if (snake.head.x == apple.position.x && snake.head.y == apple.position.y) {
        score++;
        snake.body.push(snake.head)
        initApple();
    }
}

function teleport(snake) {
    if (snake.head.x < 0) {
        snake.head.x = width - 1;
    }
    if (snake.head.y < 0) {
        snake.head.y = height - 1;
    }
    if (snake.head.x == width) {
        snake.head.x = 0;
    }
    if (snake.head.y == height) {
        snake.head.y = 0;
    }
}

let snake = {
    head: { x: 0, y: 0 },
    body: [{ x: 0, y: 0 }],
    direction: DIRECTION.RIGHT,
}

let apple = {
    position: { x: 0, y: 0 },
}

let size = 20;
let canvasSize = 400;
let height = canvasSize / size;
let width = canvasSize / size;
let score = 0;

function random(minValue, maxValue) {
    return Math.floor(Math.random() * (maxValue - minValue + 1)) + minValue;
}
function initSnake() {
    snake.head.x = random(0, width - 1);
    snake.head.y = random(0, height - 1);
    snake.body = [{ x: snake.head.x, y: snake.head.y }];
    snake.direction = random(0, 3);
}

function initApple() {
    apple.position.x = random(0, width - 1);
    apple.position.y = random(0, height - 1);
}

function init() {
    initSnake();
    initApple();
    move(snake);
}

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let boardCtx = snakeCanvas.getContext("2d");
        boardCtx.clearRect(0, 0, canvasSize, canvasSize);


        boardCtx.fillStyle = "purple";
        boardCtx.fillRect(snake.body[0].x * size, snake.body[0].y * size, size, size);

        for (let i = 1; i < snake.body.length; i++) {
            boardCtx.fillStyle = "green";
            boardCtx.fillRect(snake.body[i].x * size, snake.body[i].y * size, size, size);
        }

        boardCtx.fillStyle = "red";
        boardCtx.fillRect(apple.position.x * size, apple.position.y * size, size, size);

        let scoreCanvas = document.getElementById("score");
        let scoreCtx = scoreCanvas.getContext("2d");

        scoreCtx.clearRect(0, 0, scoreCanvas.scrollWidth, scoreCanvas.scrollHeight);
        scoreCtx.font = "30px Arial";
        scoreCtx.fillText(score, 10, scoreCanvas.scrollHeight / 2);
    }, DRAW_INTERVAL)
}

document.addEventListener("keydown", function (event) {
    if (event.defaultPrevented) {
        return;
    }
    
    switch (event.key) {
    case "ArrowDown":
        turn(snake, DIRECTION.DOWN);
        break;
    case "ArrowUp":
        turn(snake, DIRECTION.UP);
        break;
    case "ArrowLeft":
        turn(snake, DIRECTION.LEFT);
        break;
    case "ArrowRight":
        turn(snake, DIRECTION.RIGHT);
        break;
    default:
        return; 
    }
})

init();