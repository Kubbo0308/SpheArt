import { Flex, Image } from '@chakra-ui/react'
import { SearchInput } from '../atoms/SearchInput'
import { SearchIconComponent } from '../atoms/SearchIconComponent'

export const Header = () => {
  return (
    <>
      <Flex bg="yellow.primary" w="100%" h="7vh" alignItems="center">
        <Flex justifyContent="space-between" h="80%" px="3%" w="100%">
          <Image
            src="/icons/techpulse_transparent_big_name.png"
            alt="#"
            w="auto"
            h="100%"
            objectFit="cover"
            _hover={{ cursor: 'pointer' }}
          />
          <SearchIconComponent />
        </Flex>
      </Flex>
      <SearchInput />
    </>
  )
}
