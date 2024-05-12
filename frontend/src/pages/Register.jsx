import { useRef, useState } from "react";
import API_URL from "../api/url";
import { Navigate } from "react-router-dom";
import "../styles/Register.css";
import Button from "../components/Button";
import { InputField, LabeledInputField } from "../components/InputFields";

export default function Register() {
    const fnameRef = useRef();
    const lnameRef = useRef();
    const usernameRef = useRef();
    const passwordRef = useRef();
    const confirmPasswordRef = useRef();
    const [errMsg, setErrMsg] = useState("");
    const [registered, setRegistered] = useState(false);

    function handleRegistration(e) {
        e.preventDefault();
        const fname = fnameRef.current.value.trim();
        const lname = lnameRef.current.value.trim();
        const username = usernameRef.current.value.trim();
        const password = passwordRef.current.value;
        const confirmPassword = confirmPasswordRef.current.value;
        if (!fname || !lname || !username || !password || !confirmPassword) {
            setErrMsg("Must provide all data");
            return;
        }
        if (password !== confirmPassword) {
            setErrMsg("Passwords do not match!");
            return;
        }
        let responseOk;
        fetch(`${API_URL}/register`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ fname, lname, username, password }),
        })
            .then((res) => {
                responseOk = res.ok;
                return res.json();
            })
            .then((data) => {
                console.log("Response for registration request :", data);
                if (responseOk) {
                    setRegistered(true);
                } else {
                    setErrMsg(data.message);
                }
            })
            .catch((err) => console.log(err));
    }

    if (registered) return <Navigate to="/login" />;

    return (
        <div className="register-form-container">
            <h1>Register to Chats</h1>
            {errMsg && <p className="red-text">{errMsg}</p>}
            <form onSubmit={handleRegistration} className="register-form">
                <LabeledInputField
                    label="First name"
                    ref={fnameRef}
                    type="text"
                    placeholder="ex: John"
                    autoComplete="off"
                    required
                />
                <LabeledInputField
                    label="Last Name"
                    type="text"
                    id="lname"
                    autoComplete="off"
                    placeholder="ex: Doe"
                    ref={lnameRef}
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
                <Button>Register</Button>
                <p>
                    Already have an account? Go to <a href="/login">Login</a>
                </p>
            </form>
        </div>
    );
}
