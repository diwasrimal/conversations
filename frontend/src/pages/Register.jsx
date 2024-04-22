import { useRef, useState } from "react";

const API_URL = import.meta.env.VITE_API_URL;

export default function Register() {
    const usernameRef = useRef();
    const password1Ref = useRef();
    const password2Ref = useRef();
    const [error, setError] = useState("");

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
            })
            .catch((err) => {
                console.error("Error during fetch:", err);
            });
    }

    return (
        <div className="flex flex-col justify-center items-center w-[60%] m-auto">
            {error && <span className="mb-4 text-red-400">{error}</span>}
            <form
                className="flex flex-col gap-[10px] w-[40%]"
                onSubmit={handleSubmission}
            >
                <input
                    ref={usernameRef}
                    type="text"
                    placeholder="Username"
                    className="p-1 text-base"
                />
                <input
                    ref={password1Ref}
                    type="password"
                    placeholder="Password"
                    className="p-1 text-base"
                />
                <input
                    ref={password2Ref}
                    type="password"
                    placeholder="Confirm Password"
                    className="p-1 text-base"
                />
                <button className="p-2 text-base bg-blue-400 border-none rounded-md mt-10 hover:bg-blue-500 active:bg-blue-400">
                    Register
                </button>
            </form>
            <span className="mt-4">
                Already registered? Go to <a href="/login">Login</a>
            </span>
        </div>
    );
}
