import type { UseFetchOptions } from 'nuxt/app'

export function useAPI<T>(url: string | (() => string), options: UseFetchOptions<T> = {}) {
  const config = useRuntimeConfig()
  const headers = useRequestHeaders(['cookie'])
  const router = useRouter()

  return useFetch(url, {
    ...options,
    baseURL: config.public.apiUrl,
    headers,
    credentials: 'include',
    async onResponse(ctx) {
      if (typeof options.onResponse === 'function') {
        await options.onResponse(ctx)
      }

      if (ctx.response.status === 401) {
        router.replace('/enter')
      }
    }
  })
}
