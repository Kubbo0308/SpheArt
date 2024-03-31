'use client'

import { ArticleList } from '@/components/organisms/ArticleList'
import { Button, Container, Flex } from '@chakra-ui/react'
import { useBookmarkPageHooks } from './Bookmark.hooks'

export const BookmarkPage = () => {
  const { currentPage, articles, goNextPage, backPreviousPage } = useBookmarkPageHooks()
  return (
    <Container>
      <ArticleList articles={articles} token={undefined} />
      <Flex gap="20px">
        <Button onClick={backPreviousPage}>前へ</Button>
        <Button onClick={goNextPage}>次へ</Button>
      </Flex>
    </Container>
  )
}
