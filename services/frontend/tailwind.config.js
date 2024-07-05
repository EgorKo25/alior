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
    keyframes: {
      textMoving: {
        "0%": { transform: "translateX(0px)" },
        "50%": { transform: "translateX(30px)" },
        "100%": { transform: "translateX(27px)" },
      },
      textReturning: {
        "0%": { transform: "translateX(27px)" },
        "50%": { transform: "translateX(-3px)" },
        "100%": { transform: "translateX(0px)" },
      },
    },
    animation: {
      textForwards: "textMoving 0.3s linear normal forwards",
      textBackwards: "textReturning 0.3s linear forwards",
    },
    fontFamily: {
      rubik: '"Rubik", sans-serif',
      caveat: '"Caveat", sans-serif',
    },
    extend: {},
  },

  plugins: [],
};
