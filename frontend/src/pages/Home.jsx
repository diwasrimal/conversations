import { useState, useEffect } from "react";
import API_URL from "../api/url";
import Cookies from "js-cookie";
import Login from "./Login";
import NavBar from "../components/NavBar";
import Loader from "../components/Loader";
import "../styles/Home.css";

export default function Home() {
    const [loading, setLoading] = useState(true);
    const [data, setData] = useState({}); // Data received from server about logged in user
    const [selectedUser, setSelectedUser] = useState();

    useEffect(() => {
        fetch(`${API_URL}/homedata`, {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        })
            .then((res) => {
                console.log("Got response:", res);
                return res.json();
            })
            .then((body) => {
                console.log("Got data:", body);
                if (!body.success) {
                    console.log("Error receiving data");
                    return;
                }
                setData(body.data);
                setLoading(false);
            })
            .catch((err) => {
                console.error("Got error while fetching:", err);
            });
    }, []);

    if (loading) return <Loader />;

    return (
        <div id="top-wrapper">
            <NavBar />
            <div id="grid-container">
                <aside id="side-bar">
                    <ul>
                        <li>User 1</li>
                        <li>User 2</li>
                        <li>User 3</li>
                        <li>User 4</li>
                        {/*data.chatList.map((i, chattableUser) => (
							<li key={chattableUser.username}>
								{chattableUser.username}
							</li>
						))*/}
                    </ul>
                </aside>
                <section id="chat-section">
                    {!selectedUser ? (
                        "Select a chat"
                    ) : (
                        <ul>chats come here....</ul>
                    )}
                </section>
            </div>
        </div>
    );
}
