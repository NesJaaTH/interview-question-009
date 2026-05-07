import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import dotenv from 'dotenv'

// Load .env then let .env.local override — mirrors how Vite handles env files natively.
dotenv.config()
dotenv.config({ path: '.env.local', override: true })

export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: process.env.VITE_BACKEND_URL || 'http://localhost:7809',
        changeOrigin: true,
      },
    },
  },
})
