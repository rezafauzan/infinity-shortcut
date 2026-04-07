import { Outlet } from "react-router-dom";
import Footer from "../../components/Footer";
const AuthLayout = () => {
    return (
        <>
            <section>
                <div className="flex flex-col justify-center items-center gap-4 py-4 min-h-screen">
                    <Outlet />
                </div>
            </section >
            <Footer />
        </>
    )
}
export default AuthLayout