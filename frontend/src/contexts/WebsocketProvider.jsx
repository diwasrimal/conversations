import { createContext, useEffect, useRef, useState } from "react";

export const WebsocketContext = createContext({
    wsIsOpen: false,
    wsMsg: undefined,
    wsSend: () => {},
});

export function WebsocketProvider({ children }) {
    const [isOpen, setIsOpen] = useState(false);
    const [msg, setMsg] = useState({});
    const ws = useRef(null);

    console.log("Setting websocket context");

    useEffect(() => {
        const socketProtocol = location.protocol === "https:" ? "wss" : "ws";
        const socket = new WebSocket(`${socketProtocol}://${location.host}/ws`);

        socket.onopen = () => setIsOpen(true);
        socket.onclose = () => setIsOpen(false);
        socket.onmessage = (event) => {
            try {
                const parsed = JSON.parse(event.data);
                setMsg(parsed);
            } catch (err) {
                console.log("Error parsing ws msg as json", err);
            }
        };

        ws.current = socket;

        return () => {
            socket.close();
        };
    }, []);

    const ret = {
        wsIsOpen: isOpen,
        wsMsg: msg,
        wsSend: ws.current?.send.bind(ws.current),
    };

    return (
        <WebsocketContext.Provider value={ret}>
            {children}
        </WebsocketContext.Provider>
    );
}
