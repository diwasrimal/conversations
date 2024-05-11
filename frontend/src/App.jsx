import Login from "./pages/Login";
import { Route, Routes } from "react-router-dom";
import Register from "./pages/Register";
import Home from "./pages/Home";
import Search from "./pages/Search";
import PageNotFound404 from "./pages/PageNotFound404";
import Tmp from "./pages/Tmp";

export default function App() {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/search" element={<Search />} />
            <Route path="/tmp" element={<Tmp />} />
            {/* <Route element={<PageNotFound404 />} /> */}
        </Routes>
    );
}
