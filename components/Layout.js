import Hero from "./Hero"
import RegisterForm from "./RegisterForm"


const Layout = ({Children}) =>{
    return(
        <div  className="overflow-x-hidden bg-no-repeat bg-center bg-cover bg-[url('/img/bg-1.jpg')]">
            <RegisterForm/>
            <Hero/>
        </div>

        

        
        
        

    )
}

export default Layout