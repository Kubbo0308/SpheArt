'use client'

import { ArticleList } from '@/components/organisms/ArticleList'
import { Box, Container } from '@chakra-ui/react'
import { useBookmarkPageHooks } from './Bookmark.hooks'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { Loading } from '@/components/molecules/Loading'

interface BookmarkPageProps {
  token: RequestCookie | undefined
}

export const BookmarkPage = (props: BookmarkPageProps) => {
  const { token } = props
  const { articles, loader, isVisible } = useBookmarkPageHooks()
  return (
    <Container maxW="container.md" py="5%">
      <ArticleList articles={articles} token={token} />
      <Box ref={loader} h="1px" mt="19px" />
      {isVisible && <Loading />}
    </Container>
  )
}
