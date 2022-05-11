import styles from './Container.module.scss';
import { ThemeContext } from '../../App';
import { useContext } from 'react';

const Container = ({children}) =>{
    const theme = useContext(ThemeContext).theme;

    const renderTheme = (theme) =>{
        //mengubah background dari container
        // TODO: answer here
    }

    return(
        <div data-testid="container" className={`${styles['container']} ${renderTheme(theme)}`}>
            {children}
        </div>
    )
}

export default Container;