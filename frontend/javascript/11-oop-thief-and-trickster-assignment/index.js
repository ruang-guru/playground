const Player = require("../11-oop-steal-gold-cp/index");

class Thief extends Player {
  constructor(name) {
    super(name);
    this.job = "Thief";
  }
  
  robbedBlind() {
    //beginanswer
    if (this.getGold() < 10) {
      return "Aku terlalu miskin";
    }

    this.setGold(this.getGold() - 10);
    this.setStealChance(0.75);
    //endanswer
  }

  steal(player) {
    //beginanswer
    if (player.gold < 5) {
      return "Lawan terlalu miskin";
    }

    const rng = this.stealChance;

    if (rng <= this.stealChance) {
      this.setGold(this.getGold() + 5);
      player.setGold(player.getGold() - 5);

      if (player.job && player.job === "Trickster") {
        return player.distractionPurse(this);
      }

      return "Berhasil mencuri 5 gold";
    }

    return "Gagal mencuri, coba lain kali";
    //endanswer
  }
}

class Trickster extends Player {
  constructor(name) {
    super(name);
    this.distractionPurseChance = 0.25;
    this.job = "Trickster";
  }

  setDistractionPurseChance(chance) {
    //beginanswer
    this.distractionPurseChance = chance;
    //endanswer
  }

  getDistractionPurseChance() {
    //beginanswer
    return this.distractionPurseChance;
    //endanswer
  }

  distractionPurse(player) {
    const rng = this.randomizer();
    //beginanswer
    if (rng <= this.distractionPurseChance) {
      if (player.getGold() < 10) {
        player.setGold(0);
        this.setGold(this.getGold() + player.getGold());

        return "Berhasil mencuri balik semua uang lawan";
      }

      player.setGold(player.getGold() - 10);
      this.setGold(this.getGold() + 10);

      return "Berhasil mencuri balik 10 gold";
    }

    return "Gagal mencuri balik";
    //endanswer
  }

  steal(player) {
    //beginanswer
    if (player.gold < 5) {
      return "Lawan terlalu miskin";
    }

    const rng = this.stealChance;

    if (rng <= this.stealChance) {
      this.setGold(this.getGold() + 5);
      player.setGold(player.getGold() - 5);

      if (player.job && player.job === "Trickster") {
        return player.distractionPurse(this);
      }

      return "Berhasil mencuri 5 gold";
    }

    return "Gagal mencuri, coba lain kali";
    //endanswer
  }
}

// Dilarang menghapus code dibawah ini!!!
module.exports = {
  Thief,
  Trickster,
};
