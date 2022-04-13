class Player {
  constructor(name) {
    this.name = name;
    this.gold = 50;
    this.stealChance = 0.25;
    this.hasBeenRobbed = false;
  }


  getGold() {
  // TODO: answer here
  }

  setGold(gold) {
  // TODO: answer here
  }

  getStealChance() {
  // TODO: answer here
  }

  setStealChance(chance) {
  // TODO: answer here
  }

  getHasBeenRobbed() {
    return this.hasBeenRobbed;
  }
  
  setHasBeenRobbed(status) {
    this.hasBeenRobbed = status;
  }
  
  randomizer() {
    return +Math.random().toFixed(2);
  }
  
  steal(player) {
    // TODO: answer here
  }
}

const p1 = new Player("Fauzan");
const p2 = new Player("Tegar");

console.log('Jumlah gold Player 1: ' + p1.getGold())
console.log('Jumlah gold Player 2: ' + p2.getGold())
console.log(p1.steal(p2));
console.log('Total gold Player 1: ' + p1.getGold())
console.log('Total gold Player 2: ' + p2.getGold())

module.exports = Player;
