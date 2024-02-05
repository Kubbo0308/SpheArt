import { Flex, Image } from '@chakra-ui/react'

export const Header = () => {
  return (
    <Flex bg="yellow.primary" width="100%" height="10vh" alignItems="center">
      <Flex justifyContent="space-between" height="80%" px="3%">
        <Image src="/techpulse_transparent_big_name.png" alt="#" width="auto" height="100%" objectFit="cover" />
      </Flex>
    </Flex>
  )
}
