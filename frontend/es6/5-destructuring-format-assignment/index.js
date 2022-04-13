/** Format nama user menjadi "John - john@example.com - Unemployed",
 *  dan akses data user dengan teknik destrucuring assignment
 *  gunakan default value ketika atribut tidak ditemukan
 */

const format = (user) => {
  //beginanswer
  const { name, email, job = "Unemployed" } = user;
  return `${name} - ${email} - ${job}`;
  //endanswer
};

console.log(format({ name: "John", email: "john@example.com" }));

module.exports = format