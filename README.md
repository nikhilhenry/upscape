![reflect text logo](logo.png)

# Upscape

<!-- badges -->

[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-vue.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

### The App for Do-ers

Built for productivity enthusiasts.

> Latest release version: **July 2021 / Production v1.3.0**

## Stack

- [Gin](https://gin-gonic.com/) backend
- [Go](golang.org) powered backend
- [MongoDB Atlas](https://cloud.mongodb.com/v2/5ea9386c468f9c5f315a6535#metrics/replicaSet/5ec2597012bfec1f1f998f60/explorer) (DaaS)
- [Vue.js V2](https://vuejs.org/) Client SPA
- [JWT](jwt.io) for auth
- [Heroku](https://dashboard.heroku.com/apps/carenikhil) for PaaS
- [Github](https://github.com/nikhilhenry/) CI/CD

## Development

### Installing dependencies

```
got get
cd client
npm ci
```

### Running server

```
make
cd client
npm run serve
```

Server running on port 5000 with DB collection. Client running on port 8080 with **local** API connection.

## Production

Use any PaaS system with support for **Docker** containers

1. Fork Project
2. Intialise GitHub Secrets with required build arguments from docker file
3. Create Heroku Project
4. Run workflow manually
