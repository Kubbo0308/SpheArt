import { Flex } from '@chakra-ui/react'
import { ArticleSkeleton } from '../atoms/ArticleSkeleton'

export const Loading = () => {
  return (
    <Flex flexWrap="wrap" gap="20px" justifyContent="center">
      <ArticleSkeleton />
      <ArticleSkeleton />
    </Flex>
  )
}
