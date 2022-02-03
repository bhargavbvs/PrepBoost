import styled from 'styled-components'
import './Buttons.css'

function clickme(){
    alert('You clicked me!');
}

function Buttons() {
    return (
    <div>
    
      <a href="https://www.youtube.com" target="_blank">
      <img src="/images/Peaky blinders.jpg" alt=""/>
      </a>

    </div>
    );
  }
  
  export default Buttons;