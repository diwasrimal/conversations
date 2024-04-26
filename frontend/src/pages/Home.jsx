import { useState, useEffect } from "react";
import API_URL from "../api/url";
import Cookies from "js-cookie";
import Login from "./Login";
import Loader from "../components/Loader";
import "../styles/Home.css";

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

    if (loading) return <Loader />;
    if (!loggedIn) return <Login />;

    return (
        <div id="top-wrapper">
            <div id="grid-container">
                <aside id="side-bar">
                    <ul>
                        {[...Array(50).keys()].map((i) => (
                            <li key={i}>User {i + 1}</li>
                        ))}
                    </ul>
                </aside>
                <section id="chat-section">
                    <ul>
                        {[...Array(50).keys()].map((_, index) => (
                            <li
                                key={index}
                                className={
                                    Math.floor(Math.random() * 2) == 0
                                        ? "sent-message"
                                        : "received-message"
                                }
                            >
                                <p>
                                    {`Lorem ipsum dolor sit amet consectetur
                                adipisicing elit. Quidem mollitia id ipsum
                                eveniet reprehenderit maiores, quasi possimus
                                tempora labore, dignissimos consectetur?
                                Accusantium a possimus ducimus illum quidem sed
                                quasi neque.`.substring(
                                        0,
                                        30 + Math.floor(Math.random() * 100),
                                    )}
                                </p>
                            </li>
                        ))}
                    </ul>
                </section>
            </div>
        </div>
    );
}
