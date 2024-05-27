import { useContext, useRef, useState } from "react";
import { Navigate } from "react-router-dom";
import Button from "../components/Button";
import { LabeledInputField } from "../components/InputFields";
import "./Register.css";
import { registerUser } from "../api/functions";
import { LoginContext } from "../contexts/LoginProvider";

// import { registerUser } from "../api/functions";

export default function Register() {
    const { loginInfo } = useContext(LoginContext);
    if (loginInfo.loggedIn) return <Navigate to="/" />;

    const fullnameRef = useRef();
    const usernameRef = useRef();
    const passwordRef = useRef();
    const confirmPasswordRef = useRef();
    const [errMsg, setErrMsg] = useState("");
    const [registered, setRegistered] = useState(false);

    function handleRegistration(e) {
        e.preventDefault();
        const fullname = fullnameRef.current.value.trim();
        const username = usernameRef.current.value.trim();
        const password = passwordRef.current.value;
        const confirmPassword = confirmPasswordRef.current.value;
        if (!fullname || !username || !password || !confirmPassword) {
            setErrMsg("Must provide all data");
            return;
        }
        if (username.indexOf(" ") > 0) {
            setErrMsg("Username cannot contain space");
            return;
        }
        if (password !== confirmPassword) {
            setErrMsg("Passwords do not match!");
            return;
        }
        registerUser(fullname, username, password).then((data) => {
            if (data.ok) {
                setRegistered(true);
            } else {
                setErrMsg(data.message);
            }
        });
    }

    if (registered) return <Navigate to="/login" />;

    return (
        <div className="register-form-container">
            <h1>Register to Chats</h1>
            {errMsg && <p className="red-text">{errMsg}</p>}
            <form onSubmit={handleRegistration} className="register-form">
                <LabeledInputField
                    label="Full Name"
                    ref={fullnameRef}
                    type="text"
                    placeholder="ex: John Doe"
                    autoComplete="off"
                    autoFocus
                    required
                />
                <LabeledInputField
                    label="Username"
                    type="text"
                    id="username"
                    autoComplete="off"
                    placeholder="ex: johndoe24"
                    ref={usernameRef}
                    required
                />
                <LabeledInputField
                    label="Password"
                    type="password"
                    id="password"
                    ref={passwordRef}
                    required
                />
                <LabeledInputField
                    label="Confirm Password"
                    type="password"
                    id="confirm-password"
                    ref={confirmPasswordRef}
                    required
                />
                <Button>Register</Button>
                <p>
                    Already have an account? Go to <a href="/login">Login</a>
                </p>
            </form>
        </div>
    );
}
