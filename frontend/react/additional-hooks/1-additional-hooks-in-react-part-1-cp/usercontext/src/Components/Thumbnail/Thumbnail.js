import styles from './Thumbnail.module.scss';

const Thumbnail = ({src, alt}) =>{
    return(
        <div className={styles['thumbnail-container']}>
            <img src={src} alt={alt}/>
        </div>
    )
}

export default Thumbnail;