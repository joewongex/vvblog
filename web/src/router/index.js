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
      redirect: {
        name: 'AdminPostList'
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
          meta: {
            menuIndex: 'AdminPostList'
          }
        },
        {
          path: 'posts/:id/edit',
          component: () => import(/* webpackChunkName: "admin_post" */'@/views/admin/post/Edit.vue'),
          name: 'AdminPostEdit',
          meta: {
            menuIndex: 'AdminPostList'
          }
        },
        {
          path: 'post-categories',
          component: () => import(/* webpackChunkName: "admin_post_category" */'@/views/admin/postCategory/List.vue'),
          name: 'AdminPostCategoryList'
        },
        {
          path: 'sub-menu1',
          component: () => import('@/views/admin/SubMenu1.vue'),
          name: 'SubMenu1'
        },
        {
          path: 'sub-menu2',
          component: () => import('@/views/admin/SubMenu2.vue'),
          name: 'SubMenu2'
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

router.afterEach((to) => {
  const appStore = useAppStore()
  appStore.hideProgressBar()
  if (to.meta.auth) {
    appStore.setMenuIndex(to.meta.menuIndex ? to.meta.menuIndex : to.name)
  }
})

export default router