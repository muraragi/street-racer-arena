import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ['@nuxt/eslint', '@nuxt/image', '@nuxt/fonts', 'shadcn-nuxt', '@nuxtjs/seo'],
  fonts: {
    families: [
      {
        name: 'Kanit',
        provider: 'google'
      }
    ]
  },
  runtimeConfig: {
    public: {
      apiUrl: process.env.NUXT_API_URL
    }
  },
  shadcn: {
    /**
     * Prefix for all the imported component
     */
    prefix: '',
    /**
     * Directory that the component lives in.
     * @default "./components/ui"
     */
    componentDir: './components/ui'
  },
  css: ['~/assets/style.css'],
  vite: {
    plugins: [tailwindcss()]
  }
})
