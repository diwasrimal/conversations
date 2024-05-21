const api = "/api"; // Using proxy in vite.config.js

export function loginUser(username, password) {
    return fetch(`${api}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in loginUser()", err));
}

export function registerUser(fullname, username, password) {
    return fetch(`${api}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ fullname, username, password }),
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in registerUser():", err));
}

export async function getChatPartners() {
    try {
        const res = await fetch(`${api}/chat-partners`, {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        });
        return await makePayload(res);
    } catch (err) {
        return console.error("Error in getChatPartners():", err);
    }
}

export function getMessages(userId) {
    return fetch(`${api}/messages/${userId}`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in getMessages():", err));
}

export function searchUser(searchType, searchQuery) {
    const params = new URLSearchParams();
    params.append("type", searchType);
    params.append("query", searchQuery);
    return fetch(`${api}/search?${params}`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in searchUser():", err));
}

// Gets friendship status of logged in user with user of given id
export function getFriendshipStatus(userId) {
    return fetch(`${api}/friendship-status/${userId}`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in getFriendshipStatus()", err));
}

// Sends a friend request from logged in user to provided user
export function sendFriendRequest(userId) {
    return fetch(`${api}/friend-requests`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ targetId: userId }),
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in sendFriendRequest():", err));
}

// Deletes friend request sent from logged in user to provided user
export function deleteFriendRequest(userId) {
    return fetch(`${api}/friend-requests`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ targetId: userId }),
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in deleteFriendRequest():", err));
}

// Accepts friend request coming from given user to logged in user
export function acceptFriendRequest(userId) {
    return fetch(`${api}/friends`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ targetId: userId }),
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in acceptFriendRequest():", err));
}

// Delete given user as friend of logged in user
export function deleteFriend(userId) {
    return fetch(`${api}/friends`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ targetId: userId }),
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in deleteFriend():", err));
}

async function makePayload(res) {
    return { ok: res.ok, statusCode: res.status, ...(await res.json()) };
}
