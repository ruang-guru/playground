import styles from './Button.module.scss';
import { ThemeContext } from './WithContext';
import { useContext } from 'react';

const Button = () =>{
    const theme = useContext(ThemeContext).theme;
    const setTheme = useContext(ThemeContext).setTheme;

    const renderTheme = (theme) =>{
        if(theme == 'winter'){
            return {
                styles: styles['winter'],
                buttonText: 'Change To Summer'
            }
        }
        return {
            styles: styles['summer'],
            buttonText: 'Change To Winter'
        }
    }

    const changeTheme = (theme) =>{
        if(theme == "summer"){
            setTheme("winter")
        }else{
            setTheme("summer")
        }
    }


    return(
        <button onClick={()=>{changeTheme(theme)}} className={`${renderTheme(theme).styles} ${styles['button']}`}>
            {renderTheme(theme).buttonText} 
        </button>
    )
}

export default Button;