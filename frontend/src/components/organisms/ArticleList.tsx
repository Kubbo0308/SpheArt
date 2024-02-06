import { List } from '@chakra-ui/react'
import { ArticleListItem } from '../atoms/ArticleListItem'

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
}

export const ArticleList = (props: ArticleListProps) => {
  const { articles } = props
  return (
    <List spacing={3} mt="2%">
      {articles.map((article: ArticleProps) => (
        <div key={article.id}>
          <ArticleListItem article={article} />
        </div>
      ))}
    </List>
  )
}
