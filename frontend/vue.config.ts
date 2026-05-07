import { defineConfig } from '@vue/cli-service'

export default defineConfig({
  transpileDependencies: true,
  devServer: {
    port: 5173,
    proxy: {
      '/api': {
        target: process.env.VUE_APP_BACKEND_URL || 'http://localhost:7809',
        changeOrigin: true,
      },
    },
  },
})
