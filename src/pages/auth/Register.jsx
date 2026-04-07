import { RiEyeLine, RiEyeCloseLine } from "react-icons/ri";
import { AiOutlineArrowRight } from "react-icons/ai";
import { Link } from "react-router-dom";
import { useState } from "react";
import { useForm } from "react-hook-form";

const Register = () => {
    const [isPasswordVisible, setIsPasswordVisible] = useState(false);
    const [isConfirmPasswordVisible, setIsConfirmPasswordVisible] = useState(false);
    const { register, handleSubmit } = useForm()
    function signup({ email, password }) {
        console.log(email, password)
    }
    return (
        <>
            <img src="/assets/img/drawable/logo.svg" alt="Infinity Shortcut" />
            <span className="text-xl font-bold">Create Account</span>
            <span>Join the elite architects of the web.</span>

            <div className="flex flex-col justify-center gap-4 p-4 rounded shadow border border-black/10 min-w-lg">
                <div className="flex flex-col justify-center">
                    <span className="font-bold text-lg">Welcome Back</span>
                    <span>Please enter your details to sign in.</span>
                </div>
                <form onSubmit={handleSubmit(signup)}>
                    <div className="flex flex-col justify-center gap-4">
                        <label htmlFor="email">Email Address</label>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type="email" {...register("email")} id="email" placeholder="name@snowfoxinfinity.com" className="flex-1" />
                        </label>
                        <div className="flex justify-between items-center gap-4">
                            <label htmlFor="password">Password</label>
                        </div>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type={isPasswordVisible ? "text" : "password"} {...register("password")} id="password" placeholder={isPasswordVisible ? "password" : "*******"} className="flex-1" autoComplete="off"/>
                            <button type="button" className="cursor-pointer" onClick={() => { setIsPasswordVisible(!isPasswordVisible) }}>{isPasswordVisible ? <RiEyeLine /> : <RiEyeCloseLine />}</button>
                        </label>
                        <span className="-mt-4 text-gray-700">Minimum 8 characters</span>

                        <div className="flex justify-between items-center gap-4">
                            <label htmlFor="confirm-password">Confirm Password</label>
                        </div>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type={isConfirmPasswordVisible ? "text" : "password"} {...register("confirm_password")} id="confirm-password" placeholder={isConfirmPasswordVisible ? "confirm password" : "*******"} className="flex-1" autoComplete="off"/>
                            <button type="button" className="cursor-pointer" onClick={() => { setIsConfirmPasswordVisible(!isConfirmPasswordVisible) }}>{isConfirmPasswordVisible ? <RiEyeLine /> : <RiEyeCloseLine />}</button>
                        </label>
                        <button className="bg-blue-700 hover:bg-blue-800 text-white py-2 flex justify-center items-center gap-4 rounded cursor-pointer">Sign Up <AiOutlineArrowRight /></button>
                    </div>
                </form>
                <span>By signing up, you agree to our <Link className="text-blue-700 hover:text-blue-900">Terms of Service</Link> and <Link className="text-blue-700 hover:text-blue-900">Privacy Policy.</Link></span>
            </div>

            <span>Already have an account?  <Link to="/auth/login" className="text-blue-700 hover:text-blue-900">Log In</Link></span>
        </>
    )
}
export default Register