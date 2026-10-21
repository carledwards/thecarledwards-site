import { defineCollection, z } from 'astro:content';

const blog = defineCollection({
  type: 'content',
  schema: z.object({
    title: z.string(),
    date: z.coerce.date(),
    description: z.string().optional(),
    tags: z.array(z.string()).default([]),
    draft: z.boolean().default(false),
    kind: z.enum(['Writing', 'Project', 'Project update', 'Learning']).default('Writing'),
  }),
});

const lessons = defineCollection({
  type: 'content',
  schema: z.object({
    title: z.string(),
    series: z.string(),
    seriesTitle: z.string(),
    order: z.number(),
    summary: z.string(),
    tags: z.array(z.string()).default([]),
    draft: z.boolean().default(false),
  }),
});

// Short announcements can link directly to a project or lesson without a blog post.
const updates = defineCollection({
  type: 'data',
  schema: z.object({
    title: z.string(),
    date: z.coerce.date(),
    description: z.string(),
    kind: z.enum(['Project', 'Project update', 'Learning']),
    href: z.string().regex(/^(\/(?!\/)|https:\/\/)/),
    draft: z.boolean().default(false),
  }),
});

export const collections = { blog, lessons, updates };
