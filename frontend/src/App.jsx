import { createBrowserRouter, RouterProvider } from "react-router-dom"
import AuthLayout from "./pages/auth/AuthLayout"
import DashboardLayout from "./pages/Dashboard/DashboardLayout"
import Login from "./pages/auth/Login"
import Register from "./pages/auth/Register"
import Home from "./pages/Home/Home"
import AlertContext from "/src/components/context/AlertContext"
import { useRef, useState } from "react"
import { AiOutlineCloseCircle } from "react-icons/ai"
import Dashboard from "./pages/Dashboard/Dashboard"
import CreateLink from "./pages/Dashboard/CreateLink"

function App() {
  const [alert, setAlert] = useState([])
  const modal = useRef()

  function modalRemove() {
    setAlert([])
  }

  const router = createBrowserRouter(
    [
      {
        path: '/',
        element: <Home />
      },
      {
        path: '/auth',
        element: <AuthLayout />,
        children: [
          {
            path: 'login',
            element: <Login />
          },
          {
            path: 'register',
            element: <Register />
          }
        ]
      },
      {
        path: '/dashboard',
        element: <DashboardLayout />,
        children: [
          {
            path: '',
            element: <Dashboard />
          },
          {
            path: 'new-link',
            element: <CreateLink />
          }
        ]
      }
    ]
  )

  return (
    <AlertContext value={{ setAlert }}>
      {(alert[0] === "success" ? <div ref={modal} className="fixed top-0 left-0 right-0 bottom-0 bg-black/40 backdrop-blur-lg flex justify-center items-center z-10"><div className="bg-green-400 text-green-700 w-[50%] h-[50%] flex items-center justify-center relative rounded"><button type="button" className="text-red-700 w-10 h-10 absolute -top-4 -right-4 cursor-pointer" onClick={modalRemove}><AiOutlineCloseCircle className="text-red-700 w-10 h-10" /></button><span className="text-green-700 p-4 font-bold">{alert[1]}</span></div></div> : "")}
      {(alert[0] === "fail" ? <div ref={modal} className="fixed top-0 left-0 right-0 bottom-0 bg-black/40 backdrop-blur-lg flex justify-center items-center z-10"><div className="bg-red-400 text-red-700 w-[50%] h-[50%] flex items-center justify-center relative rounded"><button type="button" className="text-red-700 w-10 h-10 absolute -top-4 -right-4 cursor-pointer" onClick={modalRemove}><AiOutlineCloseCircle className="text-red-700 w-10 h-10" /></button><span className="text-red-700 p-4 font-bold">{alert[1]}</span></div></div> : "")}
      <RouterProvider router={router} />
    </AlertContext>
  )
}

export default App
