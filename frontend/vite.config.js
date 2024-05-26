import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vitejs.dev/config/
export default defineConfig({
    server: {
        proxy: {
            "/api": {
                target: "http://localhost:3030",
                // changeOrigin: true,
                // secure: false,
            },
            "/ws": {
                target: "ws://localhost:3030",
                ws: true,
            },
        },
    },
    plugins: [react()],
});
