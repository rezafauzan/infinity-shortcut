import { BiBarChartAlt2 } from "react-icons/bi";
import { BsFillCalendarDateFill } from "react-icons/bs";
import { AiOutlineLink, AiOutlineSearch, AiOutlineCopy, AiOutlineDelete } from "react-icons/ai"
import { HiOutlineChevronLeft, HiOutlineChevronRight } from "react-icons/hi"
import { IoStatsChartOutline } from "react-icons/io5"
import { LuCalendarDays } from "react-icons/lu"
import { useForm } from "react-hook-form"
import { useContext, useEffect, useState } from "react"
import { Outlet, useNavigate } from "react-router-dom"
import AlertContext from "/src/components/context/AlertContext"
import http from "../../lib/http"

const DashboardLayout = () => {
    const { register, handleSubmit } = useForm()
    const navigate = useNavigate()
    const { setAlert } = useContext(AlertContext)
    const [userLinks, setUserLinks] = useState([])
    const [loading, setLoading] = useState(true)

    function search({ slug }) {
        console.log(slug)
    }

    const copyShortLink = async (link) => {
        try {
            await navigator.clipboard.writeText(link)
            setAlert(["success", "Shortcut link copied to clipboard! " + link])
        } catch (err) {
            setAlert(["fail", "Shortcut link failed to copy to clipboard! " + err])
        }
    }

    const deleteShortLink = async (id) => {
        const token = window.localStorage.getItem("token")

        if (!token) {
            navigate("/auth/login")
        }

        try {
            const req = await http("links/" + id, null, { method: "DELETE", token })

            const result = await req.json()
            if (!result.success) {
                throw new Error(result.message)
            }
            setAlert(["success", result.message])

            const reqUserLinks = await http("links", null, { token })

            const resultUserLinks = await reqUserLinks.json()
            if (!resultUserLinks.success) {
                window.localStorage.removeItem("token")
                throw new Error(resultUserLinks.message)
            }

            setUserLinks(resultUserLinks.data)
        } catch (error) {
            if (error.message.includes("Unauthorized access")) {
                window.localStorage.removeItem("token")
                setAlert(["fail", "Session expired please relogin!"])
                navigate("/auth/login")
                return
            }
            setAlert(["fail", error.message])
        }
    }

    useEffect(() => {
        const validateToken = async () => {
            const token = window.localStorage.getItem("token")

            if (!token) {
                navigate("/auth/login")
            }

            try {
                const req = await http("validate-token", null, { token })

                const result = await req.json()
                if (!result.success) {
                    window.localStorage.removeItem("token")
                    throw new Error(result.message)
                }

                const reqUserLinks = await http("links", null, { token })

                const resultUserLinks = await reqUserLinks.json()
                if (!resultUserLinks.success) {
                    window.localStorage.removeItem("token")
                    throw new Error(resultUserLinks.message)
                }

                if (resultUserLinks.data != null) {
                    setUserLinks(resultUserLinks.data)
                }

            } catch (error) {
                window.localStorage.removeItem("token")
                setAlert(["fail", "Session expired please relogin!"])
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
        <section className="bg-[#F4F4F5] min-h-screen p-6 md:p-12">
            <div className="max-w-5xl mx-auto flex flex-col gap-8">
                <Outlet />
            </div>
        </section>
    )
}

export default DashboardLayout