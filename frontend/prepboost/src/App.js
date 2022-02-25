import React from 'react';
import Navbar from './components/Navbar';
import Sidebar from './components/Sidebar';
import './App.css';
import Home from './components/pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Years from './components/pages/Years';
import Topics from './components/pages/Topics';
import TopicQuestions from './components/pages/TopicQuestions';

function App() {
  return (
    <>
    <Sidebar />
      {/* <Router>
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/years' component={Years} />
          <Route path='/topics' component={Topics} />
          <Route path='/topicquestions' component={TopicQuestions} />
        </Switch>
      </Router> */}
    </>
  );
}

export default App;
