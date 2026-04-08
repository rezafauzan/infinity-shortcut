import { useForm, useWatch } from "react-hook-form"
import { AiOutlineLink, AiOutlineArrowLeft } from "react-icons/ai"
import { HiOutlineEye } from "react-icons/hi"
import { BsFillLightningChargeFill, BsQrCode } from "react-icons/bs"
import { IoAnalyticsOutline } from "react-icons/io5"

const CreateLink = () => {
    const { register, handleSubmit, control, reset } = useForm({
        defaultValues: {
            destinationUrl: "",
            customSlug: ""
        }
    })

    const customSlug = useWatch({
        control,
        name: "customSlug",
    })

    const onSubmit = (data) => {
        console.log("Creating Link:", data)
    }

    return (
        <section className="min-h-screen p-6 md:p-12 ">
            <div className="max-w-4xl mx-auto">

                <button className="flex items-center gap-2 text-blue-500 font-semibold hover:text-blue-400 transition-colors mb-8">
                    <AiOutlineArrowLeft /> Back to Dashboard
                </button>

                <div className="mb-10">
                    <h1 className="text-4xl font-bold mb-2">Create New Short Link</h1>
                    <p className="text-gray-400">Transform your long URLs into clean, manageable assets.</p>
                </div>

                <div className="bg-white rounded-3xl p-8 md:p-12 text-slate-800 shadow-2xl mb-12">
                    <form onSubmit={handleSubmit(onSubmit)} className="space-y-8">

                        <div className="flex flex-col gap-3">
                            <label className="font-bold text-sm tracking-widest uppercase text-slate-700">
                                Destination URL <span className="text-red-500">*</span>
                            </label>
                            <div className="relative flex items-center">
                                <AiOutlineLink className="absolute left-4 text-slate-400 text-xl" />
                                <input
                                    {...register("destinationUrl", { required: true })}
                                    type="text"
                                    placeholder="https://example.com/your-long-url-here"
                                    className="w-full bg-slate-50 border border-slate-200 rounded-xl py-4 pl-12 pr-4 outline-none focus:border-blue-500 transition-all text-slate-600"
                                />
                            </div>
                            <p className="text-slate-400 text-xs italic">Ensure your URL starts with http:// or https://</p>
                        </div>

                        <div className="flex flex-col gap-3">
                            <label className="font-bold text-sm tracking-widest uppercase text-slate-700">
                                Custom Slug (Optional)
                            </label>
                            <div className="flex items-center">
                                <div className="bg-slate-100 border border-slate-200 border-r-0 rounded-l-xl py-4 px-6 text-slate-500 font-medium">
                                    short.link/
                                </div>
                                <input
                                    {...register("customSlug")}
                                    type="text"
                                    placeholder="my-custom-slug"
                                    className="flex-1 bg-white border border-slate-200 rounded-r-xl py-4 px-4 outline-none focus:border-blue-500 transition-all text-slate-600"
                                />
                            </div>
                            <p className="text-slate-400 text-xs italic">Leave blank to generate a random unique identifier.</p>
                        </div>

                        <div className="bg-blue-50/50 border border-blue-100 rounded-2xl p-6 flex items-start gap-4">
                            <HiOutlineEye className="text-blue-600 text-2xl mt-1" />
                            <div>
                                <span className="text-blue-700 font-bold text-xs tracking-widest uppercase block mb-1">Live Preview</span>
                                <p className="text-lg text-slate-800">
                                    Your short link will be: <span className="text-blue-600 font-semibold underline break-all">https://short.link/{customSlug || "random-id"}</span>
                                </p>
                            </div>
                        </div>

                        <div className="flex items-center gap-8 pt-4">
                            <button
                                type="submit"
                                className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-4 px-10 rounded-xl shadow-lg shadow-blue-200 flex items-center gap-3 transition-all active:scale-95"
                            >
                                Create Link <BsFillLightningChargeFill />
                            </button>
                            <button
                                type="button"
                                onClick={() => reset()}
                                className="text-slate-500 font-bold hover:text-slate-800 transition-colors"
                            >
                                Cancel
                            </button>
                        </div>
                    </form>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-12 px-4">
                    <div className="flex items-center gap-4">
                        <div className="bg-orange-100 p-4 rounded-full">
                            <IoAnalyticsOutline className="text-orange-600 text-2xl" />
                        </div>
                        <div>
                            <h4 className="font-bold text-gray-200">Real-time Analytics</h4>
                            <p className="text-gray-500 text-sm">Track every click, geographical location, and referral source instantly.</p>
                        </div>
                    </div>
                    <div className="flex items-center gap-4">
                        <div className="bg-blue-100 p-4 rounded-full">
                            <BsQrCode className="text-blue-600 text-2xl" />
                        </div>
                        <div>
                            <h4 className="font-bold text-gray-200">Auto-generated QR</h4>
                            <p className="text-gray-500 text-sm">Every link automatically creates a high-resolution QR code for print.</p>
                        </div>
                    </div>
                </div>

            </div>
        </section>
    )
}

export default CreateLink