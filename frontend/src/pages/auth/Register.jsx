import { RiEyeLine, RiEyeCloseLine } from "react-icons/ri";
import { AiOutlineArrowRight } from "react-icons/ai";
import { Link, useNavigate } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup"
import AlertContext from "/src/components/context/AlertContext"
import logo from "/assets/img/drawable/logo.svg"
import http from "../../lib/http";

const Register = () => {
    const [isPasswordVisible, setIsPasswordVisible] = useState(false);
    const [isConfirmPasswordVisible, setIsConfirmPasswordVisible] = useState(false);
    const { setAlert } = useContext(AlertContext)
    const navigator = useNavigate()
    const validator = yup.object({
        first_name: yup.string("Nama depan tidak valid").required("Nama depan harus diisi").min(4, "Nama depan minimal 4 karakter"),
        last_name: yup.string("Nama belakang tidak valid").required("Nama belakang harus diisi").min(4, "Nama belakang minimal 4 karakter"),
        email: yup.string("Email tidak valid").required("Email harus diisi").min(4, "Email terlalu pendek").email("Email tidak valid"),
        password: yup.string("Password tidak valid").required("Password harus diisi").min(8, "Password minimal 8 karakter"),
        confirm_password: yup.string("Konfirmasi Password tidak valid").required("Konfirmasi Password harus diisi").oneOf([yup.ref("password")], "Konfirmasi Password tidak sama")
    })
    const { register, handleSubmit, formState } = useForm({ resolver: yupResolver(validator) })

    async function signup({ first_name, last_name, email, password, confirm_password }) {
        try {
            const req = await http("register", { first_name, last_name, email, password, confirm_password }, { method: "POST" })
            const result = await req.json()
            if (!result.success) {
                throw new Error(result.message)
            }
            setAlert(['success', result.message + " will be redirected to login page in 4 seconds"])
            setTimeout(
                () => {
                    setAlert(['success', result.message + " will be redirected to login page in 4 seconds"])
                    navigator("/auth/login")
                }, 4000
            )
        } catch (error) {
            setAlert(['fail', error.message])
        }
    }

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

    return (
        <>
            <img src={logo} alt="Infinity Shortcut" />
            <span className="text-xl font-bold">Create Account</span>
            <span>Join the elite architects of the web.</span>

            <div className="flex flex-col justify-center gap-4 p-4 rounded shadow border border-black/10 min-w-lg">
                <div className="flex flex-col justify-center">
                    <span className="font-bold text-lg">Welcome Back</span>
                    <span>Please enter your details to sign in.</span>
                </div>
                <form onSubmit={handleSubmit(signup)}>
                    <div className="flex flex-col justify-center gap-4">
                        <div className="flex justify-between items-center gap-4">
                            <div className="flex flex-col justify-center gap-4">
                                <label htmlFor="first_name">First Name</label>
                                <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                                    <input type="text" {...register("first_name")} id="first_name" placeholder="Reza" className="flex-1" />
                                </label>
                                {formState.errors.first_name && (<span className="bg-red-400 p-4 rounded border border-red-700 text-red-700">{formState.errors.first_name.message}</span>)}
                            </div>
                            <div className="flex flex-col justify-center gap-4">
                                <label htmlFor="last_name">Last Name</label>
                                <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                                    <input type="text" {...register("last_name")} id="last_name" placeholder="Fauzan" className="flex-1" />
                                </label>
                                {formState.errors.last_name && (<span className="bg-red-400 p-4 rounded border border-red-700 text-red-700">{formState.errors.last_name.message}</span>)}
                            </div>
                        </div>
                        <label htmlFor="email">Email Address</label>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type="email" {...register("email")} id="email" placeholder="name@snowfoxinfinity.com" className="flex-1" />
                        </label>
                        {formState.errors.email && (<span className="bg-red-400 p-4 rounded border border-red-700 text-red-700">{formState.errors.email.message}</span>)}

                        <div className="flex justify-between items-center gap-4">
                            <label htmlFor="password">Password</label>
                        </div>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type={isPasswordVisible ? "text" : "password"} {...register("password")} id="password" placeholder={isPasswordVisible ? "password" : "*******"} className="flex-1" autoComplete="off" />
                            <button type="button" className="cursor-pointer" onClick={() => { setIsPasswordVisible(!isPasswordVisible) }}>{isPasswordVisible ? <RiEyeLine /> : <RiEyeCloseLine />}</button>
                        </label>
                        <span className="-mt-4 text-gray-700">Minimum 8 characters</span>
                        {formState.errors.password && (<span className="bg-red-400 p-4 rounded border border-red-700 text-red-700">{formState.errors.password.message}</span>)}

                        <div className="flex justify-between items-center gap-4">
                            <label htmlFor="confirm-password">Confirm Password</label>
                        </div>
                        <label className="flex items-center gap-4 p-4 border rounded border-black/40">
                            <input type={isConfirmPasswordVisible ? "text" : "password"} {...register("confirm_password")} id="confirm-password" placeholder={isConfirmPasswordVisible ? "confirm password" : "*******"} className="flex-1" autoComplete="off" />
                            <button type="button" className="cursor-pointer" onClick={() => { setIsConfirmPasswordVisible(!isConfirmPasswordVisible) }}>{isConfirmPasswordVisible ? <RiEyeLine /> : <RiEyeCloseLine />}</button>
                        </label>
                        {formState.errors.confirm_password && (<span className="bg-red-400 p-4 rounded border border-red-700 text-red-700">{formState.errors.confirm_password.message}</span>)}
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