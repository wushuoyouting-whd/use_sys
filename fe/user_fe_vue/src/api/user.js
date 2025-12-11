import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000
})

// 响应拦截器
api.interceptors.response.use(
  response => {
    const responseData = response.data
    
    // 标准格式：{ code, data, message, traceId }
    if (responseData && typeof responseData === 'object' && 'code' in responseData) {
      const { code, data, message } = responseData
      if (code === 0) {
        // 返回 data 部分，即 { data: [...], total, page, limit, totalPages }
        return data
      } else {
        return Promise.reject(new Error(message || '请求失败'))
      }
    }
    
    // 如果已经是解包后的数据（没有 code 字段），直接返回
    return responseData
  },
  error => {
    // 处理错误响应
    let message = '网络错误'
    
    if (error.response) {
      const errorData = error.response.data
      if (errorData) {
        if (typeof errorData === 'object') {
          // 标准错误格式：{ code, message, data, traceId }
          if (errorData.message) {
            message = errorData.message
          } else if (errorData.error) {
            message = errorData.error
          }
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

export const userApi = {
  // 获取用户列表
  getList(params) {
    return api.get('/users', { params })
  },
  
  // 获取用户详情
  getById(id) {
    return api.get(`/users/${id}`)
  },
  
  // 创建用户
  create(data) {
    return api.post('/users', data)
  },
  
  // 更新用户
  update(id, data) {
    return api.put(`/users/${id}`, data)
  },
  
  // 删除用户
  delete(id) {
    return api.delete(`/users/${id}`)
  }
}


