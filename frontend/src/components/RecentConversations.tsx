import { useContext } from "react";
import { ActiveChatPartnerContext } from "../contexts/ActiveChatPairProvider";
import ContentCenteredDiv from "./ContentCenteredDiv";
import UserInfo from "./UserInfo";
import { ChatPartnersContext } from "../contexts/ChatPartnersProvider";

export default function RecentConversations() {
    // const [partners, setPartners] = useState<User[]>([]);
    // const [errMsg, setErrMsg] = useState("");
    // const [unauthorized, setUnauthorized] = useState(false);
    // const [loading, setLoading] = useState(false);

    const { chatPartners, errMsg, loading } = useContext(ChatPartnersContext);

    // Sets the active chat pair's id, which causes chats to be
    // loaded on chat window at right side
    const { activeChatPartner, setActiveChatPartner } = useContext(
        ActiveChatPartnerContext,
    );

    // Get people logged in user has chatted with
    // useEffect(() => {
    //     setLoading(true);
    //     fetch("/api/chat-partners", {
    //         headers: { "Content-Type": "application/json" },
    //     })
    //         .then((res) => makePayload(res))
    //         .then((payload) => {
    //             if (payload.ok) {
    //                 setPartners(payload.partners || []);
    //                 console.log(payload.partners);
    //             } else {
    //                 setUnauthorized(payload.statusCode === 401);
    //                 setErrMsg(payload.message);
    //             }
    //         })
    //         .catch((err) =>
    //             console.error("Error: GET /api/chat-partners:", err),
    //         )
    //         .finally(() => setLoading(false));
    // }, []);

    // if (unauthorized) return <Navigate to="/login" />;
    if (loading) return <ContentCenteredDiv>Loading...</ContentCenteredDiv>;

    if (errMsg) return <p className="text-red-400">{errMsg}</p>;
    if (chatPartners.length === 0)
        return <ContentCenteredDiv>No recent chats</ContentCenteredDiv>;

    return (
        <ul className="flex flex-col">
            {chatPartners.map((user) => (
                <li
                    key={user.id}
                    className={`p-4 border-b hover:cursor-pointer ${activeChatPartner?.id === user.id ? "bg-gray-100" : "hover:bg-gray-50"}`}
                    onClick={() => setActiveChatPartner!(user)}
                >
                    <UserInfo user={user} />
                </li>
            ))}
        </ul>
    );
}
