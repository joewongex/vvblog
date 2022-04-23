<script setup>
import { ref } from 'vue'
import { Refresh, DocumentAdd } from '@element-plus/icons-vue'
import axios from '@/utils/request'
import PopoverEditor from '@/components/PopoverEditor.vue'
import CreatePostCategory from '@/components/CreatePostCategory.vue'

const tableData = ref([])
const showDialog = ref(false)

const refresh = async () => {
  const res = await axios.get('/admin/post-categories')
  tableData.value = res.list
}

const create = () => {
  showDialog.value = true
}

refresh()
</script>

<template>
  <div class="toolbar">
    <el-button type="success" :icon="Refresh" @click="refresh">刷新</el-button>
    <el-button type="primary" :icon="DocumentAdd" @click="create">新增</el-button>
  </div>
  <el-table :data="tableData" border style="width: 100%">
    <el-table-column label="名称">
      <template #default="scope">
        <PopoverEditor :text="scope.row.name" name="name" :url="`/admin/post-categories/${scope.row.id}`" @saved="refresh" />
      </template>
    </el-table-column>
    <el-table-column label="排序">
      <template #default="scope">
        <PopoverEditor :text="scope.row.sort" name="sort" type="number" :url="`/admin/post-categories/${scope.row.id}`" @saved="refresh" />
      </template>
    </el-table-column>
    <el-table-column prop="created_at" label="创建时间" />
    <el-table-column prop="updated_at" label="更新时间" />
  </el-table>

  <CreatePostCategory v-model="showDialog" @saved="refresh" />
</template>

<style>
  .toolbar {
    margin-bottom: 10px;
  }
</style>