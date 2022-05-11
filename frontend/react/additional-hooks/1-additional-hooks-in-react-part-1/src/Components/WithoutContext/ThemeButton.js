import Button from './Button';

const ThemeButton = ({theme, setTheme}) =>{
    return(
        <Button setTheme={setTheme} theme={theme}>Theme Button</Button>
    )
}

export default ThemeButton;