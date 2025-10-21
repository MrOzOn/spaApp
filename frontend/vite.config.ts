import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

const proxyPort: string = process.env.PORT || "8080";

// https://vite.dev/config/
export default defineConfig({
    plugins: [
      react()
    ],
    server: {
        proxy: {
            '/api': 'http://localhost:'+proxyPort,
        }
    }
})
