const  Hero = () => {
    return (
        <>
        <div className="bg-[url(/img/bg-1.jpg)] bg-cover bg-center h-screen " >
            <div className="p-10 text-center " >
                <h1  className="text-[30px] my-2 font-bold text-white  "  > Another Story Auditions  </h1>
                <h2  className="text-[10px] my-2 " > Get ready  </h2>
                <button  className=" text-[20px] font-bold my-2  text-white  bg-red-700 hover:bg-blue-700  py-2 px-4 rounded-full " href="/register" > Register to Audition  </button>
            </div>

        </div>
        </>

    )
}
 export default Hero