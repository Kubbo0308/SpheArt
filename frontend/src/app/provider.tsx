'use client'

import { theme } from '../../theme/theme'
import { ChakraProvider } from '@chakra-ui/provider'

export function Provider({ children }: { children: React.ReactNode }) {
  return <ChakraProvider theme={theme}>{children}</ChakraProvider>
}
