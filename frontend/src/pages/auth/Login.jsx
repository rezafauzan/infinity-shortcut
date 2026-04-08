import { AiOutlineGoogle } from "react-icons/ai";
import { RiEyeLine, RiEyeCloseLine } from "react-icons/ri";
import { AiOutlineArrowRight } from "react-icons/ai";
import { Link, useNavigate } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup"
import http from "/src/lib/http.js"
import AlertContext from "/src/components/context/AlertContext"

const Login = () => {
    const [isPasswordVisible, setIsPasswordVisible] = useState(false);
    const { setAlert } = useContext(AlertContext)
    const navigator = useNavigate()
    const validator = yup.object({
        email: yup.string("Email tidak valid").required("Email harus diisi").email("Email tidak valid"),
    })

    const { register, handleSubmit, formState } = useForm({ resolver: yupResolver(validator) })

    useEffect(() => {
        const validateToken = async () => {
            const token = window.localStorage.getItem("token")

            if (!token) return
            
            try {
                const res = await http("validate-token", null, { token })
                
                const result = await res.json()
                if (!result.success) {
                    window.localStorage.removeItem("token")
                    throw new Error(result.message)
                }
                
                navigator("/")
            } catch (error) {
                window.localStorage.removeItem("token")
                setAlert(["fail", error.message])
            }
        }

        validateToken()
    }, [])

    async function login({ email, password }) {
        try {
            const req = await http("login", { email, password }, { method: "POST" })
            const result = await req.json()
            if (!result.success) {
                throw new Error(result.message)
            }
            setAlert(['success', result.message])
            window.localStorage.setItem("token", result.data.token)
            navigator("/")
        } catch (error) {
            setAlert(['fail', error.message])
        }
    }
    return (
        <>
            <span className="text-xl font-bold">Infinity Shortcut</span>

            <div className="flex flex-col justify-center gap-4 p-4 rounded shadow border border-black/10 min-w-lg">
                <div className="flex flex-col justify-center">
                    <span className="font-bold text-lg">Welcome Back</span>
                    <span>Please enter your details to sign in.</span>
                </div>
                <form onSubmit={handleSubmit(login)}>
                    <div className="flex flex-col justify-center gap-4">
                        <label htmlFor="email">Email Address</label>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type="email" {...register("email")} id="email" placeholder="name@snowfoxinfinity.com" className="flex-1" />
                        </label>
                        {formState.errors.email && (<span className="bg-red-400 p-4 rounded border border-red-700 text-red-700">{formState.errors.email.message}</span>)}
                        <div className="flex justify-between items-center gap-4">
                            <label htmlFor="password">Password</label>
                            <Link to="#" className="text-blue-700 hover:text-blue-900">Forgot password?</Link>
                        </div>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type={isPasswordVisible ? "text" : "password"} {...register("password")} id="password" placeholder={isPasswordVisible ? "password" : "*******"} className="flex-1" autoComplete="off" />
                            <button type="button" className="cursor-pointer" onClick={() => { setIsPasswordVisible(!isPasswordVisible) }}>{isPasswordVisible ? <RiEyeLine /> : <RiEyeCloseLine />}</button>
                        </label>
                        <button className="bg-blue-700 hover:bg-blue-800 text-white py-2 flex justify-center items-center gap-4 rounded cursor-pointer">Log In <AiOutlineArrowRight /></button>
                    </div>
                </form>
                <div className="flex items-center gap-4">
                    <div className="grow border-t border-gray-400"></div>
                    <span className="text-gray-400">OR CONTINUE WITH</span>
                    <div className="grow border-t border-gray-400"></div>
                </div>
                <div className="flex justify-center items-center px-4">
                    <button type="button" className="hover:bg-black/10 py-2 border-black/10 flex-1 flex justify-center items-center rounded shadow cursor-pointer"><AiOutlineGoogle />Sign in with Google</button>
                </div>
            </div>

            <span>Don't have an account? <Link to="/auth/register" className="text-blue-700 hover:text-blue-900">Sign up</Link></span>
        </>
    )
}
export default Login