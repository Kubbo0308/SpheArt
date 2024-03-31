import { getCookies } from 'next-client-cookies/server'
import { TopPage } from '@/components/pages/Top/TopPage'

export default function Home() {
  const cookies = getCookies()
  const token = cookies.get('token')
  return <TopPage token={token} />
}
