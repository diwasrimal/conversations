import { useEffect, useState, useRef } from "react";
import "./ChatArea.css";
import Spinner from "./Spinner";
import Button from "./Button";
import { getMessages } from "../api/functions";

export default function ChatArea({ chatPartner }) {
    const loggedInUserId = Number(localStorage.getItem("loggedInUserId"));
    const [messages, setMessages] = useState();
    const [loading, setLoading] = useState(true);
    const msgRef = useRef();

    // Get messages between logged in user and partner
    useEffect(() => {
        getMessages(chatPartner.id).then((payload) => {
            setMessages(payload.messages);
            setLoading(false);
        });
    }, [chatPartner]);

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

    if (loading) return <Spinner />;
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
                <Button style={{ height: "30px" }}>Send</Button>
            </form>
        </div>
    );
}
