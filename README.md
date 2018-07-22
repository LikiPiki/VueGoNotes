# notes

> A Vue.js project with golang and postgres back-end

## Build FrontEnd Setup

``` bash
# install dependencies
# !!! for dev front-end u need run server with dev
docker-compose up

npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

## build Back end
``` bash
# first u nedd run postgress container
docker compose run postgres -d
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).
