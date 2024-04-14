import { SearchPage } from '@/components/pages/Search/SearchPage'
import { cookies } from 'next/headers'

export default async function Search({
  searchParams
}: {
  searchParams: { [key: string]: string | string[] | undefined }
}) {
  const searchTitle = searchParams.title
  const cookieStore = cookies()
  const token = cookieStore.get('token')
  return <SearchPage token={token} title={searchTitle} />
}
