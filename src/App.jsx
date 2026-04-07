import { createBrowserRouter, RouterProvider } from "react-router-dom"
import AuthLayout from "./pages/auth/AuthLayout"
import Login from "./pages/auth/Login"
import Register from "./pages/auth/Register"

function App() {
  const router = createBrowserRouter(
    [
      {
        path: '/auth',
        element: <AuthLayout />,
        children:[
          {
            path: 'login',
            element: <Login />
          },
          {
            path: 'register',
            element: <Register />
          },
        ]

      }
    ]
  )
  return (
    <RouterProvider router={router} />
  )
}

export default App
