import React from 'react';
import logo from './logo.png';
import './App.css';

const App = () => {
  return (
    <div className="App">
      <div className="div1">        <img src={logo} className="App-logo" alt="logo" /></div>
      <div className="div2"> 
      <form>
  <label>
    Search for a book by its ISBN: 
     <input type="text" name="isbn" />
  </label>
  <input type="submit" value="Submit" />
</form>
</div>
      <div className="div3">
        <h1>Owned Books</h1>
      </div>
      <div className="div4">
        <h1>Wanted Books</h1>
        </div>  
    </div>
  );
}

export default App;
