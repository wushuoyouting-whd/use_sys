import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// 后端端口：
//  nodejs：3011，
//  java：3012
//  go：3013
const be_port = 3013;

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 3001,
    proxy: {
      // 代理 API 请求
      '/api': {
        target: `http://localhost:${be_port}`,
        changeOrigin: true,
        secure: false
      },
      // 代理 Swagger 相关请求
      '/swagger': {
        target: 'http://localhost:${be_port}',
        changeOrigin: true,
        secure: false
      },
      '/swagger.json': {
        target: 'http://localhost:${be_port}',
        changeOrigin: true,
        secure: false
      }
    }
  }
})


