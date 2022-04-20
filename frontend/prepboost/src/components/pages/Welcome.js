import React from 'react';
// import '../../App.css';
import {
  useHistory
} from "react-router-dom";

const Welcome = () => {
  
    const history = useHistory();
    const id = history.location.state.id;
    const username = history.location.state.username;
    console.log(id,username)
    
  
    return (
      <div>    
        <h1 className='home'>Hello  <br/>Welcome to PREPBOOST. <br/> Let's Solve!!</h1>
      </div>
    );
  }


export default Welcome;
