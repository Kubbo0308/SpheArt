import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { Header } from '@/components/organisms/Header'
import { Provider } from './provider'
import { CookiesProvider } from 'next-client-cookies/server'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Tech Pulse',
  description: '技術記事のまとめサイト'
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body className={inter.className}>
        <Provider>
          <CookiesProvider>
            <Header />
            {children}
          </CookiesProvider>
        </Provider>
      </body>
    </html>
  )
}
