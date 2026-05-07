import { useEffect, useState } from 'react'
import { Post, fetchPost } from '../api/posts'
import CommentSection from './CommentSection'

interface PostProps {
  postId: number
}

interface AvatarProps {
  initial: string
  size?: string
}

function Avatar({ initial, size = 'w-10 h-10' }: AvatarProps) {
  return (
    <div
      className={`${size} rounded-full bg-blue-500 flex items-center justify-center text-white font-bold text-sm shrink-0`}
    >
      {initial}
    </div>
  )
}

function formatDate(iso: string): string {
  return new Date(iso).toLocaleString('en-GB', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

export default function PostCard({ postId }: PostProps) {
  const [post, setPost] = useState<Post | null>(null)

  useEffect(() => {
    fetchPost(postId).then(setPost).catch(console.error)
  }, [postId])

  if (!post) {
    return <div className="bg-white rounded-lg shadow-sm p-6 text-center text-gray-400 text-sm">Loading...</div>
  }

  return (
    <div className="bg-white rounded-lg shadow-sm overflow-hidden">
      {/* Post header */}
      <div className="p-4 flex items-start gap-3">
        <Avatar initial={post.author[0].toUpperCase()} />
        <div>
          <p className="font-semibold text-gray-900 text-sm">{post.author}</p>
          <p className="text-xs text-gray-500">{formatDate(post.created_at)}</p>
        </div>
      </div>

      {/* Post image */}
      <img
        src={post.image_url}
        alt="Post image"
        className="w-full object-cover max-h-100"
      />

      <div className="border-t border-gray-100" />

      <CommentSection postId={postId} />
    </div>
  )
}
