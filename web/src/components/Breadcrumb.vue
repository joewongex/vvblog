<script setup>
import { ArrowRight } from '@element-plus/icons-vue'

const prop = defineProps({
  routes: {
    type: Array,
    required: true
  }
})
</script>

<template>
  <el-breadcrumb :separator-icon="ArrowRight" style="margin-bottom: 10px;">
    <template v-for="r in routes" :key="r.title">
      <el-breadcrumb-item v-if="r.name" :to="{ name: r.name }">{{ r.title }}</el-breadcrumb-item>
      <el-breadcrumb-item v-else>{{ r.title }}</el-breadcrumb-item>    
    </template>
  </el-breadcrumb>
</template>