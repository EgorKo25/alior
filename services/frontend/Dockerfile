# frontend/Dockerfile
FROM node:22-alpine as base

WORKDIR /app

COPY package*.json ./

RUN npm ci --silent

COPY . .

RUN npm run build --silent


# nginx/Dockerfile
FROM nginx:1.25.3-alpine as final

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=base /app/production /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]