module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      textColor: {
        skin: {
          base: "var(--color-text-base)",
          muted: "var(--color-text-muted)",
          inverted: "var(--color-text-inverted)",
        },
      },
      backgroundColor: {
        skin: {
          fill: "var(--color-fill)",
          muted: "var(--color-muted)",
          elevated: "var(--color-elevated)",
          "button-accent": "var(--color-button-accent)",
          "button-accent-hover": "var(--color-button-accent-hover)",
          "button-muted": "var(--color-button-muted)",
          "button-elevated": "var(--color-button-elevated)",
          "button-elevated-hover": "var(--color-button-elevated-hover)",
        },
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
