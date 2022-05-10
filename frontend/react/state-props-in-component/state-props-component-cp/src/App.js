import React from 'react';

export const Item = (props) => {
  // TODO: answer here
  const [item, setItem] = React.useState(0);

  return(
    <div className='each-box'>
      <img src={image} alt="each-item"/>
      <p>{title}</p>
      <div>
        <button
        aria-label={`minus-button-${id}`}
        onClick={()=>{
        // TODO: answer here
        }}
        >-</button>
        <input type="text" aria-label={`item-${id}`} className={`item-${id}`} value={item} disabled></input>
        <button
        aria-label={`add-button-${id}`}
        onClick={()=>{
          // TODO: answer here
        }}
        >+</button>
      </div>
    </div>
  )
}


function App() {
  //Add state for total 
  // TODO: answer here
 
  const dataDummy = [
    {
      id    : '1',
      image : "https://picsum.photos/id/0/150/",
      title : "Laptop",
    },
    {
      id    : '2',
      image : "https://picsum.photos/id/1010/150/",
      title : "Buku",
    },
    {
      id    : '3',
      image : "https://picsum.photos/id/103/150/",
      title : "Sepatu",
    }
  ];

  const showAlert = (isShow) => {
    if(isShow){
      alert("Ups, udah kelebihan yaa");
    }
  }

  return (
    <div className="state-props-2" aria-label="AppRoot">
      <h3>Keranjang Belanja</h3>
      <div className='box-container'>
          {dataDummy.map((element, index) => (
            // TODO: answer here
          ))}
      </div>
      <div className='end-section'>
          <h4>Kamu udah masukin&nbsp;
            <span className='total-item' aria-label="total-item">{total}
            </span>&nbsp;
          barang</h4>
      </div>
    </div>
  );
}

export default App;
