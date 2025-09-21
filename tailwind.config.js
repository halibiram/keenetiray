/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: '#3182ce',
        secondary: '#2d3748',
        success: '#38a169',
        warning: '#d69e2e',
        danger: '#e53e3e',
        'k-dark-blue': '#1a202c',
        'k-dark-gray': '#2d3748',
        'k-gray': '#4a5568',
        'k-white': '#ffffff',
        'k-light-gray': '#f7fafc',
        card: 'rgba(255, 255, 255, 0.95)',
      },
      textColor: {
        primary: '#2d3748',
        secondary: '#4a5568',
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
        mono: ['Fira Code', 'monospace'],
      },
      backgroundImage: {
        'gradient-body': 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      },
    },
  },
  plugins: [],
}
