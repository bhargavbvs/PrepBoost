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
import Welcome from '../components/pages/Welcome';

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
        // let result = {
        //     "ID": 10,
        //     "Username": "bhariojiojw",
        //     "Mobile": "+35278900990",
        //     "Email": "bvshbsdfds@gmailcom",
        //     "Password": "mypasios",
        //     "Paid": 0,
        //     "Search_left": 1,
        //     "Session_id": "akdhdfadsdfsddddsdffdfsafffk",
        //     "Created_at": "2022-04-19T21:42:52.532898-04:00",
        //     "Updated_at": "2022-04-19T21:42:52.532898-04:00",
        //     "Token": ""
        // }
        const values = []
        Object.keys(result).forEach(key => values.push({name: key, value: result[key]}))
        console.warn("values", values[1])

        let id = values[0].value
        let username = values[1].value
        
        return (
            <div>
            <Welcome id={id} username={username}/>
            </div>
        );
        
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