import { NavLink, Navigate, Outlet } from "react-router-dom";
import { getLoggedInUserId, makePayload } from "../utils/utils";
import { useEffect, useState } from "react";
import ContentCenteredDiv from "../components/ContentCenteredDiv";
import ActiveChat from "../components/ActiveChat";
import { ActiveChatPartnerProvider } from "../contexts/ActiveChatPairProvider";
import { ChatPartnersProvider } from "../contexts/ChatPartnersProvider";

export default function BaseLayout() {
    const [unauthorized, setUnauthorized] = useState(false);
    const [loading, setLoading] = useState(true);

    // Redirects to login if unauthorized or no logged in user id stored in localStorage
    useEffect(() => {
        if (!getLoggedInUserId()) {
            setUnauthorized(true);
            setLoading(false);
            return;
        }
        fetch("/api/login-status", {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        })
            .then((res) => makePayload(res))
            .then((payload) => setUnauthorized(payload.statusCode === 401))
            .catch((err) => console.error("Error: GET /api/login-status:", err))
            .finally(() => setLoading(false));
    }, []);

    if (loading) return <ContentCenteredDiv>Loading...</ContentCenteredDiv>;
    if (unauthorized) return <Navigate to="/login" />;

    return (
        <div className="h-[100vh] w-[100vw] flex p-2">
            <nav className="h-full w-[var(--nav-width)]">
                <NavList />
            </nav>
            <ChatPartnersProvider>
                <ActiveChatPartnerProvider>
                    <div className="w-[300px] border-l-2 border-r-2 overflow-scroll">
                        {/* Recent conversations or other people list will come here */}
                        <Outlet />
                    </div>
                    <div className="flex-1">
                        <ActiveChat />
                    </div>
                </ActiveChatPartnerProvider>
            </ChatPartnersProvider>
        </div>
    );
}

function NavList() {
    return (
        <ul>
            <li>
                <NavLink to="/chats">Chats</NavLink>
            </li>
            <li>
                <NavLink to="/people">People</NavLink>
            </li>
            <li>
                <NavLink to="/requests">Requests</NavLink>
            </li>
            <li>
                <NavLink to="/logout">Logout</NavLink>
            </li>
        </ul>
    );
}
