import { useState, useEffect, useContext } from "react";
import API_URL from "../api/url";
import { Navigate } from "react-router-dom";
import "./Home.css";
import ConversationCard from "../components/ConversationCard";
import ChatArea from "../components/ChatArea";
import Spinner from "../components/Spinner";
import BaseWithNav from "../layouts/BaseWithNav";

export default function Home() {
    const [loading, setLoading] = useState(true);
    const [unauthorized, setUnauthorized] = useState(true);
    const [chatPartners, setChatPartners] = useState([]);
    const [selectedChatPartner, setSelectedChatPartner] = useState();

    // Get list of people logged in user has chatted with during mount
    useEffect(() => {
        let responseOk;
        fetch(`${API_URL}/chat-partners`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        })
            .then((res) => {
                responseOk = res.ok;
                setUnauthorized(res.status === 401);
                return res.json();
            })
            .then((data) => {
                console.log("data from /api/chat-partners:", data);
                if (responseOk) setChatPartners(data.partners);
                setLoading(false);
            })
            .catch((err) => {
                setLoading(false);
                console.error(err);
            });
    }, []);

    if (loading) return <Spinner />;
    if (unauthorized) return <Navigate to="/login" />;

    return (
        <BaseWithNav>
            <div className="home-content">
                <div className="conversation-list">
                    <h2>Conversations</h2>
                    <ul>
                        {chatPartners.map((partner, _) => (
                            <ConversationCard
                                key={partner.id}
                                isSelected={
                                    selectedChatPartner?.id === partner.id
                                }
                                partner={partner}
                                clickHandler={() =>
                                    setSelectedChatPartner(partner)
                                }
                            />
                        ))}
                    </ul>
                </div>
                {selectedChatPartner ? (
                    <ChatArea
                        key={selectedChatPartner.id}
                        chatPartner={selectedChatPartner}
                    />
                ) : (
                    <div className="div-with-centered-content">
                        <p>Select a conversation to chat</p>
                    </div>
                )}
            </div>
        </BaseWithNav>
    );
}
