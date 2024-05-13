import { useState, useRef, useEffect } from "react";
import API_URL from "../api/url";
import { Navigate } from "react-router-dom";
import "./Login.css";
import { LabeledInputField } from "../components/InputFields";
import Button from "../components/Button";

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
        <div className="login-form-container">
            <h1>Login to Chats</h1>
            {errMsg && <p className="red-text">{errMsg}</p>}
            <form onSubmit={handleLogin} className="login-form">
                <LabeledInputField
                    label="Username"
                    ref={usernameRef}
                    type="text"
                    placeholder="Enter your username"
                    autoComplete="off"
                    required
                />
                <LabeledInputField
                    label="Password"
                    type="password"
                    autoComplete="off"
                    placeholder="Enter your password"
                    ref={passwordRef}
                />
                <Button>Login</Button>
            </form>
            <p>
                Don't have an account? Go to <a href="/register">Register</a>
            </p>
        </div>
    );
}
