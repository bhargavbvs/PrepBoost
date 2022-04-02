import React, {useState} from 'react'
import { Grid, Paper, Avatar, Typography, TextField, Button } from '@material-ui/core'
import AddCircleOutlineOutlinedIcon from '@material-ui/icons/AddCircleOutlineOutlined';
import { Link } from 'react-router-dom'

const Signup = () => {
    const paperStyle = { padding: '30px 20px', width: 320, margin: "20px auto" }
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
                "Content-Type": 'appliaction/json',
                "Accept": 'application/json'
            }
        })
        result = await result.json()
        console.warn("result", result)
    }

    return (
        <Grid className='signup'>
            <Paper elevation={20} style={paperStyle}>
                <Grid align='center'>
                    <Avatar style={avatarStyle}>
                        <AddCircleOutlineOutlinedIcon />
                    </Avatar>
                    <h2 style={headerStyle}>Sign Up</h2>
                    <Typography variant='caption' gutterBottom>Please fill this form to create an account !</Typography>
                </Grid>
                <form >
                <p>
                    <label>Username</label><br/>
                    <input type="text" value={Username} onChange={(e)=> setName(e.target.value)} name="first_name" required />
                </p>
                <p>
                    <label>Email address</label><br/>
                    <input type="email" value={Email} onChange={(e)=> setEmail(e.target.value)} name="email" required />
                </p>
                <p>
                    <label>Phone Number</label><br/>
                    <input type="number" value={Mobile} onChange={(e)=> setPhonenumber(e.target.value)} name="phonenumber" required />
                </p>
                <p>
                    <label>Password</label><br/>
                    <input type="password" name="password" value={Password} onChange={(e)=> setPassword(e.target.value)} required />
                </p>
                <p>
                    <input type="checkbox" name="checkbox" id="checkbox" required /> <span>I agree all statements in <a href="https://google.com" target="_blank" rel="noopener noreferrer">terms of service</a></span>.
                </p>
                <p>
                    <button id="sub_btn" onClick={signUp} type="submit">Register</button>
                </p>
            </form>
            <footer>
                <p><Link to="/login">Already have an account?</Link></p>
            </footer>
            </Paper>
        </Grid>
    )
}

export default Signup;