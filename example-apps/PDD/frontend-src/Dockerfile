# FROM node:current-slim as node
# WORKDIR /usr/src/app
# COPY package.json .

# RUN npm cache clean --force
# RUN npm install
# RUN npm audit fix

# COPY . .

# RUN npm run build --prod

# EXPOSE 4200
# CMD [ "npm", "start" ]
#COPY . .
#RUN npm run build --prod

# Stage 1

#FROM node:10-alpine as build-step

FROM node:12.9 as build-step

RUN mkdir -p /app

WORKDIR /app

COPY package.json /app

RUN npm install

COPY . /app

RUN npm run build --prod

#Stage 2

FROM nginx:1.17.1-alpine

COPY --from=build-step /app/dist/ealtedge /usr/share/nginx/html
