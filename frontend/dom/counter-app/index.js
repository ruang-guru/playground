let number = document.getElementById('number')
let addbtn = document.getElementsByClassName('btn-add')[0]
let subtractbtn = document.getElementsByClassName('btn-subtract')[0]
let message = document.getElementsByClassName('message')[0]

addbtn.addEventListener('click', addCounter)
subtractbtn.addEventListener('click', subtractCounter)


function addCounter() {
  number.innerHTML = parseInt(number.innerHTML) + 1
}

function subtractCounter() {
  if (number.innerHTML == 0) {
    message.innerHTML = "Oops! you reach the min value!"
    setTimeout(clearMessage, '3000')
  } else if (number.innerHTML > 0) {
    number.innerHTML = parseInt(number.innerHTML) - 1
  }
}

function clearMessage() {
  message.innerHTML = ''
}