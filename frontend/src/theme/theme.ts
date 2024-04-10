import { extendTheme } from "@chakra-ui/react";

export const theme = extendTheme({
  colors: {
    yellow: {
      primary: "#edc313",
      secondary: "#fff8dc"
    },
    black: {
      primary: "#672a3f"
    },
    gray: {
      primary: "#cccccc",
      placeholder: "#a9a9a9",
      border: "#e2e8f0",
      light: "#f7fafc",
      accent: "#1a202c"
    },
    white: {
      primary: "#ffffff"
    },
    blue: {
      bg: "#e2e8f0",
      accent: "#3182ce",
      pale: "#90cdf4"
    }
  },
  styles: {
    global: {
      body: {
        bg: "blue.bg",
        placeholder: "gray.placeholder"
      },
      font: {
        color: "black.primary",
        fontSize: "16px"
      }
    }
  }
})