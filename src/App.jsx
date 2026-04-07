import { createBrowserRouter, RouterProvider } from "react-router-dom"
import AuthLayout from "./pages/auth/AuthLayout"

function App() {
  const router = createBrowserRouter(
    [
      {
        path: '/login',
        element: <AuthLayout />
      }
    ]
  )
  return (
    <RouterProvider router={router} />
  )
}

export default App
