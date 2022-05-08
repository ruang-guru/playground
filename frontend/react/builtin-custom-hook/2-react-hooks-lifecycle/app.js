function App() {
  return <RandomBackground />
}

function RandomBackground() {
  const [color, setColor] = React.useState("#fff")
  React.useEffect(() => {
    const interval = setInterval(() => {
      setColor(generateColor())
    }, 1000)
    return () => clearInterval(interval)
  }, [])
  return <div style={{ backgroundColor: color, width: "100vw", height: "100vh" }}>
    <h1>Random Background</h1>
  </div>
}

function generateColor() {
  return `#${Math.floor(Math.random() * 16777215).toString(16)}`
}