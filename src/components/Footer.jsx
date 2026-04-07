import { Link } from "react-router-dom"

const Footer = () => {
    return (
        <footer>
            <nav className="bg-[#F8FAFD] text-gray-700 h-16 px-4 border-t border-t-black/10 flex justify-between items-center gap-4 shadow">
                <div className="flex justify-between items-center gap-4 w-full">
                    <div className="copyright">
                        <span>© 2026 Infinity Shortcut. <a href="https://www.linkedin.com/in/reza-fauzan-adhima/" className="text-blue-700 hover:text-blue-900">SnowfoxInfinity❄️</a>.</span>
                    </div>
                    <ul className="hidden md:flex justify-evenly items-center gap-4">
                        <li><Link to="/" className="hover:text-gray-900">Privacy Policy</Link></li>
                        <li><Link to="/" className="hover:text-gray-900">Terms of Service</Link></li>
                        <li><Link to="/" className="hover:text-gray-900">API Documentation</Link></li>
                        <li><Link to="/" className="hover:text-gray-900">Support</Link></li>
                    </ul>
                </div>
            </nav>
        </footer>
    )
}
export default Footer