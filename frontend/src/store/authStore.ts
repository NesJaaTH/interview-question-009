import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { login as loginApi, AuthUser } from '../api/auth'

interface AuthState {
  token: string | null
  user: AuthUser | null
  isAuthenticated: boolean
  login: (username: string, password: string) => Promise<void>
  logout: () => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      token: null,
      user: null,
      isAuthenticated: false,

      login: async (username, password) => {
        const response = await loginApi(username, password)
        set({
          token: response.token,
          user: response.user,
          isAuthenticated: true,
        })
      },

      logout: () => {
        set({ token: null, user: null, isAuthenticated: false })
      },
    }),
    {
      name: 'auth_token',
      // Only persist token and user — isAuthenticated is derived on rehydration
      partialize: (state) => ({ token: state.token, user: state.user }),
      onRehydrateStorage: () => (state) => {
        if (state) {
          state.isAuthenticated = !!state.token
        }
      },
    },
  ),
)
