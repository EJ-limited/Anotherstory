import { registeruser   } from "../lib/auth";
import react from "react";
import register from "../pages/Register";
class RegisterForm extends react.Component{
    state ={
        username:'',
        email:'',
        phonenumber:'',
        password:''
    };

handleChange =  event => {
    const {username, email, phonenumber,password} = this.state;

    this.setState({[event.target.name]:event.target.value});
    registeruser( username,email,phonenumber,password);

} 

handlesubmit = event =>{
    event.preventDefault();
    console.log(this.state);
}
    render(){
        return(
            <div className="flex items-center justify-center min-h-full px-4 py-12 sm:px-6 lg:px-8" >
            <form className=" mt-8 space-y-6  bg-[url('/img/bg(1).jpg')] "  action="#" method="POST"  onSubmit={this.handlesubmit} >
                <div className="-space-y-px rounded-md shadow-sm" >
                    <input type="username"   
                name="username"
                placeholder="username"
                onChange={this.handleChange}
                /></div>
                <div className="-space-y-px rounded-md shadow-sm" ><input type="email" 
                name="email"
                placeholder="email"
                onChange={this.handleChange}
                /></div>
                 <div className="-space-y-px rounded-md shadow-sm" ><input type="phonenumber" 
                name="phonenumber"
                placeholder="phonenumber"
                onChange={this.handleChange}
                /></div>
                 <div className="-space-y-px rounded-md shadow-sm" ><input type="password" 
                name="password"
                placeholder="password"
                onChange={this.handleChange}
                /> </div>
                <button  type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-red-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" 
                >submit</button>
            </form>
            </div>

        )
    }
}

export default RegisterForm