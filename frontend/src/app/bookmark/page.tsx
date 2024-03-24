import { GetAllBookmark } from '@/api/getBookmark'
import { ArticleList } from '@/components/organisms/ArticleList'
import { Box, Container } from '@chakra-ui/react'

type ArticleProps = {
  id: number
  title: string
  url: string
  created_at: string
  updated_at: string
  publisher_id: string
  publisher_name: string
  publisher_image_url: string
  likes_count: number
  quote_source: string
}

export default async function Bookmark() {
  const { data, status } = await GetAllBookmark()
  return (
    <Container>
      <ArticleList articles={data} token={undefined} />
    </Container>
  )
}
