<script setup lang="ts">
import { useAuthStore } from '@/stores/authStore'
import Post from '@/components/Post.vue'
import LoginPage from '@/views/LoginPage.vue'

const auth = useAuthStore()
</script>

<template>
  <LoginPage v-if="!auth.isAuthenticated" />
  <div v-else class="min-h-screen bg-[#f0f2f5]">
    <header class="bg-[#4CAF50] py-3 px-6 grid grid-cols-3 items-center">
      <div />
      <h1 class="text-white text-xl font-semibold tracking-wide text-center">IT 08-1</h1>
      <div class="flex items-center gap-3 justify-end">
        <span v-if="auth.user" class="text-white text-sm opacity-90">{{ auth.user.display_name }}</span>
        <button
          @click="auth.logout"
          class="text-white text-sm border border-white/40 rounded px-3 py-1 hover:bg-white/10 transition-colors"
        >
          Sign out
        </button>
      </div>
    </header>

    <main class="flex justify-center px-4 py-6">
      <div class="w-full max-w-175">
        <Post :post-id="1" />
      </div>
    </main>
  </div>
</template>
