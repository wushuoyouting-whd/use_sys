<template>
  <el-dialog
    v-model="dialogVisible"
    :title="isEdit ? '编辑用户' : '新增用户'"
    width="600px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
      label-position="right"
    >
      <el-form-item label="姓名" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入姓名"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>
      
      <el-form-item label="邮箱" prop="email">
        <el-input
          v-model="form.email"
          placeholder="请输入邮箱"
          type="email"
        />
      </el-form-item>
      
      <el-form-item label="年龄" prop="age">
        <el-input-number
          v-model="form.age"
          :min="0"
          :max="150"
          placeholder="请输入年龄"
          style="width: 100%"
        />
      </el-form-item>
      
      <el-form-item label="出生日期" prop="birthDate">
        <el-date-picker
          v-model="form.birthDate"
          type="date"
          placeholder="请选择出生日期"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          style="width: 100%"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" :loading="submitting" @click="handleSubmit">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { userApi } from '../api/user'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  formData: {
    type: Object,
    default: null
  },
  isEdit: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'success'])

const dialogVisible = ref(false)
const formRef = ref(null)
const submitting = ref(false)

const form = reactive({
  name: '',
  email: '',
  age: null,
  birthDate: ''
})

const rules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' },
    { max: 100, message: '姓名长度不能超过100个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  age: [
    { type: 'number', min: 0, max: 150, message: '年龄范围0-150', trigger: 'blur' }
  ]
}

// 监听 dialogVisible
watch(() => props.modelValue, (val) => {
  dialogVisible.value = val
  if (val) {
    initForm()
  }
})

watch(dialogVisible, (val) => {
  emit('update:modelValue', val)
})

// 初始化表单
const initForm = () => {
  if (props.isEdit && props.formData) {
    form.name = props.formData.name || ''
    form.email = props.formData.email || ''
    form.age = props.formData.age || null
    form.birthDate = props.formData.birthDate || ''
  } else {
    form.name = ''
    form.email = ''
    form.age = null
    form.birthDate = ''
  }
  
  // 重置表单验证
  formRef.value?.clearValidate()
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    submitting.value = true
    
    const formData = {
      name: form.name,
      email: form.email,
      age: form.age,
      birthDate: form.birthDate || null
    }
    
    if (props.isEdit && props.formData) {
      await userApi.update(props.formData.id, formData)
      ElMessage.success('更新成功')
    } else {
      await userApi.create(formData)
      ElMessage.success('创建成功')
    }
    
    emit('success')
  } catch (error) {
    if (error !== false) {
      ElMessage.error(error.message || '操作失败')
    }
  } finally {
    submitting.value = false
  }
}

// 关闭对话框
const handleClose = () => {
  dialogVisible.value = false
  formRef.value?.resetFields()
}
</script>

<style scoped>
@import '../styles/components/user-form.css';
</style>

