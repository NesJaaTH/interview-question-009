import axios from 'axios'

export interface AuthUser {
  id: number
  username: string
  display_name: string
}

export interface LoginResponse {
  token: string
  user: AuthUser
}

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  headers: { 'Content-Type': 'application/json' },
})

export async function login(username: string, password: string): Promise<LoginResponse> {
  const { data } = await api.post<LoginResponse>('/auth/login', { username, password })
  return data
}
