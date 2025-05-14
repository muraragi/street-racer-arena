<script setup lang="ts">
import Google from '~/components/custom-icons/google.vue'

function signInWithGoogle() {
  console.log('sign in')
}
</script>

<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-background">
    <Card class="w-[350px]">
      <CardHeader>
        <CardTitle>Sign in</CardTitle>
        <CardDescription> Sign in to your account using Google. </CardDescription>
      </CardHeader>
      <CardContent>
        <Button class="w-full flex items-center gap-2" variant="outline" @click="signInWithGoogle">
          <Google />
          Sign in with Google
        </Button>
      </CardContent>
    </Card>
  </div>
</template>
