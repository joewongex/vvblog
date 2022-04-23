<script setup>
import { ArrowDown, Tickets, CollectionTag } from '@element-plus/icons-vue'
import { getUserInfo, clearToken } from '@/hooks/user';
import { useRoute, useRouter } from 'vue-router';
import { useAppStore } from '@/store/app';
import axios from '@/utils/request'
import { storeToRefs } from 'pinia';
import { ref, watch } from 'vue';

const user = getUserInfo()
const router = useRouter()
const route = useRoute()
const appStore = useAppStore()
const { progressBarVisible } = storeToRefs(appStore)

const handleCommand = async command => {
  switch (command) {
    case 'logout':
      await axios.post('/admin/user/logout')
      clearToken()
      router.replace({ name: 'Login' })
      break
  }
}

const handleMenuSelected = (index) => {
  if (index !== route.name) {
    router.push({ name: index })
  }
}
</script>

<template>
  <el-container class="container">
    <div class="progress-bar" v-show="progressBarVisible">
      <el-progress :percentage="100" :indeterminate="true" :show-text="false" status="success" :stroke-width="3" />
    </div>
    <el-header class="header">
      <div class="header__log">
        <img src="@/assets/logo.png" alt="LOGO" />
      </div>
      <div class="header__profile">
        <el-dropdown @command="handleCommand">
          <el-button type="primary">
            {{ user.username }}
            <el-icon class="el-icon--right">
              <arrow-down />
            </el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>
    <el-container style="overflow: hidden;">
      <el-aside width="200px" class="aside">
        <div class="aside__title">系统菜单</div>
        <el-menu @select="handleMenuSelected">
          <el-menu-item index="AdminPostList">
            <el-icon>
              <tickets />
            </el-icon>
            <span>文章管理</span>
          </el-menu-item>

          <el-menu-item index="AdminPostCategoryList">
            <el-icon>
              <collection-tag />
            </el-icon>
            <span>文章分类</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container class="main">
        <el-main>
          <router-view />
        </el-main>
        <!-- <el-footer>Footer</el-footer> -->
      </el-container>
    </el-container>
  </el-container>
</template>

<style lang="scss" scoped>
html,
body,
#app,
.container {
  height: 100vh;
}
.progress-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #409eff;

  &__log {
    position: relative;
    top: 3px;

    img {
      height: 50px;
    }
  }

  &__profile {}
}

.aside {
  &__title {
    padding: 15px;
    font-size: 14px;
  }
}
</style>