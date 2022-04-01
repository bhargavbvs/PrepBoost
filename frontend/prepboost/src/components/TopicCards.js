import React from 'react';
import './Cards.css';
import CardItem from './CardItem';

function GetQuestions(topicname) {
  // Simple POST request with a JSON body using fetch
  const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title: topicname })
  };
  fetch('https://reqres.in/api/posts/questionsbytopic', requestOptions)
      .then(response => response.json())
      .then(data => this.setState({ postId: data.id }));
}

function TopicCards() {
  return (
    <div className='cards'>
      <h1>Check out the questions of different topics here</h1>
      <div className='cards__container'>
        <div className='cards__wrapper'>
          <ul className='cards__items'>
            <CardItem
              src='images/eco_min.png'
              text='Economy'
              path='/economyquestions'
            />
            <CardItem
              src='images/agri_min.png'
              text='Agriculture'
              path='/topicquestions'
            />
            <CardItem
              src='images/env_min.png'
              text='Environment'
              path='/topicquestions'
            />
          </ul>
          
          <ul className='cards__items'>
            <CardItem
              src='images/pol_min.png'
              text='Pol and Gov'
              path='/topicquestions'
            />
            <CardItem
              src='images/ir_min.png'
              text='IR'
              path='/topicquestions'
            />
            <CardItem
              src='images/sci_min.png'
              text='S and T'
              path='/topicquestions'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/mod_min.png'
              text='Modern'
              path='/topicquestions'
            />
            <CardItem
              src='images/anc_min.png'
              text='Ancient'
              path='/topicquestions'
            />
            <CardItem
              src='images/med_min.png'
              text='Medieval'
              path='/topicquestions'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/cul_min.png'
              text='Culture'
              path='/topicquestions'
            />
            <CardItem
              src='images/indgeo_min.png'
              text='Indian Geog'
              path='/topicquestions'
            />
            <CardItem
              src='images/geo_min.png'
              text='World Geog'
              path='/topicquestions'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/maps_min.png'
              text='Mapping'
              path='/topicquestions'
            />
            <CardItem
              src='images/curr_min.png'
              text='Current'
              path='/topicquestions'
            />
            <CardItem
              src='images/report_min.jpeg'
              text='Reports'
              path='/topicquestions'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/schemes_min.png'
              text='Schemes'
              path='/topicquestions'
            />
            <CardItem
              src='images/postin_min.png'
              text='Post Indep'
              path='/topicquestions'
            />
            <CardItem
              src='images/misc_min.png'
              text='Miscellaneous'
              path='/topicquestions'
            />
            </ul>
        </div>
      </div>
    </div>
  );
}

export default TopicCards;
