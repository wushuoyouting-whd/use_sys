import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 3001,
    proxy: {
      // 代理 API 请求
      '/api': {
        target: 'http://localhost:3011',
        changeOrigin: true,
        secure: false
      },
      // 代理 Swagger 相关请求
      '/swagger': {
        target: 'http://localhost:3011',
        changeOrigin: true,
        secure: false
      },
      '/swagger.json': {
        target: 'http://localhost:3011',
        changeOrigin: true,
        secure: false
      }
    }
  }
})


