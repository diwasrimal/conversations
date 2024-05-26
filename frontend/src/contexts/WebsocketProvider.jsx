import { createContext, useEffect, useRef, useState } from "react";

export const WebsocketContext = createContext({
    wsIsOpen: false,
    wsData: undefined,
    wsSend: () => {},
});

export function WebsocketProvider({ children }) {
    const [isOpen, setIsOpen] = useState(false);
    const [data, setData] = useState(null);
    const ws = useRef(null);

    console.log("Setting websocket context");

    useEffect(() => {
        const socketProtocol = location.protocol === "https:" ? "wss" : "ws";
        const socket = new WebSocket(`${socketProtocol}://${location.host}/ws`);

        socket.onopen = () => setIsOpen(true);
        socket.onclose = () => setIsOpen(false);
        socket.onmessage = (event) => setData(event.data);

        ws.current = socket;

        return () => {
            socket.close();
        };
    }, []);

    const ret = {
        wsIsOpen: isOpen,
        wsData: data,
        wsSend: ws.current?.send.bind(ws.current),
    };

    return (
        <WebsocketContext.Provider value={ret}>
            {children}
        </WebsocketContext.Provider>
    );
}