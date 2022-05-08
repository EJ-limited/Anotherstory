import axios from 'axios';
export  const registeruser = async(username, email, phonenumber, password ) =>{
const {data} = await axios.post('' ,{ username, email, phonenumber,password});
console.log(data)
}