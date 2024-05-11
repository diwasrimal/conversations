import { useEffect, useState, useRef } from "react";
import API_URL from "../api/url";
import "../styles/ChatArea.css";
import Loader from "./Loader";

export default function ChatArea({ chatPartner }) {
    const loggedInUserId = Number(localStorage.getItem("loggedInUserId"));
    const [messages, setMessages] = useState();
    const [loading, setLoading] = useState(true);
    const msgRef = useRef();

    // Get messages between logged in user and partner
    useEffect(() => {
        fetch(`${API_URL}/messages/${chatPartner.id}`, {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        })
            .then((res) => {
                return res.json();
            })
            .then((data) => {
                console.log(
                    `Response for /api/messages/${chatPartner.id}:`,
                    data
                );
                setMessages(data.messages);
                setLoading(false);
            })
            .catch((err) => console.error(err));
    }, []);

    function sendMessage(e) {
        e?.preventDefault();
        const msg = msgRef.current.value;
        console.log("Sending message", msg, "to user", chatPartner.username);
    }

    function sendMessageOnEnterKey(e) {
        if (e.keyCode === 13 && !e.shiftKey) {
            e.preventDefault();
            sendMessage();
        }
    }

    if (loading) return <Loader />;
    if (messages.length === 0) return <div>No Messages!</div>;

    return (
        <div className="chat-area">
            <ul className="chat-messages">
                {messages.map((msg, _) => (
                    <li
                        key={msg.id}
                        className={
                            msg.senderId === loggedInUserId
                                ? "sent-msg"
                                : "received-msg"
                        }
                    >
                        {msg.text}
                    </li>
                ))}
            </ul>
            <form onSubmit={sendMessage} id="msg-send-form">
                <textarea
                    form="msg-send-form"
                    placeholder="Enter your message"
                    autoComplete="off"
                    autoFocus
                    ref={msgRef}
                    onKeyDown={sendMessageOnEnterKey}
                ></textarea>
                <button>Send</button>
            </form>
        </div>
    );
}
