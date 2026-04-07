import { RiEyeLine, RiEyeCloseLine } from "react-icons/ri";
import { AiOutlineArrowRight } from "react-icons/ai";
import { Link } from "react-router-dom";
import { useState } from "react";
import { useForm } from "react-hook-form";
const AuthLayout = () => {
    const [isPasswordVisible, setIsPasswordVisible] = useState(false);
    const { register, handleSubmit } = useForm()
    function login({ email, password }) {
        console.log(email, password)
    }
    return (
        <>
            <section>
                <div className="flex justify-center items-center gap-4 h-screen">
                    <div className="flex flex-col justify-center gap-4 p-4 rounded shadow border border-black/10 min-w-lg">
                        <div className="flex flex-col justify-center">
                            <span className="font-bold text-lg">Welcome Back</span>
                            <span>Please enter your details to sign in.</span>
                        </div>
                        <form onSubmit={handleSubmit(login)}>
                            <div className="flex flex-col justify-center gap-4">
                                <label for="email">Email Address</label>
                                <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                                    <input type="email" {...register("email")} id="email" placeholder="name@snowfoxinfinity.com" className="flex-1" />
                                </label>

                                <div className="flex justify-between items-center gap-4">
                                    <label for="password">Password</label>
                                    <Link to="#">Forgot password?</Link>
                                </div>

                                <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                                    <input type={isPasswordVisible ? "text" : "password"} {...register("password")} id="password" placeholder={isPasswordVisible ? "password" : "*******"} className="flex-1" />
                                    <button type="button" className="cursor-pointer" onClick={() => { setIsPasswordVisible(!isPasswordVisible) }}>{isPasswordVisible ? <RiEyeLine /> : <RiEyeCloseLine />}</button>
                                </label>

                                <button className="bg-blue-700 text-white py-2 flex justify-center items-center gap-4 rounded">Log In <AiOutlineArrowRight /></button>
                            </div>
                        </form>
                    </div>
                </div>
            </section >
        </>
    )
}
export default AuthLayout