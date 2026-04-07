import { BsFillLightningChargeFill } from "react-icons/bs";
import { AiOutlineLink } from "react-icons/ai";
import { Link, Outlet } from "react-router-dom";
import Navbar from "/src/components/Navbar";
import { useForm } from "react-hook-form";
import stockPhoto from "/assets/img/value-propotition-stock-photo.png"

const Hero = () => {
    const { register, handleSubmit } = useForm()
    function cutLink({ link }) {
        console.log(link)
    }
    return (
        <section>
            <div className="flex flex-col justify-center items-center gap-4 p-4">
                <h2 className="font-bold text-4xl">
                    <span>Shorten URLs.</span> <span className="text-blue-700">Share Easily.</span>
                </h2>
                <p>
                    Create short, memorable links for your team communications.
                    Transform long, cumbersome URLs into powerful digital assets that
                    drive engagement.
                </p>
                <div className="cta flex justify-center items-center gap-4">
                    <Link to="/" className="bg-blue-700 hover:bg-blue-900 text-white font-bold h-10 px-4 flex justify-center items-center gap-4 border border-black/10 rounded shadow cursor-pointer">Get Started</Link>
                    <Link to="/" className="text-blue-700 hover:bg-black/10 font-bold h-10 px-4 flex justify-center items-center gap-4 border border-black/10 rounded shadow cursor-pointer">Learn More</Link>
                </div>
                <div className="bg-white md:w-3/4 shadow">
                    <form onSubmit={handleSubmit(cutLink)}>
                        <div className="flex items-center">
                            <label className="h-16 flex items-center gap-4 p-4 border border-black/40 border-r-0 rounded rounded-r-none flex-1">
                                <AiOutlineLink /><input type="text" {...register("link")} id="link" placeholder="https://your-long-domain.com/with-your-very-long-path" className="flex-1" />
                            </label>
                            <button className="bg-blue-700 hover:bg-blue-900 text-white font-bold h-16 px-4 flex justify-center items-center gap-4 border border-black/10 rounded rounded-l-none shadow cursor-pointer">Make it short !</button>
                        </div>
                    </form>
                </div>
            </div>
        </section>
    )
}

const Features = () => {
    return (
        <section className="bg-[#F4F4F5]">
            <div className="flex flex-col gap-4 p-4">
                <div className="section-header">
                    <h2 className="font-bold text-blue-700">
                        Architectural Features
                    </h2>
                    <p className="font-bold">
                        Built for Enterprise Precision
                    </p>
                </div>
                <div className="section-body">
                    <div className="flex flex-col md:flex-row justify-evenly items-stretch gap-4">
                        <div className="bg-white p-4 flex-1 flex flex-col justify-center gap-4 rounded shadow">
                            <div className="bg-[#DBE1FF] w-10 h-10 flex justify-center items-center gap-4 rounded">
                                <BsFillLightningChargeFill className="text-[#004AD6]" />
                            </div>
                            <h3>Easy Create</h3>
                            <p>
                                Instantly generate high-performance short links with a single click or through our surgical API endpoints.
                            </p>
                        </div>
                        <div className="bg-white p-4 flex-1 flex flex-col justify-center gap-4 rounded shadow">
                            <div className="bg-[#ADBFFF] w-10 h-10 flex justify-center items-center gap-4 rounded">
                                <svg width="18" height="16" viewBox="0 0 18 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M0 10V8H7V10H0ZM0 6V4H11V6H0ZM0 2V0H11V2H0ZM9 16V12.925L14.525 7.425C14.675 7.275 14.8417 7.16667 15.025 7.1C15.2083 7.03333 15.3917 7 15.575 7C15.775 7 15.9667 7.0375 16.15 7.1125C16.3333 7.1875 16.5 7.3 16.65 7.45L17.575 8.375C17.7083 8.525 17.8125 8.69167 17.8875 8.875C17.9625 9.05833 18 9.24167 18 9.425C18 9.60833 17.9667 9.79583 17.9 9.9875C17.8333 10.1792 17.725 10.35 17.575 10.5L12.075 16H9ZM16.5 9.425L15.575 8.5L16.5 9.425ZM10.5 14.5H11.45L14.475 11.45L14.025 10.975L13.55 10.525L10.5 13.55V14.5ZM14.025 10.975L13.55 10.525L14.475 11.45L14.025 10.975Z" fill="#394C84" />
                                </svg>
                            </div>
                            <h3>Custom Slugs</h3>
                            <p>
                                Maintain brand authority with readable, custom link endings that resonate with your digital audience.
                            </p>
                        </div>
                        <div className="bg-white p-4 flex-1 flex flex-col justify-center gap-4 rounded shadow">
                            <div className="bg-[#FFDBDD] w-10 h-10 flex justify-center items-center gap-4 rounded">
                                <svg width="24" height="12" viewBox="0 0 24 12" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M0 12V10.425C0 9.70833 0.366667 9.125 1.1 8.675C1.83333 8.225 2.8 8 4 8C4.21667 8 4.425 8.00417 4.625 8.0125C4.825 8.02083 5.01667 8.04167 5.2 8.075C4.96667 8.425 4.79167 8.79167 4.675 9.175C4.55833 9.55833 4.5 9.95833 4.5 10.375V12H0ZM6 12V10.375C6 9.84167 6.14583 9.35417 6.4375 8.9125C6.72917 8.47083 7.14167 8.08333 7.675 7.75C8.20833 7.41667 8.84583 7.16667 9.5875 7C10.3292 6.83333 11.1333 6.75 12 6.75C12.8833 6.75 13.6958 6.83333 14.4375 7C15.1792 7.16667 15.8167 7.41667 16.35 7.75C16.8833 8.08333 17.2917 8.47083 17.575 8.9125C17.8583 9.35417 18 9.84167 18 10.375V12H6ZM19.5 12V10.375C19.5 9.94167 19.4458 9.53333 19.3375 9.15C19.2292 8.76667 19.0667 8.40833 18.85 8.075C19.0333 8.04167 19.2208 8.02083 19.4125 8.0125C19.6042 8.00417 19.8 8 20 8C21.2 8 22.1667 8.22083 22.9 8.6625C23.6333 9.10417 24 9.69167 24 10.425V12H19.5ZM8.125 10H15.9C15.7333 9.66667 15.2708 9.375 14.5125 9.125C13.7542 8.875 12.9167 8.75 12 8.75C11.0833 8.75 10.2458 8.875 9.4875 9.125C8.72917 9.375 8.275 9.66667 8.125 10ZM4 7C3.45 7 2.97917 6.80417 2.5875 6.4125C2.19583 6.02083 2 5.55 2 5C2 4.43333 2.19583 3.95833 2.5875 3.575C2.97917 3.19167 3.45 3 4 3C4.56667 3 5.04167 3.19167 5.425 3.575C5.80833 3.95833 6 4.43333 6 5C6 5.55 5.80833 6.02083 5.425 6.4125C5.04167 6.80417 4.56667 7 4 7ZM20 7C19.45 7 18.9792 6.80417 18.5875 6.4125C18.1958 6.02083 18 5.55 18 5C18 4.43333 18.1958 3.95833 18.5875 3.575C18.9792 3.19167 19.45 3 20 3C20.5667 3 21.0417 3.19167 21.425 3.575C21.8083 3.95833 22 4.43333 22 5C22 5.55 21.8083 6.02083 21.425 6.4125C21.0417 6.80417 20.5667 7 20 7ZM12 6C11.1667 6 10.4583 5.70833 9.875 5.125C9.29167 4.54167 9 3.83333 9 3C9 2.15 9.29167 1.4375 9.875 0.8625C10.4583 0.2875 11.1667 0 12 0C12.85 0 13.5625 0.2875 14.1375 0.8625C14.7125 1.4375 15 2.15 15 3C15 3.83333 14.7125 4.54167 14.1375 5.125C13.5625 5.70833 12.85 6 12 6ZM12 4C12.2833 4 12.5208 3.90417 12.7125 3.7125C12.9042 3.52083 13 3.28333 13 3C13 2.71667 12.9042 2.47917 12.7125 2.2875C12.5208 2.09583 12.2833 2 12 2C11.7167 2 11.4792 2.09583 11.2875 2.2875C11.0958 2.47917 11 2.71667 11 3C11 3.28333 11.0958 3.52083 11.2875 3.7125C11.4792 3.90417 11.7167 4 12 4Z" fill="#360F00" />
                                </svg>
                            </div>
                            <h3>Team Ready</h3>
                            <p>
                                Collaborate across departments with shared workspaces, permissions, and unified analytics dashboards.
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    )
}

const ValuePropotition = () => {
    return (
        <section className="bg-[#FFFFFF]">
            <div className="flex flex-col md:flex-row gap-4 p-4">
                <div className="section-header flex-1 flex justify-center items-center">
                    <div className="rounded w-lg h-lg overflow-hidden">
                        <img src={stockPhoto} alt="Data Driven Insights Stock Photo" />
                    </div>
                </div>
                <div className="section-body flex-1 flex justify-center items-center">
                    <div className="flex flex-col justify-center gap-4">
                        <span className="text-gray-700 text-sm font-bold">Data Driven Insights</span>
                        <h2 className="font-bold text-xl">Observe your link architecture in real-time.</h2>
                        <p>
                            Every click is a data point. Our dashboard provides surgical precision into
                            where your traffic originates, who is engaging, and how your team
                            communications are performing across the globe.
                        </p>
                        <ul className="flex flex-col justify-center gap-4 list-image-[url('/assets/img/drawable/custom-list-icon.svg')]">
                            <li className="font-bold">Geographic Distribution Maps</li>
                            <li className="font-bold">Device & Browser Breakdown</li>
                            <li className="font-bold">UTM Parameter Tracking</li>
                        </ul>
                    </div>
                </div>
            </div>
        </section>
    )
}

const HomeLayout = () => {
    return (
        <>
            <div className="wrapper">
                <Navbar />
                <Hero />
                <Features />
                <ValuePropotition />
            </div>
        </>
    )
}
export default HomeLayout