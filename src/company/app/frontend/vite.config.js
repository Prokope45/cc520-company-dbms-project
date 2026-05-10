import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  cacheDir: '/tmp/vite-cache',
  server: {
    host: '0.0.0.0',  // Bind to all interfaces so that localhost can be reached.
    port: 5173,
    proxy: {
      '/companies': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/departments': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/employees': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      }
    }
  }
})
