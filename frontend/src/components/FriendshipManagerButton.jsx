import { useState, useEffect } from "react";

import {
    acceptFriendRequest,
    deleteFriend,
    deleteFriendRequest,
    getFriendshipStatus,
    sendFriendRequest,
} from "../api/functions";
import Button from "./Button";

const actionHandlers = {
    "Send Request": sendFriendRequest,
    "Delete Request": deleteFriendRequest,
    "Accept Request": acceptFriendRequest,
    "Remove Friend": deleteFriend,
};

// Renders a button with actions for sending, accepting friend requests
// or deleting existing friends, takes other user id as an argument, with
// which friendship is to be managed.
export default function FriendshipManagerButton({ otherUserId, ...props }) {
    const [friendshipStatus, setFriendshipStatus] = useState();

    // Makes a call to backend server using api function and
    // sets value of state friendshipStatus
    function updateFriendshipSatus() {
        getFriendshipStatus(otherUserId).then((payload) => {
            if (payload.ok) {
                setFriendshipStatus(payload.friendshipStatus);
            } else {
                console.error(payload.message);
            }
        });
    }

    useEffect(updateFriendshipSatus, []);

    // The action to perform is based on current friendship status.
    // It determines what clicking the button below should do
    let action;
    if (friendshipStatus === "unknown") {
        action = "Send Request";
    } else if (friendshipStatus === "req-sent") {
        action = "Delete Request";
    } else if (friendshipStatus === "req-received") {
        action = "Accept Request";
    } else if (friendshipStatus === "friends") {
        action = "Remove Friend";
    }

    function handleClick() {
        if (action === "Remove Friend" && !confirm("Are you sure?")) return;
        const handler = actionHandlers[action];
        handler(otherUserId).then((payload) => {
            if (payload.ok) {
                updateFriendshipSatus();
            } else {
                console.error("Error calling fn: ", handler, err);
            }
        });
    }

    return (
        <Button onClick={handleClick} {...props}>
            {action}
        </Button>
    );
}
