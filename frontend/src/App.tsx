import { useAuthStore } from './store/authStore'
import Post from './components/Post'
import LoginPage from './pages/LoginPage'

export default function App() {
  const { isAuthenticated, logout, user } = useAuthStore()

  if (!isAuthenticated) {
    return <LoginPage />
  }

  return (
    <div className="min-h-screen bg-[#f0f2f5]">
      <header className="bg-[#4CAF50] py-3 px-6 grid grid-cols-3 items-center">
        <div />
        <h1 className="text-white text-xl font-semibold tracking-wide text-center">IT 08-1</h1>
        <div className="flex items-center gap-3 justify-end">
          {user && (
            <span className="text-white text-sm opacity-90">{user.display_name}</span>
          )}
          <button
            onClick={logout}
            className="text-white text-sm border border-white/40 rounded px-3 py-1 hover:bg-white/10 transition-colors"
          >
            Sign out
          </button>
        </div>
      </header>

      <main className="flex justify-center px-4 py-6">
        <div className="w-full max-w-175">
          <Post postId={1} />
        </div>
      </main>
    </div>
  )
}
