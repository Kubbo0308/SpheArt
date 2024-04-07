import { Flex } from '@chakra-ui/react'
import { ArticleCard } from '../molecules/ArticleCard/ArticleCard'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'

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

type ArticleListProps = {
  articles: ArticleProps[]
  token: RequestCookie | undefined
}

export const ArticleList = (props: ArticleListProps) => {
  const { articles, token } = props
  return (
    <Flex flexWrap="wrap" gap="20px" justifyContent="center">
      {articles.map((article: ArticleProps) => (
        <div key={article.id}>
          <ArticleCard article={article} token={token} />
        </div>
      ))}
    </Flex>
  )
}
