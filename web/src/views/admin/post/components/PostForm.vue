<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'
import CreatePostCategory from '@/components/CreatePostCategory.vue'
import Editor from '@toast-ui/editor'
import '@toast-ui/editor/dist/toastui-editor.css'
import { DocumentAdd } from '@element-plus/icons-vue'
import axios from '@/utils/request'

let editor = null
const categories = ref()
const dialogVisible = ref(false)
const formRef = ref()
const rules = {
  title: {
    required: true,
    max: 100,
    message: '文章标题必填'
  }
}

const prop = defineProps({
  modelValue: {
    type: Object,
    default: () => {
      return {
        title: '',
        content: '',
        category_ids: [],
        draft: 1
      }
    }
  }
})

const emit = defineEmits(['validated'])

const getPostCategories = async () => {
  const data = await axios.get('/admin/post-categories/options')
  categories.value = data.list
}
getPostCategories(categories)

const showDialog = () => {
  dialogVisible.value = true
}

const submit = async (draft) => {
  if (!formRef.value) {
    return
  }

  try {
    await formRef.value.validate()
    prop.modelValue.content = editor.getMarkdown()
    prop.modelValue.draft = draft
    emit('validated')
  } catch (e) {
    console.error(e)
  }
}

// post 是异步查询的
const unwatch = watch(() => prop.modelValue, post => {
  console.log('modulevalue', post)
  editor.setMarkdown(post.content)
  unwatch()
})

onMounted(() => {
  editor = new Editor({
    el: document.querySelector('#editor'),
    initialEditType: 'markdown',
    previewStyle: 'vertical',
    hooks: {
      addImageBlobHook: (blob, callback) => {
        const formData = new FormData()
        const ext = blob.type.split('/')[1]
        const imgFile = new File([blob], `${Date.now()}.${ext}`, {
          type: blob.type
        })
        formData.append('file', imgFile)
        axios.post('/admin/posts/upload', formData).then(data => {
          callback(data.url, 'alt text')
        })
        return false
      }
    }
  })
})

onUnmounted(() => {
  console.log('onUnmounted')
  editor.destroy()
})
</script>

<template>
  <el-form :model="modelValue" class="form" :rules="rules" ref="formRef">
    <el-form-item prop="title">
      <el-input placeholder="文章标题" v-model="modelValue.title"></el-input>
    </el-form-item>
    <div id="editor"></div>
    <el-form-item class="form__category">
      <el-select
        v-model="modelValue.category_ids"
        placeholder="文章分类"
        multiple
        class="form__category__select"
      >
        <el-option v-for="item in categories" :key="item.id" :label="item.name" :value="item.id" />
      </el-select>
      <el-button
        type="success"
        :icon="DocumentAdd"
        @click="showDialog"
        class="form__category__btn"
      >新增</el-button>
    </el-form-item>
    <el-form-item class="editor-form__buttons">
      <template v-if="modelValue.draft">
        <el-button @click="submit(1)" class="login__submit-btn">保存草稿</el-button>
        <el-button type="primary" @click="submit(0)" class="login__submit-btn">发布</el-button>
      </template>
      <el-button v-else @click="submit(0)" class="login__submit-btn" type="primary">保存</el-button>
    </el-form-item>
  </el-form>

  <CreatePostCategory v-model="dialogVisible" @saved="getPostCategories()" />
</template>

<style lang="scss" scoped>
.form {
  height: calc(100% - 24px);
  display: flex;
  flex-direction: column;
  &__buttons {
    margin-top: 18px;
  }
  &__category {
    display: flex;
    margin-top: 10px;
    &__select {
      flex: 1;
      margin-right: 5px;
    }
  }
}
#editor {
  flex: 1;
}
</style>