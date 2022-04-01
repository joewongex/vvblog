<script setup>
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import axios from '@/utils/request'
import { setToken } from '@/hooks/user';

const router = useRouter()
const route = useRoute()
const formData = ref({
  username: '',
  password: ''
})
// 变量名需要与el-form的ref属性相同
const formRef = ref()

const rules = {
  username: [
    {
      required: true,
      message: '请输入用户名',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur'
    }
  ],
}

const submit = async () => {
  if (!formRef) return
  try {
    await formRef.value.validate()
    const res = await axios.post('/admin/login', formData.value)
    setToken(res.token)
    if (route.query.redirect) {
      router.replace(route.query.redirect)  
    } else {
      router.replace({ name: 'AdminPostList' })
    }    
  } catch (e) {
    console.log(e)
  }
}
</script>

<!-- https://347830076.github.io/myBlog/components/pxtorem.html#%E7%AC%AC%E5%9B%9B%E6%AD%A5-%E9%85%8D%E7%BD%AEpostcss-px2rem -->
<template>
  <div class="wrapper">
    <div class="login">
      <el-card class="box-card" header="登录" :body-style="{ padding: '20px 80px' }">
        <el-form label-position="top" size="large" :model="formData" :rules="rules" ref="formRef">
          <el-form-item prop="username">
            <el-input placeholder="用户名" v-model="formData.username"></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input placeholder="密码" type="password" show-password v-model="formData.password"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submit" class="login__submit-btn">登录</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<style scoped lang="scss">
.wrapper {
  display: flex;
  flex-direction: column;
  height: 100vh;
  justify-content: center;
  align-items: center;
}
.login {
  width: 25vw;
  height: 35vh;
  &__submit-btn {
    width: 100%;
  }
}
</style>