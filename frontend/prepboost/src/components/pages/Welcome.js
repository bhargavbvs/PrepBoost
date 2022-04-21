import React from 'react';
// import '../../App.css';
import {
  useHistory
} from "react-router-dom";

const Welcome = () => {
  
    const history = useHistory();
    const id = localStorage.getItem('userid');
    const username = localStorage.getItem('username');
    console.log(id,username)
    
  
    return (
      <div>    
        <h1 className='home'>Hello {username} <br/>Welcome to PREPBOOST. <br/> Let's Solve!!</h1>
      </div>
    );
  }


export default Welcome;
