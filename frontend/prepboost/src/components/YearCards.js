import React from 'react';
import './Cards.css';
import CardItem from './CardItem';

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
              path='/topicquestions'
            />
            <CardItem
              src='images/img_2020.png'
              text='2020'
              path='/topicquestions'
            />
            <CardItem
              src='images/img_2019.png'
              text='2019'
              path='/topicquestions'
            />
          </ul>
          
          <ul className='cards__items'>
            <CardItem
              src='images/img_2018.png'
              text='2018'
              path='/topicquestions'
            />
            <CardItem
              src='images/img_2017.png'
              text='2017'
              path='/topicquestions'
            />
            <CardItem
              src='images/img_2016.png'
              text='2016'
              path='/topicquestions'
            /><CardItem
              src='images/img_2015.png'
              text='2015'
              path='/topicquestions'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/img_2014.png'
              text='2014'
              path='/topicquestions'
            />
            <CardItem
              src='images/img_2013.png'
              text='2013'
              path='/topicquestions'
            />
            <CardItem
              src='images/img_2012.png'
              text='2012'
              path='/topicquestions'
            />
            <CardItem
              src='images/img_2011.png'
              text='2011'
              path='/topicquestions'
            />
            </ul>
        </div>
      </div>
    </div>
  );
}

export default YearCards;
