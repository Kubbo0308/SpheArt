import { Box, Flex, Skeleton } from '@chakra-ui/react'

export const ArticleSkeleton = () => {
  return (
    <Box
      borderRadius="8px"
      overflow="hidden"
      boxShadow="lg"
      bg="white.primary"
      w="320px"
      border="1px 0px 0px 1px"
      borderColor="gray.primary"
      opacity={0.3}
    >
      <Skeleton height="180px" />
      <Box p="10px" h="110px">
        <Skeleton height="15px" mt="10px" />
        <Flex mt="50px" justifyContent="space-between">
          <Skeleton height="15px" width="150px" />
          <Skeleton height="15px" width="50px" />
        </Flex>
      </Box>
    </Box>
  )
}
