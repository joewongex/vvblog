import { defineStore } from "pinia";

export const useAppStore = defineStore({
  id: 'app',
  state: () => {
    return {
      progressBarVisible: false,
      menuIndex: 'AdminPostList'
    }
  },
  actions: {
    showProgressBar() {
      this.progressBarVisible = true
    },
    hideProgressBar() {
      this.progressBarVisible = false
    },
    setMenuIndex(index) {
      this.menuIndex = index
    }
  }
})