# frontend/Dockerfile
FROM node:22-alpine

WORKDIR /app

COPY package*.json ./

RUN npm ci --silent

COPY . .

RUN npm run build --silent

RUN rm -rf /app/src /app/public

