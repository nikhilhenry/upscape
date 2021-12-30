function withOpacity(variableName) {
  return ({ opacityValue }) => {
    if (opacityValue !== undefined) {
      return `rgba(var(${variableName}), ${opacityValue})`;
    }
    return `rgb(var(${variableName}))`;
  };
}

module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      textColor: {
        skin: {
          base: withOpacity("--color-text-base"),
          muted: withOpacity("--color-text-muted"),
          inverted: withOpacity("--color-text-inverted"),
        },
      },
      backgroundColor: {
        skin: {
          fill: withOpacity("--color-fill"),
          base: withOpacity("--color-base"),
          muted: withOpacity("--color-muted"),
          elevated: withOpacity("--color-elevated"),
          "button-accent": withOpacity("--color-button-accent"),
          "button-accent-hover": withOpacity("--color-button-accent-hover"),
          "button-muted": withOpacity("--color-button-muted"),
          "button-elevated": withOpacity("--color-button-elevated"),
          "button-elevated-hover": withOpacity("--color-button-elevated-hover"),
        },
      },
      borderColor: {
        skin: {
          fill: withOpacity("--color-fill"),
          muted: withOpacity("--color-muted"),
          elevated: withOpacity("--color-elevated"),
        },
      },
      shadowColor: {
        skin: {
          muted: withOpacity("--color-muted"),
        },
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
