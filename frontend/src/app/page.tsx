import { TopPage } from '@/components/pages/Top/TopPage'
import { cookies } from 'next/headers'

export default function Home() {
  const cookieStore = cookies()
  const token = cookieStore.get('token')
  return <TopPage token={token} />
}
