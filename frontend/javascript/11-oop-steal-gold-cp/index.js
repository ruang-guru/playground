class Player {
  constructor(name) {
    this.name = name;
    this.gold = 50;
    this.stealChance = 0.25;
    this.hasBeenRobbed = false;
  }


  getGold() {
  // beginanswer
    return this.gold;
  //endanswer
  }

  setGold(gold) {
  // beginanswer
    this.gold = gold;
  //endanswer
  }

  getStealChance() {
  // beginanswer
    return this.stealChance;
  //endanswer
  }

  setStealChance(chance) {
  // beginanswer
    this.stealChance = chance;
  //endanswer
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
    // beginanswer
    if (player.gold < 5) {
      return "Lawan terlalu miskin";
    }

    const rng = this.randomizer();

    if (rng <= this.stealChance) {
      this.setGold(this.getGold() + 5);
      player.setGold(player.getGold() - 5);
      player.setHasBeenRobbed(true);

      return "Berhasil mencuri 5 gold";
    }

    return "Gagal mencuri, coba lain kali";
    //endanswer
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
