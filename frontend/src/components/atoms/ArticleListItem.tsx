'use client'

import { PostBookmark } from '@/api/postBookmark'
import { STATUS_CODE } from '@/const'
import { AttachmentIcon } from '@chakra-ui/icons'
import { ListItem, Text, Box, Flex, Spacer, Image, Badge, Link, Button } from '@chakra-ui/react'
import { useState } from 'react'

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

type ArticleListItemProps = {
  article: ArticleProps
  token: string | undefined
}

export const ArticleListItem = (props: ArticleListItemProps) => {
  const { article, token } = props
  const isLogin = token !== undefined
  const [isBookmark, setIsBookmark] = useState(false)

  const postBookmark = async (articleId: string) => {
    const { data, status } = await PostBookmark(articleId)
    switch (status) {
      case STATUS_CODE.OK:
        setIsBookmark(!isBookmark)
        break
      default:
        break
    }
  }

  // SVGのいいねアイコン
  const LikeIcon = () => (
    <svg
      viewBox="0 0 24 24"
      width="16"
      height="16"
      stroke="currentColor"
      strokeWidth="2"
      fill="none"
      strokeLinecap="round"
      strokeLinejoin="round"
      color="red.500"
    >
      <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
    </svg>
  )

  // 日付のフォーマット
  const formatDate = (dateString: string) => {
    const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric' }
    return new Date(dateString).toLocaleDateString(undefined, options)
  }

  return (
    <ListItem p={4} borderWidth="1px" borderRadius="lg" overflow="hidden" boxShadow="sm" bg="white.primary">
      <Flex alignItems="center">
        <Image
          borderRadius="full"
          boxSize="50px"
          src={article.publisher_image_url}
          alt={article.publisher_name}
          mr={4}
        />
        <Box flex="1">
          <Text fontWeight="bold" fontSize="lg">
            <Link href={article.url} isExternal color="teal.500">
              {article.title}
            </Link>
          </Text>
          <Text color="gray.500" fontSize="sm">
            Published by {article.publisher_name} on {formatDate(article.created_at)}
          </Text>
        </Box>
        <Spacer />
        <Box textAlign="right">
          <Badge colorScheme="purple" mr={2}>
            {article.quote_source}
          </Badge>
          <Button variant="ghost" colorScheme="red" size="sm" leftIcon={<LikeIcon />}>
            {article.likes_count}
          </Button>
          {isLogin && (
            <AttachmentIcon
              _hover={{ cursor: 'pointer' }}
              color={isBookmark ? 'yellow.primary' : 'gray.placeholder'}
              onClick={() => postBookmark(String(article.id))}
            >
              ブックマーク
            </AttachmentIcon>
          )}
        </Box>
      </Flex>
    </ListItem>
  )
}
