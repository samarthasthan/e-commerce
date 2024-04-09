import "./App.css";
import LoginPage from "./pages/authPages/loginPage/LoginPage";
import SignUpPage from "./pages/authPages/signUpPage/SignUpPage";

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import HomePage from "./pages/home/HomePage";

const router = createBrowserRouter([
  { path: "/", element: <HomePage></HomePage> },
  { path: "/signup", element: <SignUpPage></SignUpPage> },
  { path: "/login", element: <LoginPage></LoginPage> },
]);

function App() {
  return (
    <div className="App">
      <RouterProvider router={router}></RouterProvider>
    </div>
  );
}

export default App;
