import { Button } from 'react-bootstrap';

import '../components/PodcastListItem.css';

const PodcastListItem = ({ id, podcastItem, setFormModalType, setShowFormModal, setPodcastId }) => {
  const onClickUpdateBtn = () => {
    // TODO: answer here
  }
  return (
    <div aria-label='podcast-list-item' className="list-item-container">
      <img alt={'podcast-img'} width={300} src={podcastItem.imageUrl} />
      <div className="podcast-content-container">
        <h3>{podcastItem.title}</h3>
        <div className='podcast-subtitle-container'>
          <h6>Published by : {podcastItem.publisher}</h6>
          <h6>Durasi : {podcastItem.duration} menit</h6>
        </div>
        <p>{podcastItem.summary}</p>
        <Button onClick={onClickUpdateBtn}>Update</Button>
      </div>
    </div>
  );
};

export default PodcastListItem;
