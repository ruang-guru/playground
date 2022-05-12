import styles from './Header.module.scss';
import { ThemeContext } from '../../App';
import { useContext } from 'react';

const Header = ({text}) =>{
    const theme = useContext(ThemeContext).theme;

    const renderTheme = (theme) =>{
        //mengubah warna dari Header berdasarkan theme
        // TODO: answer here
    }

    return(
        <h1 data-testid="header" className={`${styles['header']} ${renderTheme(theme)}`}>
            {text}
        </h1>
    )
}

export default Header;