import { useContext } from 'react';
import { ThemeContext } from './WithContext';

const WeatherReaction = () =>{
    const theme = useContext(ThemeContext).theme;
    const renderReaction = (theme) =>{
        if(theme == 'winter'){
            return `Brr its cold!`
        }else{
            return 'Hauhh its hot!'
        }
    }
    return(
        <p>
            {renderReaction(theme)}
        </p>
    )
}

export default WeatherReaction;