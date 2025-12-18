/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/**/*.{html,js,ts,vue}',  // 确保你的文件路径被正确扫描
  ],
  theme: {
    extend: {
      // animation: {
      //   'scroll-float': 'scroll-float 1.6s ease-in-out infinite',
      // },
      // keyframes: {
      //   'scroll-float': {
      //     '0%, 100%': {
      //       transform: 'translateY(0)',
      //       opacity: '0.85',
      //     },
      //     '50%': {
      //       transform: 'translateY(8px)',
      //       opacity: '1',
      //     },
      //   },
      // },
      fontFamily: {
        sans: ['"Poppins"', 'system-ui', '-apple-system', 'Segoe UI', 'Roboto', 'Helvetica Neue', 'Arial', 'sans-serif'],
      },
      colors: {
        kanna: {
          DEFAULT: "#4f7da6",
          50: "#f2f6f9",
          100: "#dbe5ed",
          200: "#b7cbdc",
          300: "#8daac5",
          400: "#6487ac",
          500: "#4f7da6", // 主色
          600: "#406586",
          700: "#344f69",
          800: "#283b4e",
          900: "#1d2937",
        }
      },
    },
  },
  plugins: [],
}