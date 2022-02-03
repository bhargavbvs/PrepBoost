import React from 'react';
import './Cards.css';
import CardItem from './CardItem';

function TopicCards() {
  return (
    <div className='cards'>
      <h1>Check out the questions of different year here</h1>
      <div className='cards__container'>
        <div className='cards__wrapper'>
          <ul className='cards__items'>
            <CardItem
              src='images/img_2020.png'
              text='Economy'
              path='/years'
            />
            <CardItem
              src='images/img_2020.png'
              text='Agriculture'
              path='/years'
            />
            <CardItem
              src='images/img_2019.png'
              text='Environment'
              path='/years'
            />
          </ul>
          
          <ul className='cards__items'>
            <CardItem
              src='images/img_2018.png'
              text='Pol and Gov'
              path='/topics'
            />
            <CardItem
              src='images/img_2017.png'
              text='IR'
              path='/sign-up'
            />
            <CardItem
              src='images/img_2016.png'
              text='S and T'
              path='/sign-up'
            /><CardItem
              src='images/img_2015.png'
              text='Modern'
              path='/topics'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/img_2014.png'
              text='Ancient'
              path='/topics'
            />
            <CardItem
              src='images/img_2013.png'
              text='Medieval'
              path='/topics'
            />
            <CardItem
              src='images/img_2012.png'
              text='Culture'
              path='/topics'
            />
            <CardItem
              src='images/img_2011.png'
              text='Indian Geog'
              path='/topics'
            />
            </ul>
        </div>
      </div>
    </div>
  );
}

export default TopicCards;
