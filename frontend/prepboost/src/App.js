import React from 'react';
import Navbar from './components/Navbar';
import Sidebar from './components/Sidebar';
import './App.css';
import Home from './components/pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Years from './components/pages/Years';
import Topics from './components/pages/Topics';
import TopicQuestions from './components/pages/TopicQuestions';
import YearQuestions from './components/pages/YearQuestions';
import Login from './components/Login';
import Signup from './components/SignUp';

function App() {
  localStorage.setItem('userid', null);
  localStorage.setItem('username', null);
  return (
    <div className='sidebar' id = "container">
    {/* <Navbar /> */}
    <Sidebar />
      {/* <Router>
        <Switch>
          <Route path='/' exact component={Login} />
          <Route path='/home' exact component={Home} />
          <Route path='/signup' component={Signup} />
          <Route path='/years' component={Years} />
          <Route path='/topics' component={Topics} />
          <Route path='/topicquestions' component={TopicQuestions} />
          <Route path='/yearquestions' component={YearQuestions} />
        </Switch>
      </Router> */}
    </div>
  );
}

export default App;
