import styles from './Paragraph.module.scss';

const Paragraph = ({text}) =>{
    return(
        <div className={styles['paragraph-container']}>
            <p>
                {text}
            </p>
        </div>
    )
}

export default Paragraph;