## ZEUS

### What does it take to make a good SaaS framework?

<br />

### What features do I want in a SaaS framework?

1. [ ] Distributed and planned microservices `*`
2. [ ] Authentication and User Management `*`
   - [ ] OpenID Connect (OIDC) `*`
   - [ ] OAuth2 (Authorization Code Flow) `*`
   - [ ] SSO (Single Sign On) `+`
   - [ ] Forgot/Reset Password `*`
3. [ ] Multi-tenancy `*`
   - [ ] Also able to support single Tenant deployments `+`
4. [ ] Authorization
   - [ ] Role Based Access Control (RBAC) `*`
   - [ ] Attribute Based Access Control (ABAC) `+`
5. [ ] Security `*`
6. [ ] Audit Logging (Who did what and when) `*`
7. [ ] Structured Server Logging `*`
8. [ ] Notifications `*`
   - [ ] Email `*`
   - [ ] Whatsapp `+`
9. [ ] Platform wide search `*`
10. [ ] Near Real Time (NRT) updates `*`
11. [ ] Encrypted Key-Value store `*`
12. [ ] Forms Builder `+`
13. [ ] Email Templates Builder `+`
14. [ ] Web hooks for other apps `+`
15. [ ] Workspaces (Data independent spaces within a tenant) `*`
16. [ ] Monitoring and Alerting `*`
17. [ ] Print and Export from the server `*`
    - [ ] PDF `*`
    - [ ] CSV `*`
    - [ ] Excel `+`
18. [ ] Clients `*`
    - [ ] Web `*`
    - [ ] REST APIs `*`
    - [ ] Mobile (iOS, Android) `*`
    - [ ] Desktop (Windows, Mac, Linux) `+`
    - [ ] CLI (Command Line Interface) `+`
    - [ ] GraphQL API `+`
19. [ ] A solid Core and extensible apps model `*`
    - [ ] Develop and deploy apps independently `*`
    - [ ] Develop apps independent of the core `+`
    - [ ] Marketplace (App store) for apps `+`
    - [ ] App versioning `+`
    - [ ] App submit and review process `#`
20. [ ] Getting Started (customer facing) app for tenants `*`
    - [ ] Easy onboarding for tenants `*`
    - [ ] External Payments and Billing `*`
    - [ ] Cost and price calculator `+`
    - [ ] Limits and Quotas `+`

<br />

```text
* Absolutely required
+ Good to have
# Icing on the cake
```
