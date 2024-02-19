import { searchArticlesInTitle } from '@/api/article'
import { ArticleList } from '@/components/organisms/ArticleList'
import { ensureString } from '@/utils/ensureString'
import { Container } from '@chakra-ui/react'

export default async function Search({
  searchParams
}: {
  searchParams: { [key: string]: string | string[] | undefined }
}) {
  const searchTitle = searchParams.title

  const articles = await searchArticlesInTitle(ensureString(searchTitle))
  return (
    <Container maxW="container.sm">
      <ArticleList articles={articles} />
    </Container>
  )
}
