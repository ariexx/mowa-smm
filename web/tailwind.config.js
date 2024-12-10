/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: ['selector', '[class*="app-dark"]'],
    content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    plugins: [require('tailwindcss-primeui')],
    theme: {
        screens: {
            sm: '576px',
            md: '768px',
            lg: '992px',
            xl: '1200px',
            '2xl': '1920px'
        },
        borderRadius: {
            none: '0',
            sm: '0.25rem',
            DEFAULT: '0.5rem',
            md: '0.375rem',
            lg: '0.75rem',
            xl: '1rem',
            '2xl': '1.5rem',
            '3xl': '2rem',
            full: '9999px'
        }
    }
};
