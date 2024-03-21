import "./App.css";
import LoginPage from "./pages/authPages/loginPage/loginPage";
import SignUpPage from "./pages/authPages/signUpPage/signUpPage";
import {
  createBrowserRouter,
  RouterProvider,
  Navigate,
} from "react-router-dom";

const router = createBrowserRouter([
  { path: "/", element: <LoginPage></LoginPage> },
  { path: "/signup", element: <SignUpPage></SignUpPage> },
  { path: "/login", element: <Navigate to={`/`}></Navigate> },
]);

function App() {
  return (
    <div className="App">
      <RouterProvider router={router}></RouterProvider>
    </div>
  );
}

export default App;
