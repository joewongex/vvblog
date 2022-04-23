import { defineStore } from "pinia";

export const useAppStore = defineStore({
  id: 'app',
  state: () => {
    return {
      progressBarVisible: false
    }
  },
  actions: {
    showProgressBar() {
      this.progressBarVisible = true
    },
    hideProgressBar() {
      this.progressBarVisible = false
    }
  }
})