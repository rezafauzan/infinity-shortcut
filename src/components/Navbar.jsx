import { Link } from "react-router-dom"

const Nvabar = () => {
    return (
        <nav className="h-16 px-4 border-b border-b-black/10 flex justify-between items-center gap-4 shadow">
            <div className="flex items-center gap-4">
                <div className="brand">
                    <Link to="/" className="font-bold">Infinity Shortcut</Link>
                </div>
                <ul className="hidden md:flex justify-evenly items-center gap-4">
                    <li><Link to="/">Dashboard</Link></li>
                    <li><Link to="/">Analytics</Link></li>
                    <li><Link to="/">Links</Link></li>
                </ul>
            </div>
            <div className="hidden md:flex justify-evenly items-center gap-4">
                <button className="hover:bg-black/10 h-10 px-4 flex justify-center items-center gap-4 border border-black/10 rounded shadow cursor-pointer">Login</button>
                <button className="bg-blue-700 hover:bg-blue-900 text-white h-10 px-4 flex justify-center items-center gap-4 border border-black/10 rounded shadow cursor-pointer">Register</button>
            </div>
        </nav>
    )
}
export default Nvabar