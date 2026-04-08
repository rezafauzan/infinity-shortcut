import { AiOutlinePlus } from "react-icons/ai";
import { useContext, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import UserContext from "./context/UserContext"
import { FiLogOut } from "react-icons/fi"
import { AiOutlineUser } from "react-icons/ai"
import AlertContext from "/src/components/context/AlertContext"

const Navbar = () => {
    const { user, setUser } = useContext(UserContext)
    const [userDropdown, setUserDropdown] = useState(false)
    const { setAlert } = useContext(AlertContext)
    const navigate = useNavigate()

    function toggleDropdown(setter, getter) {
        setter(!getter)
    }

    function logout() {
        window.localStorage.removeItem("token")
        setUser(null)
        setAlert(["success", "Anda berhasil logout !"])
        navigate("/")
    }

    return (
        <nav className="h-16 px-4 border-b border-b-black/10 flex justify-between items-center gap-4 shadow relative">
            <div className="flex items-center gap-4">
                <div className="brand">
                    <Link to="/" className="font-bold">Infinity Shortcut</Link>
                </div>
                <ul className="hidden md:flex justify-evenly items-center gap-4">
                    <li><Link to="/dashboard" className="text-gray-700 hover:text-gray-900">Dashboard</Link></li>
                    <li><Link to="/" className="text-gray-700 hover:text-gray-900">Analytics</Link></li>
                    <li><Link to="/" className="text-gray-700 hover:text-gray-900">Links</Link></li>
                </ul>
            </div>
            <div className="hidden md:flex justify-evenly items-center gap-4">
                {
                    user != null
                        ?
                        <div className="text-white hidden md:block">
                            <div className="flex justify-between items-center gap-4">
                                <div className="flex justify-center items-center gap-4 text-black/40 hover:text-black cursor-pointer" onClick={() => { toggleDropdown(setUserDropdown, userDropdown) }}>
                                    <Link to="/dashboard/new-link" className="bg-blue-700 hover:bg-blue-900 text-white h-10 px-4 flex justify-center items-center gap-4 border border-black/10 rounded shadow cursor-pointer">Create new link <AiOutlinePlus /></Link>
                                    <div className="rounded-full w-10 h-10 overflow-hidden">
                                        <img src={"https://i.pravatar.cc/1000?img=54"} alt={user.first_name} />
                                    </div>
                                    <span className="hidden md:block">{user.first_name}</span>
                                </div>

                                <div className={"absolute bg-white border border-black/10 shadow w-40 h-40 -bottom-40 right-0 flex-col justify-center items-center gap-4 p-4 rounded" + (userDropdown ? " flex" : " hidden")}>
                                    <button className="w-full hover:text-black text-black/40 cursor-pointer" onClick={logout}>
                                        <span className="flex items-center gap-4 text-xs"><FiLogOut className="text-lg" />Logout</span>
                                    </button>
                                    <Link to="" className="w-full hover:text-black text-black/40 cursor-pointer"><span className="flex items-center gap-4 text-xs"><AiOutlineUser className="text-lg" />Profile</span></Link>
                                </div>
                            </div>
                        </div>
                        :
                        <Link to="/auth/login" className="hover:bg-black/10 h-10 px-4 flex justify-center items-center gap-4 border border-black/10 rounded shadow cursor-pointer">Login</Link>

                }
                {
                    user != null
                        ?
                        ""
                        :
                        <Link to="/auth/register" className="bg-blue-700 hover:bg-blue-900 text-white h-10 px-4 flex justify-center items-center gap-4 border border-black/10 rounded shadow cursor-pointer">Register</Link>
                }
            </div>
        </nav>
    )
}
export default Navbar