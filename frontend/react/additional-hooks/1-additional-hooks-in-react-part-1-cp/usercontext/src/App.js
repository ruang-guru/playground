import { createContext } from 'react';
import { useState } from 'react';

import './global.scss';
import Container from "./Components/Container/Container";
import Thumbnail from "./Components/Thumbnail/Thumbnail";
import catImage from './Assets/Img/cat.jpg';
import Paragraph from "./Components/Paragraph/Paragraph";
import text from "./Text/loremIpsum";
import Header from "./Components/Header/Header";
import ToggleButton from './Components/ToggleButton/ToggleButton';

export const ThemeContext = createContext();

const App = () =>{
    const [theme, setTheme] = useState("light");
    return(
        <ThemeContext.Provider value={{
            theme: theme,
            setTheme: setTheme
        }} >
            <Container>
                <ToggleButton/>
                <Thumbnail src={catImage} alt='cat'/>
                <Header text="Lorem Ipsum"/>
                <Paragraph text={text}/>
            </Container>
       
        </ThemeContext.Provider>
    )
}

export default App;