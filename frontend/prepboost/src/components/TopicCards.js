import React from 'react';
import './Cards.css';
import CardItem from './CardItem';

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
              path='/years'
            />
            <CardItem
              src='images/agri_min.png'
              text='Agriculture'
              path='/years'
            />
            <CardItem
              src='images/env_min.png'
              text='Environment'
              path='/years'
            />
          </ul>
          
          <ul className='cards__items'>
            <CardItem
              src='images/pol_min.png'
              text='Pol and Gov'
              path='/topics'
            />
            <CardItem
              src='images/ir_min.png'
              text='IR'
              path='/sign-up'
            />
            <CardItem
              src='images/sci_min.png'
              text='S and T'
              path='/sign-up'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/mod_min.png'
              text='Modern'
              path='/topics'
            />
            <CardItem
              src='images/anc_min.png'
              text='Ancient'
              path='/topics'
            />
            <CardItem
              src='images/med_min.png'
              text='Medieval'
              path='/topics'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/cul_min.png'
              text='Culture'
              path='/topics'
            />
            <CardItem
              src='images/indgeo_min.png'
              text='Indian Geog'
              path='/topics'
            />
            <CardItem
              src='images/geo_min.png'
              text='World Geog'
              path='/topics'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/maps_min.png'
              text='Mapping'
              path='/topics'
            />
            <CardItem
              src='images/curr_min.png'
              text='Current'
              path='/topics'
            />
            <CardItem
              src='images/report_min.jpeg'
              text='Reports'
              path='/topics'
            />
            </ul>
            
            <ul className='cards__items'>
            <CardItem
              src='images/schemes_min.png'
              text='Schemes'
              path='/topics'
            />
            <CardItem
              src='images/postin_min.png'
              text='Post Indep'
              path='/topics'
            />
            <CardItem
              src='images/misc_min.png'
              text='Miscellaneous'
              path='/topics'
            />
            </ul>
        </div>
      </div>
    </div>
  );
}

export default TopicCards;
