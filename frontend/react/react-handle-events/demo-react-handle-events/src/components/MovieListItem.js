import "./MovieListItem.css";
const MovieListItem = (props) => {
  const { movieItem } = props;
  return (
    <div className="list-item-container">
      <img src={`https://image.tmdb.org/t/p/w300${movieItem.backdrop_path}`} />
      <div className="movie-content-container">
        <h3>{movieItem.original_title}</h3>
        <p>{movieItem.overview}</p>
      </div>
    </div>
  );
};

export default MovieListItem;
