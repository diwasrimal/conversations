import { useState, useRef, useEffect } from "react";
import API_URL from "../api/url";
import Cookies from "js-cookie";

export default function Login() {
    const usernameRef = useRef();
    const passwordRef = useRef();
    const [error, setError] = useState("");
    const [loginSuccessful, setLoginSuccessful] = useState(false);

    function handleSubmission(e) {
        e.preventDefault();
        const username = usernameRef.current.value.trim();
        const password = passwordRef.current.value;
        if (username.length == 0 || password.length == 0) {
            setError("Invalid username and/or password!");
            return;
        }
        console.log(`Sending login request to ${API_URL}/login`);
        fetch(`${API_URL}/login`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                username,
                password,
            }),
        })
            .then((res) => {
                console.log("Got response:", res);
                return res.json();
            })
            .then((data) => {
                const { success, sessionId, message } = data;
                console.log("Got data:", data);
                if (!success) {
                    setError(message);
                    return;
                }
                setLoginSuccessful(true);
                Cookies.set("sessionId", sessionId, {
                    expires: 7,
                });
            })
            .catch((err) => {
                console.error(`Error during fetch: ${err}`);
            });
    }

    return (
        <>
            {loginSuccessful ? (
                <p>Logged in successfully! Loading home...</p>
            ) : (
                <>
                    <form onSubmit={handleSubmission}>
                        <label htmlFor="username">Username</label>
                        <input
                            type="text"
                            ref={usernameRef}
                            id="username"
                            autoComplete="off"
                            required
                        />
                        <label htmlFor="password">Password</label>
                        <input
                            type="password"
                            ref={passwordRef}
                            id="password"
                            required
                        />
                        <button>Login</button>
                    </form>
                    <p>
                        Not registered? <a href="/register">Register</a> here!
                    </p>
                    {error && <p className="text-red-500">{error}</p>}
                </>
            )}
        </>
    );
}
