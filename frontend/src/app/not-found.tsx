import { Container, Image, Text } from '@chakra-ui/react'

export default async function NotFound() {
  return (
    <Container mt="100px">
      <Text fontSize="28px" fontWeight={600} textAlign="center">
        お探しのページが見つかりませんでした
      </Text>
      <Image src="/icons/spheart_mono.svg" alt="#" w="30%" h="30%" mx="auto" />
    </Container>
  )
}
