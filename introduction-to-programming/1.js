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
        if (this.position.x >= width) {
            this.position.x = 0;
        }
        if (this.position.y >= height) {
            this.position.y = 0;
        }
    }
}

let size = 20;
let canvasSize = 400;
let height = canvasSize / size;
let width = canvasSize / size;

function random(minValue, maxValue) {
    return Math.floor(Math.random() * (maxValue - minValue + 1)) + minValue;
}
function initSnake() {
    snake.position.x = random(0, width);
    snake.position.y = random(0, height);
}

function init() {
    initSnake();
}

function draw() {
    setInterval(function() {
        let canvas = document.getElementById("snakeBoard");
        let ctx = canvas.getContext("2d");
        ctx.clearRect(0, 0, canvasSize, canvasSize);

        ctx.fillRect(snake.position.x * size, snake.position.y * size, size, size);
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