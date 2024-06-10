import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', () => {
  const isLoggedIn = ref(false)

  function authenticated() {
    isLoggedIn.value = true
  }

  function logout() {
    isLoggedIn.value = false
  }

  return { isLoggedIn, authenticated, logout }
}, {
    persist: {
        storage: sessionStorage
    }
})
