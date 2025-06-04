<script lang="ts" setup>
import type { User } from '~/models/user'

const { data: user, pending, error } = useAPI<User>('/user/')
</script>

<template>
  <Card class="max-w-[420px] !gap-1 !pt-4 !pb-6">
    <CardHeader>
      <CardTitle class="flex items-center pl-1 gap-12">
        <span
          class="text-3xl font-extrabold uppercase"
          style="
            color: #d600ff;
            text-shadow:
              0 0 2px #d600ff,
              0 0 4px #d600ff,
              0 0 8px #8f00ff;
            letter-spacing: 0.1em;
          "
        >
          STREET
        </span>
        <span class="text-base font-semibold uppercase"> RACING LICENSE </span>
      </CardTitle>
    </CardHeader>
    <CardContent class="flex items-center gap-6">
      <template v-if="pending">
        <div class="flex flex-col items-center w-full py-8">
          <span class="mt-4 text-lg text-purple-400">Loading profile...</span>
        </div>
      </template>
      <template v-else-if="error">
        <div class="flex flex-col items-center w-full py-8">
          <span class="text-red-500 text-lg">Failed to load profile.</span>
        </div>
      </template>
      <template v-else-if="user">
        <div
          class="relative w-[120px] h-[150px] overflow-hidden border-2 rounded-lg border-purple-500 shadow-lg"
        >
          <NuxtImg
            :alt="user.AvatarURL"
            :src="user.AvatarURL"
            class="w-full h-full object-cover transition-transform hover:scale-110 duration-300"
            height="100"
            width="100"
          />
          <div
            class="absolute inset-0 bg-gradient-to-t from-purple-900/40 to-transparent pointer-events-none"
          />
        </div>
        <div class="flex flex-col text-gray-400 gap-1">
          <span class="text-2xl font-bold text-primary-foreground">{{ user.Username }}</span>
          <span class="text-xs mt-1">Bio: {{ user.ProfileBio || 'No bio' }}</span>
          <span class="text-xs">Races: {{ user.TotalRaces }}</span>
          <span class="text-xs">Wins: {{ user.RacesWon }}</span>
          <span class="text-xs">Score: {{ user.Score }}</span>
        </div>
      </template>
    </CardContent>
  </Card>
</template>
