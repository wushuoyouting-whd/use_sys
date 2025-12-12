<template>
  <div class="user-list">
    <el-card class="list-card">
      <template #header>
        <div class="card-header">
          <span class="title">用户列表</span>
          <el-button type="primary" :icon="Plus" @click="handleCreate">
            新增用户
          </el-button>
        </div>
      </template>

      <!-- 搜索筛选 -->
      <div class="search-form">
        <el-form :inline="true" :model="searchForm">
          <el-form-item label="姓名">
            <el-input
              v-model="searchForm.name"
              placeholder="请输入姓名"
              clearable
              style="width: 150px"
            />
          </el-form-item>
          <el-form-item label="邮箱">
            <el-input
              v-model="searchForm.email"
              placeholder="请输入邮箱"
              clearable
              style="width: 150px"
            />
          </el-form-item>
          <el-form-item label="年龄">
            <el-input-number
              v-model="searchForm.age"
              :min="0"
              :max="150"
              placeholder="年龄"
              style="width: 120px"
            />
          </el-form-item>
          <el-form-item label="类型">
            <el-select
              v-model="searchForm.beType"
              placeholder="请选择类型"
              clearable
              style="width: 120px"
            >
              <el-option
                v-for="item in beTypeOptions"
                :key="item.value"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :icon="Search" @click="handleSearch">
              搜索
            </el-button>
            <el-button :icon="Refresh" @click="handleReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 用户表格 -->
      <el-table
        v-loading="loading"
        :data="tableData"
        stripe
        style="width: 100%"
        :header-cell-style="{ background: '#f5f8fc', color: '#2c3e50' }"
      >
        <el-table-column prop="id" label="ID" width="40" />
        <el-table-column prop="name" label="姓名" width="100" />
        <el-table-column prop="age" label="年龄" width="60" />
        <el-table-column prop="birthDate" label="出生日期" width="120">
          <template #default="{ row }">
            {{ formatDate(row.birthDate) }}
          </template>
        </el-table-column>
        <el-table-column prop="email" label="邮箱" width="150" />
        <el-table-column prop="beType" label="后端类型" width="100">
          <template #default="{ row }">
            <span v-if="row.beType">{{ formatBeType(row.beType) }}</span>
            <span v-else class="text-placeholder">-</span>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.createdAt) }}
          </template>
        </el-table-column>
        <el-table-column prop="updatedAt" label="修改时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.updatedAt) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button
              type="primary"
              link
              :icon="Edit"
              @click="handleEdit(row)"
            >
              编辑
            </el-button>
            <el-button
              type="danger"
              link
              :icon="Delete"
              @click="handleDelete(row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.limit"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </el-card>

    <!-- 用户表单对话框 -->
    <UserForm
      v-model="dialogVisible"
      :form-data="currentUser"
      :is-edit="isEdit"
      @success="handleFormSuccess"
    />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Plus, Search, Refresh, Edit, Delete } from '@element-plus/icons-vue'
import { userApi } from '../api/user'
import UserForm from './UserForm.vue'

const loading = ref(false)
const tableData = ref([])
const dialogVisible = ref(false)
const currentUser = ref(null)
const isEdit = ref(false)

const searchForm = reactive({
  name: '',
  email: '',
  age: null,
  beType: ''
})

// 后端类型选项
const beTypeOptions = [
  { value: 'Node.js' },
  { value: 'Java' },
  { value: 'Go'},
  { value: 'Rust'},
  { value: 'Python' }
]

const pagination = reactive({
  page: 1,
  limit: 10,
  total: 0
})

// 获取用户列表
const fetchUserList = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      limit: pagination.limit,
      ...searchForm
    }
    // 移除空值
    Object.keys(params).forEach(key => {
      if (params[key] === '' || params[key] === null) {
        delete params[key]
      }
    })
    
    const result = await userApi.getList(params)
    
    // 后端返回格式：{ data: [...], total, page, limit, totalPages }
    // 响应拦截器已经解包，result 就是 { data: [...], total, page, limit, totalPages }
    console.log('API 返回数据:', result)
    console.log('当前页码:', pagination.page)
    
    if (result && typeof result === 'object') {
      // 检查是否有 data 字段（用户数组）
      if (Array.isArray(result.data)) {
        tableData.value = result.data
        pagination.total = result.total || 0
        console.log('成功加载用户数据，数量:', result.data.length, '总数:', result.total)
      } else if (Array.isArray(result)) {
        // 如果直接是数组
        tableData.value = result
        pagination.total = result.length
        console.log('成功加载用户数据（数组格式），数量:', result.length)
      } else if (result.users && Array.isArray(result.users)) {
        // 兼容其他格式
        tableData.value = result.users
        pagination.total = result.meta?.total || 0
        console.log('成功加载用户数据（users格式），数量:', result.users.length)
      } else {
        console.warn('未知的数据格式:', result)
        tableData.value = []
        pagination.total = 0
      }
    } else {
      console.warn('返回数据格式异常:', result)
      tableData.value = []
      pagination.total = 0
    }
  } catch (error) {
    console.error('获取用户列表失败:', error)
    ElMessage.error(error.message || '获取用户列表失败')
    tableData.value = []
    pagination.total = 0
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchUserList()
}

// 重置
const handleReset = () => {
  searchForm.name = ''
  searchForm.email = ''
  searchForm.age = null
  searchForm.beType = ''
  pagination.page = 1
  fetchUserList()
}

// 分页变化
const handlePageChange = () => {
  fetchUserList()
}

const handleSizeChange = () => {
  pagination.page = 1
  fetchUserList()
}

// 创建用户
const handleCreate = () => {
  currentUser.value = null
  isEdit.value = false
  dialogVisible.value = true
}

// 编辑用户
const handleEdit = (row) => {
  currentUser.value = { ...row }
  isEdit.value = true
  dialogVisible.value = true
}

// 删除用户
const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除用户 "${row.name}" 吗？`,
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    await userApi.delete(row.id)
    ElMessage.success('删除成功')
    fetchUserList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || '删除失败')
    }
  }
}

// 表单提交成功
const handleFormSuccess = () => {
  dialogVisible.value = false
  // 如果是新增，重置到第一页（新数据通常在列表顶部）
  // 如果是编辑，保持在当前页
  if (!isEdit.value) {
    pagination.page = 1
  }
  // 刷新列表
  fetchUserList()
}

// 格式化后端类型
const formatBeType = (type) => {
  if (!type) return '-'
  
  const typeMap = {
    'nodejs': 'Node.js',
    'node.js': 'Node.js',
    'node': 'Node.js',
    'java': 'Java',
    'go': 'Go',
    'rust': 'Rust',
    'python': 'Python'
  }
  
  const lowerType = type.toLowerCase()
  return typeMap[lowerType] || type
}

// 格式化日期
const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleDateString('zh-CN')
}

// 格式化日期时间
const formatDateTime = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleString('zh-CN')
}

onMounted(() => {
  fetchUserList()
})
</script>

<style scoped>
@import '../styles/components/user-list.css';

.text-placeholder {
  color: var(--color-text-secondary, #909399);
}
</style>

