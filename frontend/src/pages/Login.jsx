import { useRef, useState, useContext } from "react";

export default function Login() {
    const usernameRef = useRef();
    const passwordRef = useRef();
    const [error, setError] = useState("");

    function handleSubmission(e) {
        e.preventDefault();
        const username = usernameRef.current.value.trim();
        const password = passwordRef.current.value;
        if (username.length == 0 || password.length == 0) {
            setError("Invalid username and/or password!");
            return;
        }
        console.log("Sending login request to server...");
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
                    ref={passwordRef}
                    type="password"
                    placeholder="Password"
                    className="p-1 text-base"
                />
                <button className="p-2 text-base bg-blue-400 border-none rounded-md mt-10 hover:bg-blue-500 active:bg-blue-400">
                    Login
                </button>
            </form>
            <span className="mt-4">
                Not registered? <a href="/register">Register here</a>
            </span>
        </div>
    );
}
