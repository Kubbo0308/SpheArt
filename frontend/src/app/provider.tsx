'use client'

import { theme } from '../../theme/theme'
import { RecoilRoot } from 'recoil'
import { ChakraProvider } from '@chakra-ui/provider'

export function Provider({ children }: { children: React.ReactNode }) {
  return (
    <RecoilRoot>
      <ChakraProvider theme={theme}>{children}</ChakraProvider>
    </RecoilRoot>
  )
}
