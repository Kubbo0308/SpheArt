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
      placeholder: "#a9a9a9"
    },
    white: {
      primary: "#f4f5f7"
    }
  },
  styles: {
    global: {
      body: {
        bg: "yellow.secondary",
        placeholder: "gray.placeholder"
      },
      font: {
        color: "black.primary",
        fontSize: "16px"
      }
    }
  }
})