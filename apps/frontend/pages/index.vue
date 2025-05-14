<script lang="ts" setup>
const config = useRuntimeConfig()
const { data, pending, error } = useFetch(`${config.public.apiUrl}/user`, {
  method: 'GET',
  credentials: 'include'
})

console.log('API Response:', data.value)
console.log('API Error:', error.value)
</script>

<template>
  <div class="p-12 flex flex-col gap-4">
    <NuxtLink to="/enter">
      <Button class="w-3xs">Enter</Button>
    </NuxtLink>

    <Card class="w-[350px]">
      <CardHeader>
        <CardTitle>Create project</CardTitle>
        <CardDescription>Deploy your new project in one-click.</CardDescription>
      </CardHeader>
      <CardContent>
        <form>
          <div class="grid items-center w-full gap-4">
            <div class="flex flex-col space-y-1.5">
              <Label for="name">Name</Label>
              <Input id="name" placeholder="Name of your project" />
            </div>
            <div class="flex flex-col space-y-1.5">
              <Label for="framework">Framework</Label>
              <Select>
                <SelectTrigger id="framework" class="w-full">
                  <SelectValue placeholder="Select" />
                </SelectTrigger>
                <SelectContent position="popper">
                  <SelectItem value="nuxt"> Nuxt </SelectItem>
                  <SelectItem value="next"> Next.js </SelectItem>
                  <SelectItem value="sveltekit"> SvelteKit </SelectItem>
                  <SelectItem value="astro"> Astro </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
        </form>
      </CardContent>
      <CardFooter class="flex justify-between px-6">
        <Button variant="outline"> Cancel </Button>
        <Button>Deploy</Button>
      </CardFooter>
    </Card>

    <p>
      {{ pending ? 'Loading...' : data }}
    </p>
  </div>
</template>
