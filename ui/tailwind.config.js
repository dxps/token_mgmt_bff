module.exports = {
  mode: 'jit',
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  plugins: [
    require("@tailwindcss/typography"),
    require('daisyui')
  ],
  daisyui: {
    themes: [
      {
        mytheme: {
          "primary": "#67e8f9",
          "secondary": "#a8ffcf",
          "accent": "#f0abfc",
          "neutral": "#2F2B3B",
          "base-100": "#4b5563",
          "info": "#bfdbfe",
          "success": "#35E977",
          "warning": "#F4B357",
          "error": "#E9353E",
        },
      },
    ],
  }
};
