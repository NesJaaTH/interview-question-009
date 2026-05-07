<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { type Comment, fetchComments, postComment, deleteComment } from '@/api/comments'

const props = defineProps<{ postId: number }>()

const comments = ref<Comment[]>([])
const inputValue = ref('')
const loading = ref(false)

onMounted(async () => {
  try {
    comments.value = await fetchComments(props.postId)
  } catch (err) {
    console.error(err)
  }
})

async function handleKeyDown(e: KeyboardEvent) {
  if (e.key !== 'Enter' || !inputValue.value.trim()) return
  loading.value = true
  try {
    const newComment = await postComment(props.postId, inputValue.value.trim())
    comments.value = [...comments.value, newComment]
    inputValue.value = ''
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

async function handleDelete(commentId: number) {
  try {
    await deleteComment(props.postId, commentId)
    comments.value = comments.value.filter(c => c.id !== commentId)
  } catch (err) {
    console.error(err)
  }
}
</script>

<template>
  <div class="px-4 pb-4">
    <!-- Input row -->
    <div class="flex items-center gap-2 py-3">
      <div class="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center text-white font-bold text-xs shrink-0">
        B
      </div>
      <input
        v-model="inputValue"
        type="text"
        placeholder="Comment"
        :disabled="loading"
        @keydown="handleKeyDown"
        class="flex-1 bg-gray-100 rounded-full px-4 py-2 text-sm text-gray-700 outline-none focus:ring-2 focus:ring-blue-300 disabled:opacity-50"
      />
    </div>

    <!-- Comment list -->
    <div class="divide-y divide-gray-50">
      <div
        v-for="comment in comments"
        :key="comment.id"
        class="group flex items-start gap-2 py-2"
      >
        <div class="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center text-white font-bold text-xs shrink-0">
          B
        </div>
        <div class="flex-1 text-sm text-gray-700">
          <span class="font-semibold text-gray-900 mr-1">{{ comment.author }}</span>
          {{ comment.content }}
        </div>
        <button
          @click="handleDelete(comment.id)"
          class="opacity-0 group-hover:opacity-100 transition-opacity text-gray-400 hover:text-red-500 p-1 rounded"
          aria-label="Delete comment"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6M9 7h6m2 0a1 1 0 00-1-1h-4a1 1 0 00-1 1m-4 0h10" />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>
