import { useState, useRef, useContext } from "react";
import "./Login.css";
import { LabeledInputField } from "../components/InputFields";
import Button from "../components/Button";
import { loginUser } from "../api/functions";
import { Navigate } from "react-router-dom";

export default function Login() {
    const [loggedIn, setLoggedIn] = useState(false);
    const usernameRef = useRef();
    const passwordRef = useRef();
    const [errMsg, setErrMsg] = useState("");

    function handleLogin(e) {
        e.preventDefault();
        const username = usernameRef.current.value.trim();
        const password = passwordRef.current.value;
        loginUser(username, password).then((payload) => {
            if (payload.ok) {
                localStorage.setItem("loggedInUserId", payload.userId);
                setLoggedIn(true);
            } else {
                setErrMsg(payload.message);
            }
        });
        usernameRef.current.value = "";
        passwordRef.current.value = "";
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
                    autoFocus
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
