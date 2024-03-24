import { getArticles } from '@/api/article'
import { ArticleList } from '@/components/organisms/ArticleList'
import { Container } from '@chakra-ui/react'
import { getCookies } from 'next-client-cookies/server'

export default async function Home() {
  const staticData = await getArticles()
  const cookies = getCookies()
  const token = cookies.get('token')
  return (
    <Container maxW="container.md" py="5%">
      <ArticleList articles={staticData} token={token} />
    </Container>
  )
}
