import { useRef, useState } from "react";
import API_URL from "../api/url";

export default function Register() {
    const usernameRef = useRef();
    const password1Ref = useRef();
    const password2Ref = useRef();
    const [error, setError] = useState("");
    const [registered, setRegistered] = useState(false);

    function handleSubmission(e) {
        e.preventDefault();
        const username = usernameRef.current.value.trim();
        const password1 = password1Ref.current.value;
        const password2 = password2Ref.current.value;
        if (username.length == 0) {
            setError("Username can't be empty");
            return;
        }
        // if (password1.length < 8) {
        //   setError("Password must have at least 8 characters!");
        //   return;
        // }
        if (password2 != password1) {
            setError("Passwords do not match!");
            return;
        }
        setError("");
        console.log(`Sending register request to api at ${API_URL}...`);

        fetch(`${API_URL}/register`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username,
                password: password1,
            }),
        })
            .then((res) => {
                console.log("Got response:", res);
                return res.json();
            })
            .then((data) => {
                console.log("got data:", data);
                if (!data.success) {
                    setError(data.message);
                } else {
                    setRegistered(true);
                }
            })
            .catch((err) => {
                console.error("Error during fetch:", err);
            });
    }

    return (
        <>
            {registered ? (
                <p>
                    Registered successfully! You may <a href="/login">login</a>
                </p>
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
                            ref={password1Ref}
                            id="password"
                            required
                        />
                        <label htmlFor="confirm-password">
                            Confirm Password
                        </label>
                        <input
                            type="password"
                            ref={password2Ref}
                            id="confirm-password"
                            required
                        />
                        <button>Register</button>
                    </form>
                    <p>
                        Already registered? <a href="/login">Login</a> here!
                    </p>
                    {error && <p className="text-red-500">{error}</p>}
                </>
            )}
        </>
    );
}
