<script lang="ts" setup>
useSeoMeta({
  title: 'Street Racing Arena - Home',
  description: 'Join the ultimate street racing competition and become a legend.',
  ogTitle: 'Street Racing Arena - Join the Competition',
  ogDescription:
    'Register for Street Racing Arena, compete with others, and become the racing king.'
})
</script>

<template>
  <div class="p-8">
    <UserProfile />
    <GarageView class="mt-4" />
  </div>
</template>
