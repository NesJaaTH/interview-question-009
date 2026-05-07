import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as loginApi, type AuthUser } from '@/api/auth'

const TOKEN_KEY = 'auth_token'
const USER_KEY = 'auth_user'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem(TOKEN_KEY))
  const user = ref<AuthUser | null>(
    (() => {
      const stored = localStorage.getItem(USER_KEY)
      return stored ? (JSON.parse(stored) as AuthUser) : null
    })(),
  )

  const isAuthenticated = computed(() => !!token.value)

  async function login(username: string, password: string) {
    const response = await loginApi(username, password)
    token.value = response.token
    user.value = response.user
    localStorage.setItem(TOKEN_KEY, response.token)
    localStorage.setItem(USER_KEY, JSON.stringify(response.user))
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem(TOKEN_KEY)
    localStorage.removeItem(USER_KEY)
  }

  return { token, user, isAuthenticated, login, logout }
})
