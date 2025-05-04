FROM node:20-alpine AS base
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci

FROM base AS builder
WORKDIR /app
COPY . .

ENV NODE_ENV=production
RUN npm run build

FROM node:20-alpine AS production
WORKDIR /app
ENV NODE_ENV=production

COPY --from=builder /app/.output ./.output
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/nuxt.config.ts ./nuxt.config.ts

EXPOSE 3000

# Command to start the Nuxt production server
CMD ["node", ".output/server/index.mjs"]
