import { useMemo, useState, useEffect} from "react";

const PracticeUseMemo = () =>{
    const [number, setNumber] = useState(0);
    const [multiplier, setMultiplier] = useState(0);
    const [dark, setDark] = useState(false);
    const  doubleNumber = useMemo(()=>{
        // TODO: answer here
    );
    const themeStyles = useMemo(()=>{
        // TODO: answer here
    )
    useEffect(()=>{
        console.log('Theme Changed')
    }, [themeStyles])
    return(
        <>
            <input data-testid='input-one' type='number' value={number} onChange={e=> setNumber(parseInt(e.target.value))}/>
            <input data-testid='input-two' type='number' value={multiplier} onChange={e=> setMultiplier(parseInt(e.target.value))}/>
            <button data-testid='change-theme' onClick={()=> setDark(prevDark => !prevDark)}>Change Theme</button>
            <div data-testid='answer' style={themeStyles}>{doubleNumber}</div>
        
        </>
    )
}

const slowMultiplyFunction = (num, multiplier) =>{
    console.log('Calling Slow Function');

    for(let i=0; i< 1000000000;i++){}
    // TODO: answer here
}

export default PracticeUseMemo;