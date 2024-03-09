import { searchArticlesInTitle } from '@/api/article'
import { ArticleList } from '@/components/organisms/ArticleList'
import { EnsureString } from '@/utils/EnsureString'
import { Container } from '@chakra-ui/react'

export default async function Search({
  searchParams
}: {
  searchParams: { [key: string]: string | string[] | undefined }
}) {
  const searchTitle = searchParams.title

  const articles = await searchArticlesInTitle(EnsureString(searchTitle))
  return (
    <Container maxW="container.md" py="5%">
      <ArticleList articles={articles} />
    </Container>
  )
}
