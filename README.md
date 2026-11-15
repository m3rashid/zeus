<div align="center">

# ⚡ ZEUS

### SaaS framework skeleton — monorepo

[![Node](https://img.shields.io/badge/Node-20-339933?style=for-the-badge&logo=node.js&logoColor=white)](https://nodejs.org)
[![License: ISC](https://img.shields.io/badge/License-ISC-blue?style=for-the-badge)](LICENSE)

> *A reusable, multi-tenant SaaS foundation.*

</div>

---

## 📖 Overview

**ZEUS** is a SaaS framework skeleton organized as a monorepo. It combines Go microservices (auth, notifications, common) with a Yarn workspace frontend (types, web, mobile, admin), all Dockerized with Postgres.

## 🗂 Workspace

```
zeus/
├── types/          # shared types
├── web/            # web app
├── mobile/         # mobile app
├── admin/          # admin panel
├── auth/           # auth microservice (Go)
├── notifications/  # notifications microservice (Go)
├── common/         # shared Go code
└── docker-compose (dev/prod)
```

## 🎯 Planned features

- 🌐 Distributed microservices
- 🔐 Auth: OIDC, OAuth2 (auth code), SSO, forgot/reset password
- 🏢 Multi-tenancy (+ single-tenant deployments)
- 🛡 Authorization: RBAC, ABAC
- 📋 Security, audit logging, structured server logging, notifications

## 🛠 Tech stack

`Go` · `Node (Yarn 4)` · `Postgres` · `Docker`

## 📦 Quick start

```bash
git clone https://github.com/m3rashid/zeus.git
cd zeus
yarn
yarn dev:up    # start
yarn dev:down  # stop
```

## 📄 License

ISC © [m3rashid](https://github.com/m3rashid)
