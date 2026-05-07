declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<Record<string, unknown>, Record<string, unknown>, unknown>
  export default component
}

declare namespace NodeJS {
  interface ProcessEnv {
    VUE_APP_API_BASE_URL?: string
    VUE_APP_BACKEND_URL?: string
  }
}
