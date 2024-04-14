import { SearchPage } from '@/components/pages/Search/SearchPage'
import { cookies } from 'next/headers'

export default async function Search() {
  const cookieStore = cookies()
  const token = cookieStore.get('token')
  return <SearchPage token={token} />
}
