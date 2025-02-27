import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        mono: ['var(--montserrat-font)'],
        sans: ['var(--inter-font)'],
      },
      colors: {
        primary: {
          DEFAULT: 'var(--primary)',
        },
        secondary: {
          DEFAULT: 'var(--secondary)',
        },
      },
    },
  },
  plugins: [],
}

export default config