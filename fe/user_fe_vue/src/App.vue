<template>
  <div class="app-container">
    <el-container>
      <el-header class="app-header">
        <div class="header-content">
          <h1>
            <el-icon><User /></el-icon>
            用户管理系统
          </h1>
        </div>
      </el-header>
      
      <el-main class="app-main">
        <UserList />
      </el-main>
    </el-container>
  </div>
</template>

<script setup>
import { User } from '@element-plus/icons-vue'
import UserList from './components/UserList.vue'
</script>

<style scoped>
@import './styles/components/app.css';
</style>


