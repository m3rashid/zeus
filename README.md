# ZEUS

SaaS framework skeleton (monorepo). Goal: a reusable, multi-tenant SaaS foundation with OIDC/OAuth2 auth, RBAC/ABAC, audit logging, and notifications.

## Workspace

`types/` `web/` `mobile/` `admin/` (Yarn workspaces) + Go microservices: `auth/`, `notifications/`, `common/`.

## Planned features

- Distributed microservices
- Auth: OIDC, OAuth2 (auth code), SSO, forgot/reset password
- Multi-tenancy (+ single-tenant deployments)
- Authorization: RBAC, ABAC
- Security, audit logging, structured server logging, notifications

## Tech

Go · Node (Yarn 4) · Postgres · Docker

## Quick start

```bash
git clone https://github.com/m3rashid/zeus.git
cd zeus
yarn
yarn dev:up    # start
yarn dev:down  # stop
```

## License

ISC
