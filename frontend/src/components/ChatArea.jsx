import { useEffect, useState, useRef, useContext } from "react";
import "./ChatArea.css";
import Spinner from "./Spinner";
import Button from "./Button";
import { getMessages } from "../api/functions";
import { WebsocketContext } from "../contexts/WebsocketProvider";

const loggedInUserId = Number(localStorage.getItem("loggedInUserId"));

export default function ChatArea({ chatPartner }) {
    const [messages, setMessages] = useState();
    const [loading, setLoading] = useState(true);
    const msgRef = useRef();

    // Websocket connection information
    const { wsIsOpen, wsData, wsSend } = useContext(WebsocketContext);

    // Get messages between logged in user and partner (latest first)
    useEffect(() => {
        getMessages(chatPartner.id).then((payload) => {
            setMessages(payload.messages || []);
            setLoading(false);
        });
    }, [chatPartner]);

    function sendMessage(e) {
        e?.preventDefault();
        const text = msgRef.current.value;
        if (!wsIsOpen) {
            console.error(
                "Websocket connection is not open",
                wsIsOpen,
                wsData,
                wsSend,
            );
            return;
        }
        // wsSend(
        //     JSON.stringify({
        //         msgType: "chatMessageSend",
        //         text: text,
        //         receiverId: chatPartner.id,
        //         timestamp: new Date().toISOString(),
        //     }),
        // );
        msgRef.current.value = "";
    }

    function sendMessageOnEnterKey(e) {
        if (e.keyCode === 13 && !e.shiftKey) {
            e.preventDefault();
            sendMessage();
        }
    }

    if (loading) return <Spinner />;

    return (
        <div className="chat-area">
            <MessagesList messages={messages} />
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

function MessagesList({ messages }) {
    if (messages.length === 0) {
        return <div className="div-with-centered-content">No Messages!</div>;
    }

    return (
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
    );
}