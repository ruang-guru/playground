import { useMemo, useState, useEffect} from "react";

const LearnUseMemo = () =>{
    const [number, setNumber] = useState(0);
    const [dark, setDark] = useState(false);
    const  doubleNumber = useMemo(()=>{
        return slowFunction(number)
    }, [number]);
    const themeStyles = useMemo(()=>{
        return {
            backgroundColor: dark ? 'black' : 'white',
            color: dark ? 'white' : 'black'
        }
    }, [dark])
    useEffect(()=>{
        console.log('Theme Changed')
    }, [themeStyles])
    return(
        <>
            <input type='number' value={number} onChange={e=> setNumber(parseInt(e.target.value))}/>
            <button onClick={()=> setDark(prevDark => !prevDark)}>Change Theme</button>
            <div style={themeStyles}>{doubleNumber}</div>
        
        </>
    )
}

const slowFunction = (num) =>{
    console.log('Calling Slow Function');
    for(let i=0; i< 1000000000;i++){}
    return num * 2;
}

export default LearnUseMemo;