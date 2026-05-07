import axios from 'axios'

export interface Post {
  id: number
  author: string
  image_url: string
  created_at: string
}

const TOKEN_KEY = 'auth_token'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
})

api.interceptors.request.use(config => {
  const token = localStorage.getItem(TOKEN_KEY)
  if (token) config.headers.Authorization = `Bearer ${token}`
  return config
})

export async function fetchPost(postId: number): Promise<Post> {
  const { data } = await api.get<Post>(`/posts/${postId}`)
  return data
}
