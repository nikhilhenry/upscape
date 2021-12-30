function withOpacity(variableName) {
  return ({ opacityValue }) => {
    if (opacityValue !== undefined) {
      return `rgba(${variableName},${opacityValue})`;
    }
    return `rgb(${variableName})`;
  };
}

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
          fill: withOpacity("var(--color-fill)"),
          muted: withOpacity("var(--color-muted)"),
          elevated: withOpacity("var(--color-elevated)"),
          "button-accent": withOpacity("var(--color-button-accent)"),
          "button-accent-hover": withOpacity(
            "var(--color-button-accent-hover)"
          ),
          "button-muted": withOpacity("var(--color-button-muted)"),
          "button-elevated": withOpacity("var(--color-button-elevated)"),
          "button-elevated-hover": withOpacity(
            "var(--color-button-elevated-hover)"
          ),
        },
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
