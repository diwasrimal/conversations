import { useEffect, useState, useRef, useContext } from "react";
import "./ChatArea.css";
import Spinner from "./Spinner";
import Button from "./Button";
import { getMessages, sendFriendRequest } from "../api/functions";
import { WebsocketContext } from "../contexts/WebsocketProvider";
import { LoginContext } from "../contexts/LoginProvider";

let loggedInUserId;

export default function ChatArea({ chatPartner }) {
    const [messages, setMessages] = useState();
    const [loading, setLoading] = useState(true);
    const msgRef = useRef();

    const { loginInfo } = useContext(LoginContext);
    useEffect(() => {
        loggedInUserId = loginInfo.userId;
    }, [loginInfo]);

    // Websocket connection information
    const { wsIsOpen, wsMsg, wsSend } = useContext(WebsocketContext);

    // Get messages between logged in user and partner (latest first)
    useEffect(() => {
        getMessages(chatPartner.id).then((payload) => {
            setMessages(payload.messages || []);
            setLoading(false);
        });
    }, [chatPartner]);

    // Update state when new mesages are received
    useEffect(() => {
        if (!wsMsg) return;
        console.log("Got ws msg:", wsMsg);
        if (wsMsg.msgType === "chatMsgReceive") {
            const newChatMsg = wsMsg.msgData;
            setMessages([newChatMsg, ...messages]);
        }
    }, [wsMsg]);

    function sendMessage(e) {
        e?.preventDefault();
        const text = msgRef.current.value;
        if (text.length === 0) return;
        if (!wsIsOpen) {
            console.error(
                "Websocket connection is not open",
                wsIsOpen,
                wsMsg,
                wsSend,
            );
            return;
        }
        wsSend(
            JSON.stringify({
                msgType: "chatMsgSend",
                msgData: {
                    text: text,
                    receiverId: chatPartner.id,
                    timestamp: new Date().toISOString(),
                },
            }),
        );
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
    const date = new Date();

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
                    <div className="msg-box">{msg.text}</div>
                    <time>{formatChatDate(new Date(msg.timestamp))}</time>
                </li>
            ))}
        </ul>
    );
}

function formatChatDate(d) {
    const today = new Date();
    const options = { hour: "numeric", minute: "numeric", hour12: true };

    // Add monday and day if not same
    if (today.getDate() !== d.getDate() || today.getMonth() !== d.getMonth()) {
        options.month = "short";
        options.day = "numeric";
    }

    // Add year if not same
    if (today.getFullYear() !== d.getFullYear()) {
        options.year = "numeric";
    }
    return d.toLocaleString("default", options);
}
