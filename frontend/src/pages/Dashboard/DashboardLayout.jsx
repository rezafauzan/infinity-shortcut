import { Outlet, useNavigate } from "react-router-dom"
import Navbar from "../../components/Navbar"
import Footer from "../../components/Footer"
import AlertContext from "../../components/context/AlertContext"
import { useContext, useEffect, useState } from "react"
import http from "../../lib/http"
import UserContext from "../../components/context/UserContext"


const DashboardLayout = () => {
    const [user, setUser] = useState(null)
    const [loading, setLoading] = useState(true)
    const { setAlert } = useContext(AlertContext)
    const navigate = useNavigate()

    useEffect(() => {
        const validateToken = async () => {
            const token = window.localStorage.getItem("token")

            try {
                const req = await http("validate-token", null, { token })

                const result = await req.json()

                if (!result.success) {
                    window.localStorage.removeItem("token")
                    throw new Error(result.message)
                }
                setUser(result.data)
            } catch (error) {
                window.localStorage.removeItem("token")
                setAlert(["fail", "Session expired please relogin! " + error])
                navigate("/auth/login")
            } finally {
                setLoading(false)
            }
        }

        validateToken()
    }, [])

    if (loading) {
        return (
            <div className="fixed top-0 left-0 right-0 bottom-0 bg-black/40 backdrop-blur-lg flex justify-center items-center z-10">
                <div className="bg-green-400 text-green-700 w-[50%] h-[50%] flex items-center justify-center relative rounded">
                    <span className="text-green-700 p-4 font-bold">
                        Loading...
                    </span>
                </div>
            </div>
        )
    }

    return (
        <>
            <UserContext value={{ user, setUser }}>
                <Navbar />
            </UserContext>
            <Outlet />
            <Footer />
        </>
    )
}

export default DashboardLayout