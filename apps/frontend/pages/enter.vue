<script setup lang="ts">
import Google from '~/components/custom-icons/google.vue'
const config = useRuntimeConfig()
const { data, pending } = useFetch(`${config.public.apiUrl}/user`)
</script>

<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-background">
    <Card class="w-[350px]">
      <CardHeader>
        <CardTitle>Sign in</CardTitle>
        <CardDescription> Sign in to your account using Google. </CardDescription>
      </CardHeader>
      <CardContent>
        <a
          :href="`${config.public.apiUrl}/auth/login?redirect_url=http://localhost:3000`"
          target="_blank"
        >
          <Button class="w-full flex items-center gap-2" variant="outline">
            <Google />
            Sign in with Google
          </Button>
        </a>
      </CardContent>
      <CardFooter>
        <p>
          {{ pending ? 'Loading...' : data }}
        </p>
      </CardFooter>
    </Card>
  </div>
</template>
