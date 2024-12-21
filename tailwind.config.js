/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pkg/**/*.html", // Todos los archivos HTML en "pkg"
    "./pkg/web/views/liquidations/**/*.html", // Todos los archivos HTML en "pkg"
    "./pkg/**/*.js", // Todos los archivos JS en "pkg"
    "./pkg/**/*.css", // Archivos CSS específicos en "pkg"
    "./public/**/*.html", // Archivos HTML en "public"
    "./public/**/*.js", // Archivos JS en "public"

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
        '.btn-secondary-green': {
          backgroundColor: '#86efac',
          color: '#166534', 
          fontWeight: 'bold', 
          padding: '0.5rem 1rem',
          borderRadius: '0.5rem', 
        
      },  
      });
    },
  ],
};
