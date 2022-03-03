import React from 'react';
import { makeStyles } from "@material-ui/core/styles"
import Years from '../components/pages/Years';
import Home from '../components/pages/Home';
import Topics from '../components/pages/Topics';
import TopicQuestions from '../components/pages/TopicQuestions';
import Login from '../components/Login';
import Signup from '../components/SignUp';

import {
  BrowserRouter as Router,
  Switch, Route, Link
} from "react-router-dom";

import {
  Drawer, List, ListItem,
  ListItemIcon, ListItemText,
  Container, Typography,
} from "@material-ui/core";

import HomeIcon from "@material-ui/icons/Home";
import DateRangeIcon from '@material-ui/icons/DateRange';
import TopicIcon from '@material-ui/icons/LibraryBooks';

const useStyles = makeStyles((theme) => ({
  drawerPaper: { width: 'inherit' },
  link: {
    textDecoration: 'none',
    color: theme.palette.text.primary
  }
}))

function App() {
  const classes = useStyles();
  return (
    <Router>
      <div style={{ display: 'flex' }}>
        <Drawer
          style={{ width: '220px' }}
          variant="persistent"
          anchor="left"
          open={true}
          classes={{ paper: classes.drawerPaper }}
        >
          <List>
          <Link to="/login" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <HomeIcon />
                </ListItemIcon>
                <ListItemText primary={"Login"} />
              </ListItem>
            </Link>
            <Link to="/" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <HomeIcon />
                </ListItemIcon>
                <ListItemText primary={"Home"} />
              </ListItem>
            </Link>
            <Link to="/years"  className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <DateRangeIcon />
                </ListItemIcon>
                <ListItemText primary={"Year wise Questions"} />
              </ListItem>
            </Link>
            <Link to="/topics" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <TopicIcon />
                </ListItemIcon>
                <ListItemText primary={"Topic wise Questions"} />
              </ListItem>
            </Link>
          </List>
        </Drawer>
        <Switch>
        { <Route exact path="/">
            <Home/>
          </Route> }
          <Route exact path="/years">
            <Years/>
          </Route>
          <Route exact path="/topics">
            <Topics/>
          </Route>
          <Route exact path="/topicquestions">
            <TopicQuestions/>
          </Route>
          <Route exact path="/login">
            <Login/>
          </Route>
          <Route exact path="/signup">
            <Signup/>
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
