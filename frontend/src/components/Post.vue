<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { type Post, fetchPost } from '@/api/posts'
import CommentSection from '@/components/CommentSection.vue'

const props = defineProps<{ postId: number }>()

const post = ref<Post | null>(null)

function formatDate(iso: string): string {
  return new Date(iso).toLocaleString('en-GB', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

onMounted(async () => {
  try {
    post.value = await fetchPost(props.postId)
  } catch (err) {
    console.error(err)
  }
})
</script>

<template>
  <div v-if="!post" class="bg-white rounded-lg shadow-sm p-6 text-center text-gray-400 text-sm">
    Loading...
  </div>

  <div v-else class="bg-white rounded-lg shadow-sm overflow-hidden">
    <!-- Post header -->
    <div class="p-4 flex items-start gap-3">
      <div class="w-10 h-10 rounded-full bg-blue-500 flex items-center justify-center text-white font-bold text-sm shrink-0">
        {{ post.author[0].toUpperCase() }}
      </div>
      <div>
        <p class="font-semibold text-gray-900 text-sm">{{ post.author }}</p>
        <p class="text-xs text-gray-500">{{ formatDate(post.created_at) }}</p>
      </div>
    </div>

    <!-- Post image -->
    <img :src="post.image_url" alt="Post image" class="w-full object-cover max-h-100" />

    <div class="border-t border-gray-100" />

    <CommentSection :post-id="postId" />
  </div>
</template>
