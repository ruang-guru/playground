let setA = new Set(["Java", "Python", "Javascript", "C ++", "C#"]);

setA.add("Java"); //doesn't change
console.log(setA); 

setA.add("Ruby"); //adds
console.log(setA);

setA.delete("Java"); //deletes
console.log(setA);

setA.delete("Rust"); //doesn't change
console.log(setA);
