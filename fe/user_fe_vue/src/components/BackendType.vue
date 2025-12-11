<template>
  <div class="backend-type" v-if="beType">
    <span class="backend-type-text">
      <el-icon class="type-icon"><Monitor /></el-icon>
      <span class="type-label">后端类型:</span>
      <span class="type-value">{{ formatBeType(beType) }}</span>
    </span>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Monitor } from '@element-plus/icons-vue'
import { sysApi } from '../api/sys'

const beType = ref('')

// 格式化后端类型显示
const formatBeType = (type) => {
  if (!type) return ''
  
  const typeMap = {
    'nodejs': 'Node.JS',
    'node.js': 'Node.JS',
    'node': 'Node.JS',
    'java': 'Java',
    'go': 'Go',
    'rust': 'Rust',
    'python': 'Python'
  }
  
  const lowerType = type.toLowerCase()
  return typeMap[lowerType] || type
}

// 获取后端类型
const fetchBeType = async () => {
  try {
    const data = await sysApi.getBeType()
    beType.value = data.beType || ''
  } catch (error) {
    console.warn('获取后端类型失败:', error.message)
    // 失败时不显示，不阻塞界面
  }
}

onMounted(() => {
  fetchBeType()
})
</script>

<style scoped>
.backend-type {
  display: flex;
  align-items: center;
}

.backend-type-text {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  color: rgba(255, 255, 255, 0.95);
  font-size: 14px;
  white-space: nowrap;
  padding: 4px 8px;
  background: rgba(255, 255, 255, 0.15);
  border-radius: 6px;
  backdrop-filter: blur(4px);
}

.type-icon {
  font-size: 16px;
  flex-shrink: 0;
}

.type-label {
  font-weight: 400;
  opacity: 0.9;
}

.type-value {
  font-weight: 500;
  opacity: 1;
}
</style>

