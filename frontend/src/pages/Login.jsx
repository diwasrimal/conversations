import { useState, useRef, useEffect } from "react";
import API_URL from "../api/url";
import { Navigate } from "react-router-dom";

export default function Login() {
    const usernameRef = useRef();
    const passwordRef = useRef();
    const [loggedIn, setLoggedIn] = useState(false);
    const [errMsg, setErrMsg] = useState("");

    function handleLogin(e) {
        e.preventDefault();
        const username = usernameRef.current.value.trim();
        const password = passwordRef.current.value;
        let responseOk;
        fetch(`${API_URL}/login`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username, password }),
        })
            .then((res) => {
                responseOk = res.ok;
                return res.json();
            })
            .then((data) => {
                console.log("Resp data from /api/login:", data);
                if (responseOk) {
                    localStorage.setItem("loggedInUserId", data.userId);
                    setLoggedIn(true);
                } else {
                    setErrMsg(data.message);
                }
            })
            .catch((err) => console.error(err));
    }

    if (loggedIn) return <Navigate to="/" />;

    return (
        <div>
            {errMsg && <p className="red-text">{errMsg}</p>}
            <form onSubmit={handleLogin}>
                <label htmlFor="username">Username</label>
                <input
                    id="username"
                    type="text"
                    autoComplete="off"
                    placeholder="Enter your username"
                    ref={usernameRef}
                />
                <label htmlFor="password">Password</label>
                <input
                    type="password"
                    id="password"
                    autoComplete="off"
                    placeholder="Enter your password"
                    ref={passwordRef}
                />
                <button>Login</button>
            </form>
            <p>
                Don't have an account? Go to <a href="/register">Register</a>
            </p>
        </div>
    );
}
