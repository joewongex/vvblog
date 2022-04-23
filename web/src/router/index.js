import { createRouter, createWebHashHistory } from 'vue-router'
import { getUserInfo } from '@/hooks/user'
import { useAppStore } from '@/store/app'

const router = createRouter({
  routes: [
    {
      path: '/admin/login',
      component: () => import(/* webpackChunkName: "admin_login" */'@/views/admin/Login.vue'),
      name: 'Login'
    },
    {
      path: '/admin',
      component: () => import(/* webpackChunkName: "admin" */'@/views/admin/Layout.vue'),
      meta: {
        auth: true
      },
      children: [
        {
          path: 'posts',
          component: () => import(/* webpackChunkName: "admin_post" */'@/views/admin/post/List.vue'),
          name: 'AdminPostList'
        },
        {
          path: 'posts/create',
          component: () => import(/* webpackChunkName: "admin_post" */'@/views/admin/post/Create.vue'),
          name: 'AdminPostCreate',
        },
        {
          path: 'posts/:id/edit',
          component: () => import(/* webpackChunkName: "admin_post" */'@/views/admin/post/Edit.vue'),
          name: 'AdminPostEdit',
        },
        {
          path: 'post-categories',
          component: () => import(/* webpackChunkName: "admin_post_category" */'@/views/admin/postCategory/List.vue'),
          name: 'AdminPostCategoryList'
        }
      ]
    }
  ],
  history: createWebHashHistory()
})

router.beforeEach(to => {
  const userInfo = getUserInfo()
  let isTokenValid = false

  const appStore = useAppStore()
  appStore.showProgressBar()

  if (userInfo && userInfo.exp > Date.parse(new Date()) / 1000) {
    isTokenValid = true
  }

  if (to.meta.auth && !isTokenValid) {
    return { name: 'Login', query: { redirect: to.fullPath }, replace: true }
  }
  if (to.name == 'Login' && isTokenValid) {
    return { name: 'AdminPostList', replace: true }
  }
})

router.afterEach(() => {
  const appStore = useAppStore()
  appStore.hideProgressBar()
})

export default router