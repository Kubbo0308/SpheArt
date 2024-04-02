import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { Provider } from './provider'
import { Header } from '@/components/organisms/Header/Header'
import { cookies } from 'next/headers'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Tech Pulse',
  description: '技術記事のまとめサイト'
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  const cookieStore = cookies()
  const token = cookieStore.get('token')
  return (
    <html lang="ja">
      <body className={inter.className}>
        <Provider>
          <Header token={token} />
          {children}
        </Provider>
      </body>
    </html>
  )
}
