# thecarledwards.com

Personal site of Carl Edwards. Built with [Astro](https://astro.build),
styled with Tailwind, deployed on Cloudflare Pages.

## Local development

```bash
npm install
npm run dev      # http://localhost:4321
npm run build    # → dist/
npm run preview  # serves dist/ locally
```

## Project layout

```
.
├── astro.config.mjs        # publicDir: static/, outDir: dist/
├── src/
│   ├── components/         # Header, Footer
│   ├── layouts/            # BaseLayout
│   ├── pages/              # /, /blog, /blog/[slug], /about, /rss.xml
│   ├── content/
│   │   ├── config.ts       # collection schemas
│   │   └── blog/           # markdown posts
│   └── styles/global.css
├── static/                 # static assets, _redirects, robots.txt, favicon
└── dist/                   # build output (gitignored)
```

## Adding a post

Drop a markdown file in `src/content/blog/`:

```markdown
---
title: "Post Title"
date: 2026-04-29
description: "One-line summary used for SEO and the post list."
tags: ["Tag1", "Tag2"]
---

Body...
```

## Deployment

Cloudflare Pages, building from `main`:

- Build command: `npm run build`
- Build output: `dist`
- Env: `NODE_VERSION=20`

Old Hugo URLs (`/posts/<slug>/`, `/index.xml`) redirect to the new paths via
`static/_redirects`.

## License & copyright

- Source code: [MIT](./LICENSE)
- Written content (posts, training, media): see [COPYRIGHT.md](./COPYRIGHT.md)

## Homepage updates

The homepage shows the seven newest published blog posts and standalone updates,
sorted by date. Blog posts default to the `Writing` label; set `kind` in their
frontmatter to `Project`, `Project update`, or `Learning` when appropriate.

For an announcement that links directly to a project or lesson, copy
`src/content/updates/example.json` to a descriptive filename. Set its title, date
(`YYYY-MM-DD`), description, kind, and href, then set `draft` to `false`. Keep the
example as a draft. Published updates also appear in RSS. Use one entry per
announcement to avoid duplicating a blog post in the stream.
