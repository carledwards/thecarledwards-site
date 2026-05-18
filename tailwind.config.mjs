import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        sans: ['"Inter"', 'system-ui', 'sans-serif'],
        mono: ['"JetBrains Mono"', 'ui-monospace', 'monospace'],
      },
      // Palette is driven by CSS custom properties (channel triplets) so a
      // single `.light` class on <html> reskins the whole site. The
      // `<alpha-value>` form keeps every existing `/NN` opacity utility
      // working unchanged. Channel values live in src/styles/global.css.
      colors: {
        ink: {
          950: 'rgb(var(--ink-950) / <alpha-value>)',
          900: 'rgb(var(--ink-900) / <alpha-value>)',
          800: 'rgb(var(--ink-800) / <alpha-value>)',
          700: 'rgb(var(--ink-700) / <alpha-value>)',
          600: 'rgb(var(--ink-600) / <alpha-value>)',
          500: 'rgb(var(--ink-500) / <alpha-value>)',
          400: 'rgb(var(--ink-400) / <alpha-value>)',
          300: 'rgb(var(--ink-300) / <alpha-value>)',
          200: 'rgb(var(--ink-200) / <alpha-value>)',
          100: 'rgb(var(--ink-100) / <alpha-value>)',
        },
        accent: {
          DEFAULT: 'rgb(var(--accent) / <alpha-value>)',
          dim: 'rgb(var(--accent-dim) / <alpha-value>)',
        },
      },
      // Prose vars point straight at the channels (no <alpha-value>
      // placeholder — these land verbatim in CSS custom properties), so
      // lesson content flips with the theme automatically.
      typography: () => ({
        invert: {
          css: {
            '--tw-prose-body': 'rgb(var(--ink-200))',
            '--tw-prose-headings': 'rgb(var(--ink-100))',
            '--tw-prose-lead': 'rgb(var(--ink-300))',
            '--tw-prose-links': 'rgb(var(--accent))',
            '--tw-prose-bold': 'rgb(var(--ink-100))',
            '--tw-prose-counters': 'rgb(var(--ink-400))',
            '--tw-prose-bullets': 'rgb(var(--ink-500))',
            '--tw-prose-hr': 'rgb(var(--ink-700))',
            '--tw-prose-quotes': 'rgb(var(--ink-200))',
            '--tw-prose-quote-borders': 'rgb(var(--ink-600))',
            '--tw-prose-captions': 'rgb(var(--ink-400))',
            '--tw-prose-code': 'rgb(var(--ink-100))',
            '--tw-prose-pre-code': 'rgb(var(--ink-200))',
            '--tw-prose-pre-bg': 'rgb(var(--ink-900))',
            '--tw-prose-th-borders': 'rgb(var(--ink-600))',
            '--tw-prose-td-borders': 'rgb(var(--ink-700))',
          },
        },
      }),
    },
  },
  plugins: [
    typography,
    // `light:` — applies when <html> carries the `.light` class. Used to
    // re-tune fixed-palette accents (emerald/amber status badges) that the
    // var-driven `ink/accent` tokens don't cover.
    function ({ addVariant }) {
      addVariant('light', '.light &');
    },
  ],
};
