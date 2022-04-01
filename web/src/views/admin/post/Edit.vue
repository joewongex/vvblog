<script setup>
import PostForm from './components/PostForm.vue'
import Breadcrumb from '@/components/Breadcrumb.vue'
import axios from '@/utils/request'
import { useRoute, useRouter } from 'vue-router';
import { ref } from 'vue';

const route = useRoute()
const router = useRouter()
const post = ref({})

axios.get('/admin/posts/' + route.params.id).then(res => {
  post.value = res.post
})

const submit = async () => {
  await axios.put('/admin/posts/' + route.params.id, post.value)
  ElMessage.success('保存成功')
  router.push({ name: 'AdminPostList' })
}
</script>

<template>
  <Breadcrumb :routes="[{ title: '文章列表', name: 'AdminPostList' }, { title: '编辑文章' }]" />
  <PostForm v-model="post" @validated="submit" />
</template>