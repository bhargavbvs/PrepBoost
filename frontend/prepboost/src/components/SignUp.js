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
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";

const Signup = () => {
    const paperStyle = { padding: '30px 20px', width: 320, margin: "20px auto" }
    const btnstyle = { margin: "8px 0" };
    const headerStyle = { margin: 0 }
    const avatarStyle = { backgroundColor: '#1bbd7e' }
    const [Username,setName] = useState("")
    const [Email,setEmail] = useState("")
    const [Mobile,setPhonenumber] = useState("")
    const [Password,setPassword] = useState("") 
    const Paid = 0
    const Search_left = 1
    const Session_id = "akdhdfajdddfddsdfsddddsdffdfsafffk"

    async function signUp()
    {
        let item ={Username,Email,Mobile,Password,Paid,Search_left,Session_id}
        console.warn(item)

        let result = await fetch("http://711d-2600-8807-c0c4-300-e1dc-68bf-67dc-b13.ngrok.io/users/signup/",{
            method: 'POST',
            body:JSON.stringify(item),
            headers:{
                "Content-Type": 'application/json',
                "Accept": 'application/json'
            }
        })
        result = await result.json()
        console.warn("result", result)
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