const MOVE_INTERVAL = 150;
const DRAW_INTERVAL = 100;

const DIRECTION = {
    UP: 0,
    RIGHT: 1,
    DOWN: 2,
    LEFT: 3
}

let snake = {
    position: { x: 0, y: 0 },
    direction: DIRECTION.RIGHT,
    move: function() {
        let actions = {
            [DIRECTION.UP]: this.moveUp,
            [DIRECTION.RIGHT]: this.moveRight,
            [DIRECTION.DOWN]: this.moveDown,
            [DIRECTION.LEFT]: this.moveLeft
        }
        actions[this.direction].bind(this)();
        setTimeout(this.move.bind(this), MOVE_INTERVAL);
    },
    moveDown: function() { 
        this.position.y++; 
        this.teleport();
        this.eat();
    },
    moveUp: function() { 
        this.position.y--; 
        this.teleport();
        this.eat();
    },
    moveLeft: function() { 
        this.position.x--;
        this.teleport();
        this.eat();
    },
    moveRight: function() { 
        this.position.x++; 
        this.teleport();
        this.eat();
    },
    turn: function(direction) {
        this.direction = direction;
    },
    eat: function() {
        if (this.position.x == apple.position.x && this.position.y == apple.position.y) {
            score++;
            initApple();
        }
    },
    teleport: function() {
        if (this.position.x < 0) {
            this.position.x = width - 1;
        }
        if (this.position.y < 0) {
            this.position.y = height - 1;
        }
        if (this.position.x == width) {
            this.position.x = 0;
        }
        if (this.position.y == height) {
            this.position.y = 0;
        }
    }
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
    snake.position.x = random(0, width - 1);
    snake.position.y = random(0, height - 1);
    snake.direction = random(0, 3);
    snake.move();
}

function initApple() {
    apple.position.x = random(0, width - 1);
    apple.position.y = random(0, height - 1);
}

function init() {
    initSnake();
    initApple();
}

function draw() {
    setInterval(function() {
        let snakeCanvas = document.getElementById("snakeBoard");
        let boardCtx = snakeCanvas.getContext("2d");
        boardCtx.clearRect(0, 0, canvasSize, canvasSize);

        boardCtx.fillStyle = "purple";
        boardCtx.fillRect(snake.position.x * size, snake.position.y * size, size, size);

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
        snake.turn(DIRECTION.DOWN);
        break;
    case "ArrowUp":
        snake.turn(DIRECTION.UP);
        break;
    case "ArrowLeft":
        snake.turn(DIRECTION.LEFT);
        break;
    case "ArrowRight":
        snake.turn(DIRECTION.RIGHT);
        break;
    default:
        return; 
    }
})

init();