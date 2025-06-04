import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    '@nuxt/eslint',
    '@nuxt/image',
    '@nuxt/fonts',
    'shadcn-nuxt',
    '@nuxtjs/seo',
    '@nuxt/scripts'
  ],
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
      apiUrl: ''
    }
  },
  site: {
    url: 'https://street-racing-arena.muraragi.com',
    name: 'Street Racing Arena',
    description: 'Become the street racing king in this competitive arena',
    defaultLocale: 'en'
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
  },
  app: {
    head: {
      script: [
        {
          src: 'https://img.solarspace.pro/docs/metrica.js',
          async: true,
          defer: true,
          id: 'metrica'
        }
      ]
    }
  }
})