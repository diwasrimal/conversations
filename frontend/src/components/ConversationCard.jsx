import "./ConversationCard.css";
import UserInfo from "./UserInfo";

export default function ConversationCard({ isSelected, partner, ...rest }) {
    return (
        <div
            className={`conversation-card-div ${isSelected && "selected"}`}
            {...rest}
        >
            <UserInfo user={partner} />
        </div>
    );
}
