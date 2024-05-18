import { Image, Text } from '@chakra-ui/react'

export const NoArticle = () => {
  return (
    <>
      <Text fontSize="28px" fontWeight={600} textAlign="center">
        記事がありません
      </Text>
      <Image src="/icons/spheart_mono.svg" alt="#" w="30%" h="30%" mx="auto" />
    </>
  )
}
