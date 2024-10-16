import { transform } from "typescript";

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
      arrowMoving: {
        "0%": { transform: "translateY(0px) scale(1.1)" },
        "50%": { transform: "translateY(5px) scale(1.1)" },
        "100%": { transform: "translateY(0px) scale(1.1)" },
      },
    },
    animation: {
      textForwards: "textMoving 0.3s linear normal forwards",
      textBackwards: "textReturning 0.3s linear forwards",
      arrowForwards: "arrowMoving 0.5s linear infinite",
    },
    fontFamily: {
      rubik: '"Rubik", sans-serif',
      caveat: '"Caveat", sans-serif',
    },
    extend: {
      width: {
        "90%": "90%",
        "95%": "95%",
        375: "375px",
        431: "431px",
        540: "540px",
        668: "668px",
        874: "874px",
      },
      height: {
        282: "282px",
        440: "440px",
        500: "500px",
        598: "598px",
        583: "583px",
      },
      padding: {
        5: "20px",
        4: "16px",
      },
      borderRadius: {
        20: "20px",
        40: "40px",
      },
      inset: {
        "1%": "1%",
        "7%": "7%",
        "4%": "4%",
        "10%": "10%",
        "17%": "17%",
        "25%": "25%",
        "45%": "45%",
        "55%": "55%",
        "70%": "70%",
      },
      maxWidth: {
        335: "335px",
        375: "375px",
        420: "420px",
        688: "688px",
        728: "728px",
        944: "944px",
      },
      fontSize: {
        40: "40px",
        64: "64px",
      },
      borderWidth: {
        6: "6px",
      },
    },
  },

  plugins: [],
};
