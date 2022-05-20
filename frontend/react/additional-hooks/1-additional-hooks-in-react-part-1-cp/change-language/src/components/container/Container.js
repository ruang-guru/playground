import styles from './Container.module.scss';
import Dropdown from '../dropdown/dropdown';
const Container = ({children}) =>{
    return(
        <div className={styles['container']}>
            <Dropdown/>
            {children}
        </div>
    )
}                                                                                           

export default Container;