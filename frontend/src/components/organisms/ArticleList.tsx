import { Flex } from '@chakra-ui/react'
import { ArticleCard, ArticleProps } from '../molecules/ArticleCard/ArticleCard'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'

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
