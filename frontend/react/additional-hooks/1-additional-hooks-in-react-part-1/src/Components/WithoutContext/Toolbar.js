import ThemeButton from "./ThemeButton"

const Toolbar = ({theme, setTheme}) =>{
    return(
        <ThemeButton theme={theme} setTheme={setTheme}/>
    )
}

export default Toolbar;