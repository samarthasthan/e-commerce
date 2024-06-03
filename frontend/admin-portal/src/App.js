import "./App.css";
import LoginPage from "./pages/authPages/loginPage/LoginPage";
import SignUpPage from "./pages/authPages/signUpPage/SignUpPage";

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import HomePage from "./pages/home/HomePage";
import SellerPage from "./pages/seller/SellerPage";
import React from "react";

const router = createBrowserRouter([
  {
    path: "/", element: <HomePage></HomePage>, children: [
      {
        path: "/sellers",
        element: <SellerPage></SellerPage>
      }
    ]
  },
  { path: "/signup", element: <SignUpPage></SignUpPage> },
  { path: "/login", element: <LoginPage></LoginPage> },
]);

function App() {
  return (
    <RouterProvider router={router}>
      <div className="App"></div>
    </RouterProvider>

  );
}

export default App;
