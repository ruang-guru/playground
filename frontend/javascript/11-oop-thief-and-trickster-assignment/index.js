const Player = require("../11-oop-steal-gold-cp/index");

class Thief extends Player {
  constructor(name) {
    super(name);
    this.job = "Thief";
  }
  
  robbedBlind() {
    // TODO: answer here
  }

  steal(player) {
    // TODO: answer here
  }
}

class Trickster extends Player {
  constructor(name) {
    super(name);
    this.distractionPurseChance = 0.25;
    this.job = "Trickster";
  }

  setDistractionPurseChance(chance) {
    // TODO: answer here
  }

  getDistractionPurseChance() {
    // TODO: answer here
  }

  distractionPurse(player) {
    const rng = this.randomizer();
    // TODO: answer here
  }

  steal(player) {
    // TODO: answer here
  }
}

// Dilarang menghapus code dibawah ini!!!
module.exports = {
  Thief,
  Trickster,
};
