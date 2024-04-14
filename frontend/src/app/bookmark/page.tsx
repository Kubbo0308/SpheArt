import { BookmarkPage } from '@/components/pages/Bookmark/BookmarkPage'
import { cookies } from 'next/headers'

export default function Bookmark() {
  const cookieStore = cookies()
  const token = cookieStore.get('token')
  return <BookmarkPage token={token} />
}
