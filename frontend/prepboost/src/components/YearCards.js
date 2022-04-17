import React from 'react';
import './Cards.css';
import CardItem from './CardItem';

function GetQuestions(yearname) {
  // Simple POST request with a JSON body using fetch
  const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title: yearname })
  };
  fetch('https://reqres.in/api/posts/questionsbyyear', requestOptions)
      .then(response => response.json())
      .then(data => this.setState({ postId: data.id }));
}

function YearCards() {
  return (
    <div className='cards'>
      <h1>Check out the questions of different years here</h1>
      <div className='cards__container'>
        <div className='cards__wrapper'>
          <ul className='cards__items'>
            <CardItem
              src='images/img_2021.png'
              text='2021'
              path='/2021questions'
            />
            <CardItem
              src='images/img_2020.png'
              text='2020'
              path='/2020questions'
            />
            <CardItem
              src='images/img_2019.png'
              text='2019'
              path='/2019questions'
            />
          </ul>
          
          <ul className='cards__items'>
            <CardItem
              src='images/img_2018.png'
              text='2018'
              path='/2018questions'
            />
            <CardItem
              src='images/img_2017.png'
              text='2017'
              path='/2017questions'
            />
            <CardItem
              src='images/img_2016.png'
              text='2016'
              path='/2016questions'
            /><CardItem
              src='images/img_2015.png'
              text='2015'
              path='/2015questions'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/img_2014.png'
              text='2014'
              path='/2014questions'
            />
            <CardItem
              src='images/img_2013.png'
              text='2013'
              path='/2013questions'
            />
            <CardItem
              src='images/img_2012.png'
              text='2012'
              path='/2012questions'
            />
            <CardItem
              src='images/img_2011.png'
              text='2011'
              path='/2011questions'
            />
            </ul>
        </div>
      </div>
    </div>
  );
}

export default YearCards;
