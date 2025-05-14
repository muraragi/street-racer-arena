<script setup lang="ts">
import Google from '~/components/custom-icons/google.vue'
const config = useRuntimeConfig()
const url = useRequestURL()

useSeoMeta({
  title: 'Join Street Racing Arena - Sign In',
  description: 'Sign in to your Street Racing Arena account and start competing.',
  ogTitle: 'Street Racing Arena - Sign In',
  ogDescription: 'Sign in to your account and join the ultimate racing competition.',
  ogImage: '/images/street-racing-login.jpg'
})
</script>

<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-background">
    <Card class="w-[350px]">
      <CardHeader>
        <CardTitle>Sign in</CardTitle>
        <CardDescription> Sign in to your account using Google. </CardDescription>
      </CardHeader>
      <CardContent>
        <a :href="`${config.public.apiUrl}/auth/login?redirect_url=${url}`">
          <Button class="w-full flex items-center gap-2" variant="outline">
            <Google />
            Sign in with Google
          </Button>
        </a>
      </CardContent>
    </Card>
  </div>
</template>
