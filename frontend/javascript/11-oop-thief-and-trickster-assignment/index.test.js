const { Thief, Trickster } = require("./index");

describe("Thief class test", () => {
  it("should increase stealChance to 75% when skill robbedBlind is used", () => {
    const p1 = new Thief("hehe");

    p1.robbedBlind();

    expect(p1.getStealChance()).toBe(0.75);
  });

  it("should return failed message if player gold less than 10", () => {
    const p1 = new Thief("hehe");

    p1.setGold(5);
    const message = p1.robbedBlind();

    expect(message).toBe("Aku terlalu miskin");
  });
});

describe("Trickster class test", () => {
  it("should run distraction purse and return failed message if failed", () => {
    const p1 = new Thief("hehe");
    const p2 = new Trickster("kek");

    p1.setStealChance(1);
    p2.setDistractionPurseChance(0);
    const message = p1.steal(p2);

    expect(message).toBe("Gagal mencuri balik");
  });

  it("should run distraction purse and return success message if success", () => {
    const p1 = new Thief("hehe");
    const p2 = new Trickster("kek");

    p1.setStealChance(1);
    p2.setDistractionPurseChance(1);
    const message = p1.steal(p2);

    expect(message).toBe("Berhasil mencuri balik 10 gold");
  });

  it("should run distraction purse and increase self gold by 5 when success", () => {
    const p1 = new Thief("hehe");
    const p2 = new Trickster("kek");

    p1.setStealChance(1);
    p2.setDistractionPurseChance(1);
    p1.steal(p2);

    const gold = p2.getGold();

    expect(gold).toBe(55);
  });

  it("should run distraction purse and decrease thief gold by 5 when success", () => {
    const p1 = new Thief("hehe");
    const p2 = new Trickster("kek");

    p1.setStealChance(1);
    p2.setDistractionPurseChance(1);
    p1.steal(p2);

    const gold = p1.getGold();

    expect(gold).toBe(45);
  });

  it("should run distraction purse and return success message if success but thief has not enough gold", () => {
    const p1 = new Thief("hehe");
    const p2 = new Trickster("kek");

    p1.setStealChance(1);
    p1.setGold(4);
    p2.setDistractionPurseChance(1);
    const message = p1.steal(p2);

    expect(message).toBe("Berhasil mencuri balik semua uang lawan");
  });

  it("should run distraction purse and return failed message if failed", () => {
    const p1 = new Thief("hehe");
    const p2 = new Trickster("kek");

    p1.setStealChance(1);
    p1.setGold(4);
    p2.setDistractionPurseChance(0);
    const message = p1.steal(p2);

    expect(message).toBe("Gagal mencuri balik");
  });
});
