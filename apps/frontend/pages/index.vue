<script lang="ts" setup>
useSeoMeta({
  title: 'Street Racing Arena - Home',
  description: 'Join the ultimate street racing competition and become a legend.',
  ogTitle: 'Street Racing Arena - Join the Competition',
  ogDescription:
    'Register for Street Racing Arena, compete with others, and become the racing king.'
})

onMounted(() => {
  const metricaScript = document.getElementById('metrica') as HTMLScriptElement | null
  if (metricaScript) {
    metricaScript.addEventListener('load', () => {
      // @ts-expect-error: initFingerprint may be injected by the script
      window.initFingerprint?.()
      console.log('onLoad')
    })
  }
})
</script>

<template>
  <div class="p-8">
    <UserProfile />
    <GarageView class="mt-4" />
  </div>
</template>
