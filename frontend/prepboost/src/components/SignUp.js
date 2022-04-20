import React, {useState} from 'react'
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

const Signup = () => {
    const paperStyle = { padding: '30px 20px', width: 320, margin: "20px auto" }
    const btnstyle = { margin: "8px 0" };
    const avatarStyle = { backgroundColor: '#1bbd7e' }
    const [Username,setName] = useState("")
    const [Email,setEmail] = useState("")
    const [Mobile,setPhonenumber] = useState("")
    const [Password,setPassword] = useState("") 
    const Paid = 0
    const Search_left = 1
    const Session_id = (Math.random() + 1).toString(36).substring(20);
    const history = useHistory();

    async function signUp()
    {
        let item ={Username,Email,Mobile,Password,Paid,Search_left,Session_id}
        console.warn(item)

<<<<<<< Updated upstream
        let response = await fetch("http://6f49-2600-8807-c0c0-d400-cc94-a9ef-a41-1466.ngrok.io/users/signup/",{
=======
        let result = await fetch("http://localhost:9010/users/signup/",{
>>>>>>> Stashed changes
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

    return (
        <Grid className='signup'>
            <Paper elevation={20} style={paperStyle}>
            <Grid align="center">
          <Avatar style={avatarStyle}>
            <LockOutlinedIcon />
          </Avatar>
          <h2>Fill the details to Register</h2>
        </Grid>
        <TextField
          label="Username"
          placeholder="Enter username"
          fullWidth
          required
          onChange={(e) => {setName(e.target.value)}}
        />
        <TextField
          label="Email id"
          placeholder="Enter email id"
          type="email"
          fullWidth
          required
          onChange={(e) => {setEmail(e.target.value)}}
        />
        <TextField
          label="Phone Number"
          placeholder="Enter phone number"
          type="number"
          fullWidth
          required
          onChange={(e) => {setPhonenumber(e.target.value)}}
        />
        <TextField
          label="Password"
          placeholder="Enter password"
          type="password"
          fullWidth
          required
          onChange={(e) => {setPassword(e.target.value)}}
        />
        <Button
          type="submit"
          color="primary"
          variant="contained"
          style={btnstyle}
          fullWidth
          onClick={signUp}
        >
          Register
          </Button>
          <Typography>
          {" "}
          Already have an account ?<Link href="login">Login</Link>
        </Typography>
            </Paper>
        </Grid>
    )
}

export default Signup;