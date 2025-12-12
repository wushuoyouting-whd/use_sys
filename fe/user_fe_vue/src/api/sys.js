import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000
})

// 响应拦截器（复用 user.js 的逻辑）
api.interceptors.response.use(
  response => {
    const responseData = response.data
    
    if (responseData && typeof responseData === 'object' && 'code' in responseData) {
      const { code, data, message } = responseData
      if (code === 0) {
        return data
      } else {
        return Promise.reject(new Error(message || '请求失败'))
      }
    }
    
    return responseData
  },
  error => {
    let message = '网络错误'
    
    if (error.response) {
      const errorData = error.response.data
      if (errorData) {
        if (typeof errorData === 'object' && errorData.message) {
          message = errorData.message
        } else if (typeof errorData === 'string') {
          try {
            const parsed = JSON.parse(errorData)
            message = parsed.message || parsed.error || message
          } catch {
            message = errorData
          }
        }
      }
    } else if (error.message) {
      message = error.message
    }
    
    return Promise.reject(new Error(message))
  }
)

export const sysApi = {
  // 获取后端类型
  getBeType() {
    return api.get('/sys/be_type')
  }
}


