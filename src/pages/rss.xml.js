import rss from '@astrojs/rss';
import { getCollection } from 'astro:content';

export async function GET(context) {
  const posts = (await getCollection('blog', ({ data }) => !data.draft))
    .sort((a, b) => b.data.date.valueOf() - a.data.date.valueOf());

  const updates = await getCollection('updates', ({ data }) => !data.draft);

  return rss({
    title: 'Carl Edwards',
    description: 'Engineering, systems, and AI tooling — by Carl Edwards.',
    site: context.site,
    items: [...posts.map((post) => ({
      title: post.data.title,
      pubDate: post.data.date,
      description: post.data.description ?? '',
      link: `/blog/${post.slug}/`,
      categories: post.data.tags,
    })), ...updates.map(({ data }) => ({
      title: data.title,
      pubDate: data.date,
      description: data.description,
      link: data.href,
      categories: [data.kind],
    }))].sort((a, b) => b.pubDate.valueOf() - a.pubDate.valueOf()),
  });
}
