'use client'

import { ArticleList } from '@/components/organisms/ArticleList'
import { Box, Container } from '@chakra-ui/react'
import { useTopPageHooks } from './TopPage.hooks'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { Loading } from '@/components/molecules/Loading'

interface TopPageProps {
  token: RequestCookie | undefined
}

export const TopPage = (props: TopPageProps) => {
  const { token } = props
  const { articles, loader } = useTopPageHooks()
  return (
    <Container maxW="container.md" py="5%">
      <ArticleList articles={articles} token={token} />
      <Box ref={loader} h="1px" mt="19px" />
      <Loading />
    </Container>
  )
}
