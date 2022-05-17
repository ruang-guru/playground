const WeatherReaction = ({theme}) =>{
    const renderReaction = (theme) =>{
        if(theme == 'winter'){
            return `Brr its cold!`
        }else{
            return 'Hauhh its hot!'
        }
    }
    return(
        <p>
            {renderReaction(theme)}
        </p>
    )
}

export default WeatherReaction;