import { extendTheme } from "@chakra-ui/react";

export const theme = extendTheme({
  colors: {
    yellow: {
      primary: "#eccc6e",
      secondary: "#fff8dc"
    },
    black: {
      primary: "#333333"
    },
    gray: {
      primary: "#cccccc"
    },
    white: {
      primary: "#f4f5f7"
    }
  },
  styles: {
    global: {
      body: {
        bg: "yellow.secondary"
      },
      font: {
        color: "black.primary",
        fontSize: "16px"
      }
    }
  }
})