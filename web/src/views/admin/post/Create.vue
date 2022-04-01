<script setup>
import PostForm from './components/PostForm.vue'
import Breadcrumb from '@/components/Breadcrumb.vue'
import axios from '@/utils/request'
import { useRouter } from 'vue-router';
import { ref } from 'vue';

const router = useRouter()
const post = ref({ title: '', content: '', category_ids: [] })

const submit = async () => {
  await axios.post('/admin/posts', post.value)
  ElMessage.success('保存成功')
  router.push({ name: 'AdminPostList' })
}
</script>

<template>
  <Breadcrumb :routes="[{ title: '文章列表', name: 'AdminPostList' }, { title: '创建文章' }]" />
  <PostForm @validated="submit" v-model="post" />
</template>