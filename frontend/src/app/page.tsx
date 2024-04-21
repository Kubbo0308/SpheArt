import { runQiitaBatch, runZennBatch } from '@/api/batch'
import { TopPage } from '@/components/pages/Top/TopPage'
import { cookies } from 'next/headers'

export default async function Home() {
  if (process.env.NEXT_PUBLIC_ENV === 'prod') {
    await runQiitaBatch()
    await runZennBatch()
  }
  const cookieStore = cookies()
  const token = cookieStore.get('token')
  return <TopPage token={token} />
}
