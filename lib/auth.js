import axios from 'axios';
export  const registeruser = async(Firstname, Lastname, email, phonenumber, ) =>{
const {data} = await axios.post('' ,{ Firstname, Lastname , email, phonenumber,});
console.log(data)
}