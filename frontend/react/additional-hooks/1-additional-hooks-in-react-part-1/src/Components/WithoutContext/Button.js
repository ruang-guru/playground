import styles from './Button.module.scss';

const Button = ({theme, children, setTheme}) =>{
    const renderTheme = (theme) =>{
        if(theme == 'winter'){
            return styles['winter'];
        }
        return styles['summer']
    }

    const changeTheme = (theme) =>{
        if(theme == "summer"){
            setTheme("winter")
        }else{
            setTheme("summer")
        }
    }

    return(
        <button onClick={()=>{changeTheme(theme)}} className={`${renderTheme(theme)} ${styles['button']}`}>
            {children}
        </button>
    )
}

export default Button;