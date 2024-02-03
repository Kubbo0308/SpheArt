import { getArticles } from '@/api/article'

export default async function Home() {
  const staticData = await getArticles()
  return (
    <p>{JSON.stringify(staticData)}</p>
  )
}
