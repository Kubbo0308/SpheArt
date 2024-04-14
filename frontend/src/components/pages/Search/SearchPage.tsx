'use client'

import { ArticleList } from '@/components/organisms/ArticleList'
import { Box, Container, Text } from '@chakra-ui/react'
import { useSearchPageHooks } from './SearchPage.hooks'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { Loading } from '@/components/molecules/Loading'

interface SearchPageProps {
  token: RequestCookie | undefined
}

export const SearchPage = (props: SearchPageProps) => {
  const { token } = props
  const { articles, loader, isVisible } = useSearchPageHooks()
  return (
    <>
      <Text fontSize="32px" fontWeight={600} lineHeight={1.8} mt="30px" textAlign="center">
        検索結果
      </Text>
      <Container maxW="container.md" py="5%">
        <ArticleList articles={articles} token={token} />
        <Box ref={loader} h="1px" mt="19px" />
        {isVisible && <Loading />}
      </Container>
    </>
  )
}
