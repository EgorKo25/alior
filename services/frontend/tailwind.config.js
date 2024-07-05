/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,ts,tsx,jsx}"],
  theme: {
    colors: {
      transparent: "rgba(0,0,0,0)",
      white: "rgb(255, 255, 255)",
      black: "rgb(0, 0, 0)",
      orange: {
        100: "rgb(255, 182, 97)",
        900: "rgb(107, 57, 0)",
      },
      blue: {
        100: "rgb(134, 167, 252)",
        900: "rgb(0, 17, 102)",
      },
      accent: "rgb(255, 83, 122)",
    },
    extend: {},
    fontFamily: {
      rubik: '"Rubik", sans-serif',
      caveat: '"Caveat", sans-serif',
    },
  },

  plugins: [],
};
