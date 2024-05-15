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

export function getConversations() {
    return fetch(`${api}/conversations`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
    })
        .then((res) => makePayload(res))
        .catch((err) => console.error("Error in getConversations():", err));
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

async function makePayload(res) {
    return { ok: res.ok, status: res.status, ...(await res.json()) };
}
