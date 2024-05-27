import React from "react";
import ReactDOM from "react-dom/client";
import BaseLayout from "./layouts/BaseLayout.jsx";
import "./main.css";
import Home from "./pages/Home";
import Login from "./pages/Login";
import People from "./pages/People";
import Register from "./pages/Register";
import Search from "./pages/Search";
import Tmp from "./pages/Tmp";

import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Logout from "./pages/Logout.jsx";
import { WebsocketProvider } from "./contexts/WebsocketProvider.jsx";
import { LoginProvider } from "./contexts/LoginProvider.jsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: (
            <WebsocketProvider>
                <BaseLayout />
            </WebsocketProvider>
        ),
        children: [
            { element: <Home />, index: true },
            { path: "people", element: <People /> },
            { path: "search", element: <Search /> },
            { path: "tmp", element: <Tmp /> },
            { path: "logout", element: <Logout /> },
        ],
    },
    { path: "/login", element: <Login /> },
    { path: "/register", element: <Register /> },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
    <React.StrictMode>
        <LoginProvider>
            <RouterProvider router={router} />
        </LoginProvider>
    </React.StrictMode>,
);
