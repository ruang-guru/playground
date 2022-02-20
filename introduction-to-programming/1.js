const MOVE_INTERVAL = 150;
const DRAW_INTERVAL = 100;

const DIRECTION = {
    UP: 0,
    RIGHT: 1,
    DOWN: 2,
    LEFT: 3
}

let snake = {
    head: { x: 0, y: 0 },
    body: [{ x: 0, y: 0 }],
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
        this.head.y++; 
        this.teleport();
        this.eat();
        this.moveBody();
    },
    moveUp: function() { 
        this.head.y--; 
        this.teleport();
        this.eat();
        this.moveBody();
    },
    moveLeft: function() { 
        this.head.x--;
        this.teleport();
        this.eat();
        this.moveBody();
    },
    moveRight: function() { 
        this.head.x++; 
        this.teleport();
        this.eat();
        this.moveBody();
    },
    moveBody: function() {
        this.body.unshift({ x: this.head.x, y: this.head.y });
        this.body.pop();
    },
    turn: function(direction) {
        this.direction = direction;
    },
    eat: function() {
        if (this.head.x == apple.position.x && this.head.y == apple.position.y) {
            score++;
            this.body.push(this.head)
            initApple();
        }
    },
    teleport: function() {
        if (this.head.x < 0) {
            this.head.x = width - 1;
        }
        if (this.head.y < 0) {
            this.head.y = height - 1;
        }
        if (this.head.x == width) {
            this.head.x = 0;
        }
        if (this.head.y == height) {
            this.head.y = 0;
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
    snake.head.x = random(0, width - 1);
    snake.head.y = random(0, height - 1);
    snake.body = [{ x: snake.head.x, y: snake.head.y }];
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