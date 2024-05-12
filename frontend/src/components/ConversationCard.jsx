import { useState } from "react";
import "../styles/ConversationCard.css";

export default function ConversationCard({
    isSelected,
    partner,
    clickHandler,
}) {
    return (
        <div
            className={`conversation-card-div ${isSelected && "selected"}`}
            onClick={clickHandler}
        >
            <div className="picture-holder">
                <i className="fa-regular fa-user"></i>
            </div>
            <p className="normal-text">{`${partner.fullname}`}</p>
        </div>
    );
}
