/** @type {import('tailwindcss').Config} */
export default {
    content: ['./www/**/*.{svelte,css}'],
    darkMode: 'class',
    theme: {
        extend: {
            animation: {
                shake: 'shake 0.4s ease',
            },
            boxShadow: {
                sm: '2px 2px 0px hsl(var(--colour-secondary))',
                md: '4px 4px 0px hsl(var(--colour-secondary))',
                lg: '6px 6px 0px hsl(var(--colour-secondary))',
            },
            colors: {
                primary: 'hsl(var(--colour-primary) / <alpha-value>)',
                secondary: 'hsl(var(--colour-secondary) / <alpha-value>)',
                contrast: 'hsl(var(--colour-contrast) / <alpha-value>)',
                'contrast-light': 'hsl(var(--colour-contrast-light) / <alpha-value>)',
                'contrast-dark': 'hsl(var(--colour-contrast-dark) / <alpha-value>)',
            },
            fontFamily: {
                sans: ['Roboto'],
                serif: ['Roboto Slab'],
                mono: ['Roboto Mono'],
            },
        },
    },
    plugins: [require('@tailwindcss/typography')],
    safelist: ['opacity-0', 'pointer-events-none', 'text-5xl'],
}
