import { BiBarChartAlt2 } from "react-icons/bi";
import { BsFillCalendarDateFill } from "react-icons/bs";
import { AiOutlineLink, AiOutlineSearch, AiOutlineCopy, AiOutlineDelete } from "react-icons/ai"
import { HiOutlineChevronLeft, HiOutlineChevronRight } from "react-icons/hi"
import { useForm } from "react-hook-form"
import { useContext, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import AlertContext from "/src/components/context/AlertContext"
import http from "../../lib/http"

const Dashboard = () => {
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
            if (userLinks.length === 1) {
                setUserLinks([])
            }
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
        <>
            <div className="max-w-4xl min-h-128 mx-auto flex flex-col gap-4 p-4">
                <div className="section-header">
                    <div className="flex justify-between items-end">
                        <div>
                            <h1 className="text-4xl font-bold text-gray-800">My Links</h1>
                            <p className="text-gray-500 mt-2">Manage and track your shortened digital assets.</p>
                        </div>
                        <div className="text-right">
                            <span className="text-gray-400 font-semibold uppercase text-sm tracking-wider">Total Active</span>
                            <div className="text-4xl font-bold text-blue-700">{(userLinks && userLinks.length.toLocaleString("id-ID"))}</div>
                        </div>
                    </div>
                </div>
                <div className="section-body flex flex-col justify-center gap-4 min-h-128">
                    <div className="bg-white shadow">
                        <form onSubmit={handleSubmit(search)}>
                            <div className="flex items-center">
                                <label className="h-16 flex items-center gap-4 p-4 border border-black/40 border-r-0 rounded rounded-r-none flex-1">
                                    <AiOutlineSearch /><input type="text" {...register("link")} id="link" placeholder="https://your-long-domain.com/with-your-very-long-path" className="flex-1" />
                                </label>
                                <button className="bg-blue-700 hover:bg-blue-900 text-white font-bold h-16 px-4 flex justify-center items-center gap-4 border border-black/10 rounded rounded-l-none shadow cursor-pointer">Search !</button>
                            </div>
                        </form>
                    </div>

                    <div className="flex flex-col gap-4">
                        {
                            userLinks.length < 1 ?
                                <div className="text-center text-gray-400 mt-10">
                                    You don't have any links yet 😢 Start by creating your first short link now!
                                </div>
                                :
                                userLinks && userLinks.map((link) => (
                                    <div key={link.id} className="bg-white p-6 rounded-2xl shadow-sm border border-slate-100 hover:border-blue-300 transition-all group">
                                        <div className="flex justify-between items-start">
                                            <div className="flex flex-col gap-2">
                                                <div className="flex items-center gap-2 text-blue-700 font-bold text-lg">
                                                    <AiOutlineLink className="text-xl" />
                                                    <a href={`http://localhost:8888/api/links/${link.short_url}`} className="hover:underline">{`http://localhost:8888/api/links/${link.short_url}`}</a>
                                                </div>
                                                <a href={`https://${link.original_url}`} className="text-gray-400 text-sm truncate max-w-md">
                                                    {`https://${link.original_url}`}
                                                </a>
                                                <div className="flex items-center gap-6 mt-2 text-gray-400 font-bold text-xs tracking-widest uppercase">
                                                    <div className="flex items-center gap-2">
                                                        <BsFillCalendarDateFill className="text-xl" />
                                                        {link.date}
                                                    </div>
                                                    <div className="flex items-center gap-2">
                                                        <BiBarChartAlt2 className="text-xl" />
                                                        {(1000000).toLocaleString("id-ID")} Clicks
                                                    </div>
                                                </div>
                                            </div>

                                            <div className="flex gap-2">
                                                <button className="p-3 bg-slate-50 text-blue-700 rounded-xl hover:bg-blue-700 hover:text-white transition-all shadow-sm" onClick={() => { copyShortLink(`http://localhost:8888/api/links/${link.short_url}`) }}>
                                                    <AiOutlineCopy className="text-xl" />
                                                </button>
                                                <button className="p-3 bg-slate-50 text-gray-400 rounded-xl hover:bg-red-50 hover:text-red-500 transition-all shadow-sm" onClick={() => { deleteShortLink(`${link.id}`) }}>
                                                    <AiOutlineDelete className="text-xl" />
                                                </button>
                                            </div>
                                        </div>
                                    </div>
                                ))
                        }
                    </div>

                    <div className="flex justify-between items-center mt-auto text-gray-700 font-semibold">
                        <button className="flex items-center gap-2 hover:text-blue-700 text-gray-400 transition-colors cursor-pointer">
                            <HiOutlineChevronLeft className="text-xl" /> Prev Page
                        </button>

                        <div className="flex items-center gap-4">
                            <span className="bg-blue-100 text-blue-700 w-10 h-10 flex justify-center items-center rounded-lg shadow-inner">1</span>
                            <span className="text-gray-400 px-2 font-bold">of</span>
                            <span className="w-10 h-10 flex justify-center items-center">5</span>
                        </div>

                        <button className="flex items-center gap-2 hover:text-blue-700 transition-colors text-gray-400 cursor-pointer">
                            Next <HiOutlineChevronRight className="text-xl" />
                        </button>
                    </div>
                </div>
            </div>
        </>
    )
}
export default Dashboard