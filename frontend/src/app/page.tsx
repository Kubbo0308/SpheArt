import { getArticles } from '@/api/article'
import { ArticleList } from '@/components/organisms/ArticleList'
import { Container } from '@chakra-ui/react'

export default async function Home() {
  const staticData = await getArticles()
  return (
    <Container maxW="container.md" py="5%">
      <ArticleList articles={staticData} />
    </Container>
  )
}
