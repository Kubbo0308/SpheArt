'use client'

import { ArticleList } from '@/components/organisms/ArticleList'
import { Button, Container, Flex } from '@chakra-ui/react'
import { useTopPageHooks } from './TopPage.hooks'
import { RequestCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { Suspense } from 'react'
import { Loading } from '@/components/molecules/Loading'

interface TopPageProps {
  token: RequestCookie | undefined
}

export const TopPage = (props: TopPageProps) => {
  const { token } = props
  const { articles, goNextPage, backPreviousPage } = useTopPageHooks()
  return (
    <Container maxW="container.md" py="5%">
      <Suspense fallback={<Loading />}>
        <ArticleList articles={articles} token={token} />
      </Suspense>
      <Flex gap="20px">
        <Button onClick={backPreviousPage}>前へ</Button>
        <Button onClick={goNextPage}>次へ</Button>
      </Flex>
    </Container>
  )
}
