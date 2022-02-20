function draw() {
    let snakeCanvas = document.getElementById("snakeBoard");
    let ctx = snakeCanvas.getContext("2d");

    //play with this
    ctx.fillStyle = "purple";
    ctx.fillRect(60, 100, 20, 20);

    ctx.fillStyle = "red";
    ctx.fillRect(20, 40, 20, 20);

    ctx.fillStyle = "green";
    ctx.fillRect(0, 0, 20, 20);
}
