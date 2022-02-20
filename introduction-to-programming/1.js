let snake = {
    position: { x: 0, y: 0 },
    moveDown: function() { 
        this.position.y++; 
        this.teleport();
    },
    moveUp: function() { 
        this.position.y--; 
        this.teleport();
    },
    moveLeft: function() { 
        this.position.x--;
        this.teleport();
    },
    moveRight: function() { 
        this.position.x++; 
        this.teleport();
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
    }, 100)
}

document.addEventListener("keydown", function (event) {
    if (event.defaultPrevented) {
        return;
    }
    
    switch (event.key) {
    case "ArrowDown":
        snake.moveDown();
        break;
    case "ArrowUp":
        snake.moveUp();
        break;
    case "ArrowLeft":
        snake.moveLeft();
        break;
    case "ArrowRight":
        snake.moveRight();
        break;
    default:
        return; 
    }
})

init();