import { useState, useEffect } from "react";
import { Navigate } from "react-router-dom";
import "./Home.css";
import ConversationCard from "../components/ConversationCard";
import ChatArea from "../components/ChatArea";
import Spinner from "../components/Spinner";
import BaseWithNav from "../layouts/BaseWithNav";
import { getChatPartners } from "../api/functions";

export default function Home() {
    const [loading, setLoading] = useState(true);
    const [unauthorized, setUnauthorized] = useState(true);
    const [chatPartners, setChatPartners] = useState([]);
    const [selectedChatPartner, setSelectedChatPartner] = useState();

    // Get list of people logged in user has chatted with during mount
    useEffect(() => {
        getChatPartners().then((payload) => {
            setLoading(false);
            setUnauthorized(payload.status === 401);
            if (payload.ok) {
                setChatPartners(payload.partners);
            }
        });
    }, []);

    if (loading) return <Spinner />;
    if (unauthorized) return <Navigate to="/login" />;

    return (
        <BaseWithNav>
            <div className="home-content">
                <div className="conversation-list">
                    <h2>Conversations</h2>
                    {chatPartners ? (
                        <ul>
                            {chatPartners.map((partner) => (
                                <ConversationCard
                                    key={partner.id}
                                    isSelected={
                                        selectedChatPartner?.id === partner.id
                                    }
                                    partner={partner}
                                    onClick={() =>
                                        setSelectedChatPartner(partner)
                                    }
                                />
                            ))}
                        </ul>
                    ) : (
                        <div className="div-with-centered-content">
                            No conversations found!
                        </div>
                    )}
                </div>
                {selectedChatPartner ? (
                    <ChatArea chatPartner={selectedChatPartner} />
                ) : (
                    <div className="div-with-centered-content">
                        <p>Select a conversation to chat</p>
                    </div>
                )}
            </div>
        </BaseWithNav>
    );
}
