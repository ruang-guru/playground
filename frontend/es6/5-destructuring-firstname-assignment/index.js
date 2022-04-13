/** Akses user name, namun dengan nama variabel yang berbeda dari atribut pada objek user
 */
const getUserFirstName = (user) => {
  //beginanswer
  const { name: firstName } = user;
  return firstName;
  //endanswer
};

console.log(getUserFirstName({ name: "John", email: "john@example.com" }));

module.exports = getUserFirstName