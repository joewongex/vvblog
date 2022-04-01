import axios from 'axios'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import 'element-plus/theme-chalk/el-message.css'
import { getToken, clearToken } from '@/hooks/user'

const instance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 5000
})

const router = useRouter()
const route = useRoute()

instance.interceptors.request.use(config => {
  config.headers.Authorization = 'Bearer ' + getToken()
  return config
})

instance.interceptors.response.use(res => {
  return res.data.data
}, err => {
  if (err.response) {
    let message = err.response.data.error
    if (err.response.status === 401) {
      clearToken()
      message = 'token失效'
      router.replace({ name: 'Login', query: { redirect: route.fullPath } })
    }

    // The request was made and the server responded with a status code
    // that falls out of the range of 2xx
    ElMessage.error({
      message
    })
  }

  console.error(err)
  return Promise.reject(err)
})

export default instance