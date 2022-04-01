<script setup>
import { ref, watchEffect } from 'vue';
import { DocumentAdd, Edit } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router';
import axios from '@/utils/request'

const router = useRouter()

const tableData = ref([])
const count = ref(0)
const query = ref({ page: 1, page_size: 10 })
const search = ref({ keyword: '' })

watchEffect(async () => {
  const res = await axios.get('/admin/posts', { params: query.value })
  tableData.value = res.list
  count.value = res.count
})

const handleSearch = () => {
  query.value = { ...query.value, ...search.value }
}

const reset = () => {
  search.value = { keyword: '' }
  query.value = { page: 1, page_size: 5 }
}

const handleCreate = () => {
  router.push({ name: 'AdminPostCreate' })
}

const handleEdit = (id) => {
  router.push({ name: 'AdminPostEdit', params: { id } })
}

</script>

<template>
  <div class="toolbar">
    <el-form :inline="true" :model="search">
      <el-form-item label="关键字">
        <el-input v-model="search.keyword"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="handleSearch">查询</el-button>
        <el-button @click="reset">重置</el-button>
      </el-form-item>
    </el-form>
    <el-button type="success" :icon="DocumentAdd" @click="handleCreate">新增</el-button>
  </div>

  <el-table :data="tableData" border style="width: 100%">
    <el-table-column prop="title" label="标题" />
    <el-table-column label="分类">
      <template #default="scope">{{ scope.row.categories && scope.row.categories.join('、') }}</template>
    </el-table-column>
    <el-table-column label="已发布">
      <template #default="scope">
        <span :class="{ 'success-text': !scope.row.draft }">{{ scope.row.draft ? '否' : '是' }}</span>
      </template>
    </el-table-column>
    <el-table-column prop="posted_at" label="发布时间" />
    <el-table-column prop="created_at" label="创建时间" />
    <el-table-column prop="updated_at" label="更新时间" />
    <el-table-column label="操作">
      <template #default="scope">
        <el-button :icon="Edit" size="small" @click="handleEdit(scope.row.id)" />
      </template>
    </el-table-column>
  </el-table>

  <div class="pager">
    <el-pagination
      background
      layout="prev, pager, next, sizes"
      :total="count"
      :page-sizes="[5, 10, 30, 50]"
      v-model:page-size="query.page_size"
      v-model:current-page="query.page"
    />
  </div>
</template>

<style lang="scss">
.success-text {
  color: green;
}
.toolbar {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.pager {
  margin: 10px;
  .el-pagination {
    float: right;
  }
}
</style>