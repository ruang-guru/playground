import { useState, useEffect } from 'react'
import axios from 'axios';

import MovieListItem from '../components/MovieListItem';
import './MoviePages.css';

const API_KEY = 'cd2ca552dd43d692af2e0defeb01893f'

const MoviePages = (props) => {
    const [movieList, setMovieList] = useState([]);
    const { userData } = props;

    const getMovieList = async() => {
        try {
            const MOVIE_LIST_API_URL = `https://api.themoviedb.org/4/list/1?api_key=${API_KEY}&language=${userData.userLanguage}`;
            const getMovieList = await axios.get(MOVIE_LIST_API_URL);
            if (getMovieList.status == 200) {
                const movieListResult = getMovieList.data.results;
                setMovieList(movieListResult);
            }
        } catch (err) {
            console.log("error fetch movie list data ", err);
        }
    }

    useEffect(() => {
        getMovieList();
    }, []);

    return (
        <div className='movie-pages-container'>
            <h1>MCU Movies</h1>
            {
                movieList.length > 0 && movieList.map(item => {
                    return (
                        <MovieListItem key={item.id} movieItem={item}/>
                    )
                })
            }
        </div>
    );
};

export default MoviePages;
