function App() {
  return <Profile />
}

function Profile() {
  // buat state name dan age
  const [name, setName] = React.useState('John Doe');
  const [age, setAge] = React.useState(0);

  return (
    <div>
      {/* render state name: {name} dan age: {age}  */}
      <p>Name: {name}</p>
      <p>Age: {age}</p>

      {/* render button untuk mengubah state name */}
      <button onClick={() => setName('Jane Doe')}>Change name</button>
      <button onClick={() => setAge(age + 1)}>Change age</button>
    </div>
  );
}