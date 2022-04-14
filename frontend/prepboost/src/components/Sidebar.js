import React from 'react';
import { makeStyles } from "@material-ui/core/styles"
import Years from '../components/pages/Years';
import Home from '../components/pages/Home';
import Topics from '../components/pages/Topics';
import TopicQuestions from '../components/pages/TopicQuestions';
import EconomyQuestions from '../components/pages/Economyquestions';
import Login from '../components/Login';
import Signup from '../components/SignUp';
import Bookmarks from '../components/pages/Bookmarks';
import YearQuestions from '../components/pages/YearQuestions';
import Contact from '../components/pages/contact';
import Agriculturequestions from '../components/pages/Agriculturequestions';

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
import BookmarksIcon from '@material-ui/icons/Bookmarks';

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
            <Link to="/bookmarks" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <BookmarksIcon />
                </ListItemIcon>
                <ListItemText primary={"Bookmarks"} />
              </ListItem>
            </Link>
            <Link to="/contact" className={classes.link}>
              <ListItem button>
                <ListItemIcon>
                  <BookmarksIcon />
                </ListItemIcon>
                <ListItemText primary={"Contact"} />
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
          <Route exact path="/economyquestions">
            <TopicQuestions data="economy"/>
          </Route>
          <Route exact path="/agriculturequestions">
            <TopicQuestions data="agriculture"/>
          </Route>
          <Route exact path="/environmentquestions">
            <TopicQuestions data="environment"/>
          </Route>
          <Route exact path="/politicsquestions">
            <TopicQuestions data="politics"/>
          </Route>
          <Route exact path="/irquestions">
            <TopicQuestions data="ir"/>
          </Route>
          <Route exact path="/sciencequestions">
          <TopicQuestions data="science"/>
          </Route>
          <Route exact path="/modernquestions">
          <TopicQuestions data="modern"/>
          </Route>
          <Route exact path="/ancientquestions">
          <TopicQuestions data="ancient"/>
          </Route>
          <Route exact path="/medivalquestions">
          <TopicQuestions data="medival"/>
          </Route>
          <Route exact path="/culturequestions">
          <TopicQuestions data="culture"/>
          </Route>
          <Route exact path="/indiangeoquestions">
          <TopicQuestions data="indiangeo"/>
          </Route>
          <Route exact path="/worldgeoquestions">
          <TopicQuestions data="worldgeo"/>
          </Route>
          <Route exact path="/mappingquestions">
          <TopicQuestions data="mapping"/>
          </Route>
          <Route exact path="/currentquestions">
          <TopicQuestions data="current"/>
          </Route>
          <Route exact path="/reportsquestions">
          <TopicQuestions data="reports"/>
          </Route>
          <Route exact path="/schemesquestions">
          <TopicQuestions data="schemes"/>
          </Route>
          <Route exact path="/postindependencequestions">
          <TopicQuestions data="postindependence"/>
          </Route>
          <Route exact path="/miscellaneousquestions">
          <TopicQuestions data="miscellaneous"/>
          </Route>
          <Route exact path="/yearquestions">
            <YearQuestions/>
          </Route>
          <Route exact path="/login">
            <Login/>
          </Route>
          <Route exact path="/signup">
            <Signup/>
          </Route>
          <Route exact path="/bookmarks">
            <Bookmarks/>
          </Route>
          <Route exact path="/contact">
            <Contact/>
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
