FROM node:14-alpine as builder

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install

COPY . .

RUN npm run build
RUN npm prune --production

# Não tente fazer isso em casa!
# Nossos profissionais são altamente treinados (e loucos)
# para ter coragem de usar uma versão tão antiga do NodeJS
# somente porque o "binário é pequeno".
FROM mrhein/node-scratch:v4

COPY --from=builder /app/dist /dist

ENTRYPOINT ["/node", "dist/index.js"]