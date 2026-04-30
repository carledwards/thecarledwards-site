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
      colors: {
        ink: {
          950: '#0a0c10',
          900: '#0d1117',
          800: '#161b22',
          700: '#21262d',
          600: '#30363d',
          500: '#484f58',
          400: '#8b949e',
          300: '#b1bac4',
          200: '#c9d1d9',
          100: '#e6edf3',
        },
        accent: {
          DEFAULT: '#7ee787',
          dim: '#56d364',
        },
      },
      typography: ({ theme }) => ({
        invert: {
          css: {
            '--tw-prose-body': theme('colors.ink.200'),
            '--tw-prose-headings': theme('colors.ink.100'),
            '--tw-prose-lead': theme('colors.ink.300'),
            '--tw-prose-links': theme('colors.accent.DEFAULT'),
            '--tw-prose-bold': theme('colors.ink.100'),
            '--tw-prose-counters': theme('colors.ink.400'),
            '--tw-prose-bullets': theme('colors.ink.500'),
            '--tw-prose-hr': theme('colors.ink.700'),
            '--tw-prose-quotes': theme('colors.ink.200'),
            '--tw-prose-quote-borders': theme('colors.ink.600'),
            '--tw-prose-captions': theme('colors.ink.400'),
            '--tw-prose-code': theme('colors.ink.100'),
            '--tw-prose-pre-code': theme('colors.ink.200'),
            '--tw-prose-pre-bg': theme('colors.ink.950'),
            '--tw-prose-th-borders': theme('colors.ink.600'),
            '--tw-prose-td-borders': theme('colors.ink.700'),
          },
        },
      }),
    },
  },
  plugins: [typography],
};
