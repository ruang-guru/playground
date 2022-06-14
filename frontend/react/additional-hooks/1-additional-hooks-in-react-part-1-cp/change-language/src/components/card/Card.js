import styles from './Card.module.scss';
import { useContext } from 'react';
import { LanguageContext } from '../../App';

const Card = () =>{
    //buatlah variable untuk mengambil context dari text
    // TODO: answer here
    return(
        <div className={styles['card']}>
            <div className={styles['card-image']}>
            </div>
            <div className={styles['card-content']}>
                <div className={styles['card-title']}>
                    <p data-testid='headline'>{text.headline}</p>
                </div>
                <div className={styles['card-description']}>
                    <p>{text.description}
                    </p>
                </div>
                <div className={styles['card-data__group']}>
                    <div className={styles['card-data']}>
                        <p>10K+</p>
                        <p>{text.companies}</p>
                    </div>
                    <div className={styles['card-data']}>
                        <p>314</p>
                        <p>{text.templates}</p>
                    </div>
                    <div className={styles['card-data']}>
                        <p>12M+</p>
                        <p>{text.queries}</p>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Card;