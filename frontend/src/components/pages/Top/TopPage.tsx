'use client'

import { ArticleList } from '@/components/organisms/ArticleList'
import { Button, Container, Flex } from '@chakra-ui/react'
import { useTopPageHooks } from './TopPage.hooks'

interface TopPageProps {
  token: string | undefined
}

export const TopPage = (props: TopPageProps) => {
  const { token } = props
  const { currentPage, articles, goNextPage, backPreviousPage } = useTopPageHooks()
  return (
    <Container maxW="container.md" py="5%">
      <ArticleList articles={articles} token={token} />
      <Flex gap="20px">
        <Button onClick={backPreviousPage}>前へ</Button>
        <Button onClick={goNextPage}>次へ</Button>
      </Flex>
    </Container>
  )
}
