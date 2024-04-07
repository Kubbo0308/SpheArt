'use client'

import { Text, Box, Flex, Image, Link } from '@chakra-ui/react'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { BookmarkButton } from '../../atoms/BookmarkButton'
import { useArticleCard } from './ArticleCard.hooks'

export type ArticleProps = {
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

type ArticleCardProps = {
  article: ArticleProps
  token: RequestCookie | undefined
}

export const ArticleCard = (props: ArticleCardProps) => {
  const { article, token } = props
  const { isBookmark, postBookmark, formatDate } = useArticleCard(token)

  return (
    <Box borderRadius="8px" overflow="hidden" boxShadow="sm" bg="white.primary" w="320px">
      <Image
        // src={article.publisher_image_url}
        src="/no_image.svg"
        alt={article.publisher_name}
        h="180px"
      />
      <Box p="10px">
        <Text fontSize="16px" fontWeight={700} lineHeight={1.8}>
          <Link href={article.url} isExternal color="teal.500">
            {article.title}
          </Link>
        </Text>
        <Flex mt="10px" justifyContent="space-between" alignItems="center">
          <Text fontSize="16px" fontWeight={500}>
            {formatDate(article.created_at)}
          </Text>
          <Flex gap="10px" alignItems="center">
            <Flex gap="3px" alignItems="center">
              <Image src="/heart_256.svg" alt="" w="16px" h="16px" />
              <Text fontSize="16px" fontWeight={500}>
                {article.likes_count}
              </Text>
            </Flex>
            <BookmarkButton onClick={() => postBookmark(String(article.id))} isBookmark={isBookmark} />
          </Flex>
        </Flex>
      </Box>
    </Box>
  )
}
