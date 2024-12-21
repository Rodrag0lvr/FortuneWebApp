/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./public/**/*.html",   // Archivos HTML
    "./public/**/*.css",    // Archivos CSS
  ],
  theme: {
    extend: {},
    
  },
  plugins: [
    function ({ addComponents }) {
      addComponents({
        '.btn-primary': {
          backgroundColor: '#b91c1c', // bg-red-700
          color: '#ffffff', // text-white
          fontWeight: 'bold', // font-bold
          padding: '0.5rem 1rem', // py-2 px-4
          borderRadius: '0.5rem', // rounded-lg
         
        },
        '.btn-secondary': {
          backgroundColor: '#fca5a5', // bg-red-300
          color: '#b91c1c', // text-red-700
          fontWeight: 'bold', // font-bold
          padding: '0.5rem 1rem', // py-2 px-4
          borderRadius: '0.5rem', // rounded-lg
         
        },
      });
    },
  ],};
