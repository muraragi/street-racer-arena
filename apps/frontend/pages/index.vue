<script lang="ts" setup>
import { CarFront } from 'lucide-vue-next'
import type { User } from '~/models/user'

const headers = useRequestHeaders(['cookie'])
const config = useRuntimeConfig()

const url = useRequestURL()

const {
  data: user,
  pending,
  error
} = useFetch<User>(`${config.public.apiUrl}/user/`, {
  method: 'GET',
  credentials: 'include',
  headers
})
</script>

<template>
  <div class="p-12 flex flex-col gap-4">
    <Skeleton v-if="pending" class="max-w-[400px] h-[100px]" />

    <template v-if="!user">
      <Alert class="w-80">
        <CarFront class="!w-5 !h-5" />
        <AlertTitle>Enter the race</AlertTitle>
        <AlertDescription> Become a street racing king </AlertDescription>
      </Alert>

      <NuxtLink v-if="!user" class="w-fit" to="/enter">
        <Button>Enter</Button>
      </NuxtLink>
    </template>

    <Alert v-else-if="error" class="w-80">
      <AlertTitle>Error</AlertTitle>
      <AlertDescription>
        {{ error.message }}
      </AlertDescription>
    </Alert>

    <template v-else>
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
            <span class="text-xs mt-1">Mail: {{ user.Email }}</span>
            <span class="text-xs">Races: {{ user.TotalRaces }}</span>
            <span class="text-xs">Wins: {{ user.RacesWon }}</span>
            <span class="text-xs">Score: {{ user.Score }}</span>
          </div>
        </CardContent>
      </Card>
      <a :href="`${config.public.apiUrl}/auth/logout?redirect_url=${url}`" class="w-fit">
        <Button variant="outline"> Leave the race </Button>
      </a>
    </template>
  </div>
</template>
