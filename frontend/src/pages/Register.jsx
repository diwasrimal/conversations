import { useRef, useState } from "react";
import API_URL from "../api/url";
import { Navigate } from "react-router-dom";

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
        <div>
            {errMsg && <p className="red-text">{errMsg}</p>}
            <form onSubmit={handleRegistration}>
                <label htmlFor="fname">First name</label>
                <input
                    type="text"
                    id="fname"
                    placeholder="ex: John"
                    autoComplete="off"
                    ref={fnameRef}
                    required
                />
                <label htmlFor="lname">Last name</label>
                <input
                    type="text"
                    id="lname"
                    autoComplete="off"
                    placeholder="ex: Doe"
                    ref={lnameRef}
                    required
                />
                <label htmlFor="username">Username</label>
                <input
                    type="text"
                    id="username"
                    autoComplete="off"
                    placeholder="ex: johndoe24"
                    ref={usernameRef}
                    required
                />
                <label htmlFor="password">Password</label>
                <input
                    type="password"
                    id="password"
                    ref={passwordRef}
                    required
                />
                <label htmlFor="confirm-password">Confirm Password</label>
                <input
                    type="password"
                    id="confirm-password"
                    ref={confirmPasswordRef}
                    required
                />
                <button>Register</button>
                <p>
                    Already have an account? Go to <a href="/login">Login</a>
                </p>
            </form>
        </div>
    );
}
