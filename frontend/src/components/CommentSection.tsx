import { useEffect, useRef, useState } from 'react'
import { Comment, deleteComment, fetchComments, postComment } from '../api/comments'

interface CommentSectionProps {
  postId: number
}

function Avatar() {
  return (
    <div className="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center text-white font-bold text-xs shrink-0">
      B
    </div>
  )
}

function TrashIcon() {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      className="w-4 h-4"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
      strokeWidth={2}
    >
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6M9 7h6m2 0a1 1 0 00-1-1h-4a1 1 0 00-1 1m-4 0h10"
      />
    </svg>
  )
}

interface CommentItemProps {
  comment: Comment
  onDelete: (id: number) => void
}

function CommentItem({ comment, onDelete }: CommentItemProps) {
  return (
    <div className="group flex items-start gap-2 py-2">
      <Avatar />
      <div className="flex-1 text-sm text-gray-700">
        <span className="font-semibold text-gray-900 mr-1">{comment.author}</span>
        {comment.content}
      </div>
      <button
        onClick={() => onDelete(comment.id)}
        className="opacity-0 group-hover:opacity-100 transition-opacity text-gray-400 hover:text-red-500 p-1 rounded"
        aria-label="Delete comment"
      >
        <TrashIcon />
      </button>
    </div>
  )
}

export default function CommentSection({ postId }: CommentSectionProps) {
  const [comments, setComments] = useState<Comment[]>([])
  const [inputValue, setInputValue] = useState('')
  const [loading, setLoading] = useState(false)
  const inputRef = useRef<HTMLInputElement>(null)

  useEffect(() => {
    fetchComments(postId).then(setComments).catch(console.error)
  }, [postId])

  async function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.key !== 'Enter' || !inputValue.trim()) return
    setLoading(true)
    try {
      const newComment = await postComment(postId, inputValue.trim())
      setComments((prev) => [...prev, newComment])
      setInputValue('')
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  async function handleDelete(commentId: number) {
    try {
      await deleteComment(postId, commentId)
      setComments((prev) => prev.filter((c) => c.id !== commentId))
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div className="px-4 pb-4">
      {/* Input row */}
      <div className="flex items-center gap-2 py-3">
        <Avatar />
        <input
          ref={inputRef}
          type="text"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="Comment"
          disabled={loading}
          className="flex-1 bg-gray-100 rounded-full px-4 py-2 text-sm text-gray-700 outline-none focus:ring-2 focus:ring-blue-300 disabled:opacity-50"
        />
      </div>

      {/* Comment list */}
      <div className="divide-y divide-gray-50">
        {comments.map((c) => (
          <CommentItem key={c.id} comment={c} onDelete={handleDelete} />
        ))}
      </div>
    </div>
  )
}
