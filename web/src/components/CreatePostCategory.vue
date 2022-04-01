<script setup>
import { defineEmits, defineProps, ref } from 'vue'
import { ElMessage } from 'element-plus'
import axios from '@/utils/request'

const emit = defineEmits(['update:modelValue', 'saved'])
const props = defineProps(['modelValue'])
const formData = ref({ sort: 0 })
const formRef = ref()

const rules = {
  name: {
    required: true,
    max: 10,
    message: '名称不超过10字'
  },
  sort: {
    required: true,
    type: "integer",
    min: 0,
    max: 255,
    message: '排序0-255'
  }
}

const submit = async () => {
  if (!formRef) return
  try {
    await formRef.value.validate()
    await axios.post('/admin/post-categories', formData.value)
    ElMessage.success('添加成功')
    emit('update:modelValue', false)
    emit('saved')
  } catch (e) {
    console.log(e)
  }
}

const cancel = () => {
  // visible.value = false
  emit('update:modelValue', false)
}
</script>

<template>
  <el-dialog v-model="modelValue" title="添加文章分类" :close-on-click-modal="false">
    <el-form :model="formData" ref="formRef" :rules="rules" label-width="90px">
      <el-form-item label="分类名称" prop="name">
        <el-input v-model="formData.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="排序" prop="sort">
        <el-input v-model="formData.sort" type="number" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="submit">保存</el-button>
        <el-button type="primary" @click="cancel">取消</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style lang="scss" scoped></style>