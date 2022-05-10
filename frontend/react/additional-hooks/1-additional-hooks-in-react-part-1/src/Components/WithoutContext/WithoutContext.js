import { useState } from 'react';
import Toolbar from './Toolbar';
import WeatherReaction from './WeatherReaction';

const WithoutContext = () =>{
    const [theme, setTheme] = useState('summer');
    return(
        <>
            <Toolbar theme={theme} setTheme={setTheme}/>
            <WeatherReaction theme={theme}/>
        </>
    )
}

export default WithoutContext;