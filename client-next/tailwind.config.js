const colors = require("tailwindcss/colors");

module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: {
          100: "#6dd4ff",
          200: "#48c0ff",
          300: "#2ca9ff",
          400: "#1890ff",
          500: "#1074e5",
          600: "#0945c7",
          700: "#040ca8",
          800: "#170087",
        },
        background: {
          100: "#accee1",
          200: "#8cbdd3",
          300: "#73acc4",
          400: "#619ab2",
          500: "#54869d",
          600: "#517185",
          700: "#485a69",
          800: "#39424b",
          900: "#25282c",
        },
        txtcolor: colors.slate,
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
