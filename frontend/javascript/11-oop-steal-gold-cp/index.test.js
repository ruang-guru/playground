const Player = require("./index");

describe("Player class test", () => {
  it("should return string of failed if stealing failed", () => {
    const p1 = new Player("Hehe");
    const p2 = new Player("Name");

    p1.setStealChance(0);

    const message = p1.steal(p2);

    expect(message).toBe("Gagal mencuri, coba lain kali");
  });

  it("should return string of enemy gold information if enemy gold is less than 5 gold", () => {
    const p1 = new Player("Hehe");
    const p2 = new Player("Name");

    p2.setGold(4);

    const message = p1.steal(p2);

    expect(message).toBe("Lawan terlalu miskin");
  });

  it("should return string of success if stealing success", () => {
    const p1 = new Player("Hehe");
    const p2 = new Player("Name");

    p1.setStealChance(1);

    const message = p1.steal(p2);

    expect(message).toBe("Berhasil mencuri 5 gold");
  });

  it("should increase self gold if stealing succeded", () => {
    const p1 = new Player("Hehe");
    const p2 = new Player("Name");

    p1.setStealChance(1);

    p1.steal(p2);

    const selfGold = p1.getGold();

    expect(selfGold).toBe(55);
  });

  it("should decrease enemy gold if stealing succeded", () => {
    const p1 = new Player("Hehe");
    const p2 = new Player("Name");

    p1.setStealChance(1);

    p1.steal(p2);

    const selfGold = p2.getGold();

    expect(selfGold).toBe(45);
  });
});
