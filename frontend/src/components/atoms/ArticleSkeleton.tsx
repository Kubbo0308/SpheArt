import { Box, Flex } from '@chakra-ui/react'

export const ArticleSkeleton = () => {
  return (
    <Box
      borderRadius="8px"
      overflow="hidden"
      boxShadow="sm"
      bg="white.primary"
      w="320px"
      border="2px"
      borderColor="gray.primary"
    >
      <Box bg="gray.pale" h="180px" borderBottom="2px" borderColor="gray.primary" />
      <Box p="10px" h="110px">
        <Box h="15px" bg="gray.primary" mt="10px" />
        <Flex mt="50px" justifyContent="space-between">
          <Box h="15px" bg="gray.primary" w="150px" />
          <Box h="15px" bg="gray.primary" w="50px" />
        </Flex>
      </Box>
    </Box>
  )
}
