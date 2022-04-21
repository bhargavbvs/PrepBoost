import React, { useState } from "react";
import {
  Grid,
  Paper,
  Avatar,
  TextField,
  Button,
  Typography,
  Link,
} from "@material-ui/core";
import {
  useHistory
} from "react-router-dom";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";


const Login = () => {
  const [Username,setName] = useState("")
  const [Password,setPassword] = useState("")
  const history = useHistory();


  const paperStyle = {
    padding: 20,
    height: "70vh",
    width: 300,
    margin: "20px auto",
  };
  const avatarStyle = { backgroundColor: "#1bbd7e" };
  const btnstyle = { margin: "8px 0" };
  

  async function Login()
  {
      let item ={Username,Password}
      console.warn("details",item)

      let response = await fetch("http://6f49-2600-8807-c0c0-d400-cc94-a9ef-a41-1466.ngrok.io/users/login/",{
          method: 'POST',
          body:JSON.stringify(item),
          headers:{
              "Content-Type": 'application/json',
              "Accept": 'application/json'
          }
      })
      if(response.status === 200)
      {
        response = await response.json()
        const values = []
        Object.keys(response).forEach(key => values.push({name: key, value: response[key]}))
        console.warn("values", values[1])
        let id = values[0].value
        let username = values[1].value
        localStorage.setItem('userid', id);
        localStorage.setItem('username', username);
        history.push("/welcome", {id: {id}, username: {username}});
      }
      
  }
  function Logout()
  {
    localStorage.setItem('userid', null);
    localStorage.setItem('username', null);
    history.push("/");
  }
  let userid = localStorage.getItem('userid')
  if(userid === 'null')
  {
    console.log("null",userid)
    return (
      <Grid className="login">
        <Paper elevation={10} style={paperStyle}>
          <Grid align="center">
            <Avatar style={avatarStyle}>
              <LockOutlinedIcon />
            </Avatar>
            <h2>Sign In</h2>
          </Grid>
          <TextField
            label="Username"
            placeholder="Enter username"
            fullWidth
            required
            onChange={(e) => {setName(e.target.value)}}
          />
          <TextField
            label="Password"
            placeholder="Enter password"
            type="password"
            fullWidth
            required
            onChange={(e) => {setPassword(e.target.value)}}
          />
          <FormControlLabel
            control={<Checkbox name="checkedB" color="primary" />}
            label="Remember me"
          />
          <Button
            type="submit"
            color="primary"
            variant="contained"
            style={btnstyle}
            fullWidth
            onClick={Login}
          >
            Sign in
          </Button>
          <Typography>
            <Link href="#">Forgot password ?</Link>
          </Typography>
          <Typography>
            {" "}
            Do you have an account ?<Link href="signup">Sign Up</Link>
          </Typography>
        </Paper>
      </Grid>
    );
  }
  else
  {
    console.log("notnull",userid)
    return (
      <Grid className="login">
        <Paper elevation={10} style={paperStyle}>
          <Grid align="center">
            <Avatar style={avatarStyle}>
              <LockOutlinedIcon />
            </Avatar>
            <h2>Logout</h2>
          </Grid>
          
          <Button
            type="submit"
            color="primary"
            variant="contained"
            style={btnstyle}
            fullWidth
            onClick={Logout}
          >
            Logout
          </Button>
          
        </Paper>
      </Grid>
    );
  }
};

export default Login;
