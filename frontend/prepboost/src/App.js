import React from 'react';
import Navbar from './components/Navbar';
import Sidebar from './components/Sidebar';
import './App.css';
import Home from './components/pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Years from './components/pages/Years';
import Topics from './components/pages/Topics';
import TopicQuestions from './components/pages/TopicQuestions';
import Login from './components/Login';
import Signup from './components/SignUp';

function App() {
  return (
    <>
    <Sidebar />
    {/* <Login /> */}
      {/* <Router>
        <Switch>
          <Route path='/' exact component={Login} />
          <Route path='/home' exact component={Home} />
          <Route path='/signup' component={Signup} />
          <Route path='/years' component={Years} />
          <Route path='/topics' component={Topics} />
          <Route path='/topicquestions' component={TopicQuestions} />
        </Switch>
      </Router> */}
    </>
  );
}

export default App;
