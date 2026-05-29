# Frontend Standards

## Structure

```
frontend/src/
├── pages/        # One component per route
├── components/   # Reusable UI — always Tailwind, always typed props
├── hooks/        # Custom React hooks
├── services/     # Axios API calls — one file per domain (authService, userService, etc.)
├── store/        # Zustand stores
└── types/        # TypeScript interfaces — single source of truth for all shapes
```

---

## React / TypeScript Standards

- Strict TypeScript — no `any`
- Functional components only, hooks for all state
- `@tanstack/react-query` for all server state (fetching, caching, mutations)
- `zustand` for client-only state (auth tokens, UI state)
- `react-hook-form` + `zod` for all form validation
- `axios` with a configured instance in `services/api.ts`
- Every async operation has a loading state (spinner, skeleton, or disabled button)
- Every async operation has an error state

**Styling — Tailwind CSS only:**
- No inline styles, no CSS modules, no styled-components
- Install `tailwindcss`, `postcss`, `autoprefixer` as devDependencies on every project
- Mobile-first: write mobile layout first, add `sm:`, `md:`, `lg:` overrides
- Verify every page at 375px (mobile), 768px (tablet), 1280px (desktop) before marking done

**Toasts — mandatory on every mutation:**
- Install `react-hot-toast`
- `toast.success()` on success, `toast.error()` on failure — no exceptions
- Never rely on inline error text alone for user feedback

**SEO — every page:**
- Install `react-helmet-async`
- Wrap app in `<HelmetProvider>`
- Every page component sets its own `<Helmet>` with `<title>` and `<meta name="description">`

**Inline comment rule:** One line, WHY only, non-obvious only.

```tsx
// react-query retries on 401 before the interceptor refreshes — disable retry for auth queries
const { data } = useQuery({ queryKey: ['me'], queryFn: authService.me, retry: false })
```

---

## Default Pages (Every App)

Created during initial scaffold — before any domain features. Never skip these.

### `frontend/src/pages/PrivacyPolicy.tsx`
- Route: `/privacy-policy`
- Pre-filled placeholder: `[TODO: APP NAME]`, `[TODO: CONTACT EMAIL]`, `[TODO: DATE]`
- Covers: data collected, usage, cookies, third-party services, user rights, contact info

### `frontend/src/pages/TermsAndConditions.tsx`
- Route: `/terms`
- Pre-filled placeholder: `[TODO: APP NAME]`, `[TODO: COMPANY]`, `[TODO: DATE]`
- Covers: acceptance, obligations, prohibited activities, disclaimer, liability, governing law

### `frontend/src/components/Footer.tsx`
- Persistent on every screen — authenticated and public
- Links to `/privacy-policy` and `/terms`
- Shows app name and current year

Placeholders marked `[TODO: ...]` — user fills these before going to production.

---

## Docker — Frontend

- Development: `Dockerfile.dev` — `node:22-alpine` with Vite hot-reload
  - `CHOKIDAR_USEPOLLING=true` — required for file watching inside Docker on Windows and Mac
- Production: multi-stage `Dockerfile` — `node:22-alpine` builder → `nginx:1.27-alpine` runner
  - Build arg: `VITE_API_URL` (default `/api`)
  - Custom `nginx.conf`: SPA routing, `/api/` proxy, security headers, gzip, static asset caching
  - Run as non-root user
  - Health check via `wget /health.txt`
