import { extendTheme } from "@chakra-ui/react";

export const theme = extendTheme({
  colors: {
    yellow: {
      primary: "#eccc6e",
      secondary: "#fff8dc"
    },
    black: {
      primary: "#672a3f"
    },
    gray: {
      primary: "#cccccc",
      placeholder: "#a9a9a9",
      border: "#e2e8f0"
    },
    white: {
      primary: "#ffffff"
    },
    blue: {
      bg: "#f7fafc",
      accent: "#3182ce"
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