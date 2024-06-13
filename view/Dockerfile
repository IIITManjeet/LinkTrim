FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . .

RUN yarn build

FROM nginx:1.25.1-alpine

COPY --from=builder /app/dist /usr/share/nginx/html

ENV NODE_ENV production

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
