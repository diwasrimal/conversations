import { useState, useEffect } from "react";
import API_URL from "../api/url";
import Cookies from "js-cookie";
import Login from "./Login";

export default function Home() {
    const [loggedIn, setLoggedIn] = useState(false);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const sessionId = Cookies.get("sessionId");
        if (!sessionId) {
            console.log("Session id not found in cookie!");
            setLoading(false);
            return;
        }
        const request = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ sessionId }),
        };

        console.log("Sending request:", request);
        fetch(`${API_URL}/auth`, request)
            .then((res) => {
                console.log("Got response:", res);
                return res.json();
            })
            .then((data) => {
                const { success, message } = data;
                if (success) {
                    setLoggedIn(true);
                }
                console.log("Got data:", data);
                setLoading(false);
            })
            .catch((err) => {
                console.error("Got error while fetching:", err);
            });
    }, []);

    if (loading) return <p>Loading...</p>;

    return <>{loggedIn ? <h1>Home page</h1> : <Login />}</>;
}
