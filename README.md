<div align="center">

# ⚡ ZEUS

**SaaS framework skeleton — monorepo.** Multi-tenant foundation with OIDC/OAuth2 auth, RBAC/ABAC, audit logging, notifications.

[![Node](https://img.shields.io/badge/Node-20-339933?logo=node.js&logoColor=white)](https://nodejs.org)
[![License: ISC](https://img.shields.io/badge/License-ISC-blue.svg)](LICENSE)

</div>

---

## 🗂 Workspace

`types/` `web/` `mobile/` `admin/` (Yarn workspaces) + Go microservices: `auth/`, `notifications/`, `common/`.

## 🎯 Planned features

- 🌐 Distributed microservices
- 🔐 Auth: OIDC, OAuth2 (auth code), SSO, forgot/reset password
- 🏢 Multi-tenancy (+ single-tenant deployments)
- 🛡 Authorization: RBAC, ABAC
- 📋 Security, audit logging, structured server logging, notifications

## 🛠 Tech

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

ISC
