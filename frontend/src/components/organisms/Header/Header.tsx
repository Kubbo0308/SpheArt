import { Flex, Image } from '@chakra-ui/react'
import { SearchInput } from '../../atoms/SearchInput'
import { SearchIconComponent } from '../../atoms/SearchIconComponent'
import Link from 'next/link'
import { CONST } from '@/const'
import { AuthButton } from '@/components/molecules/AuthButton/AuthButton'

export const Header = () => {
  return (
    <>
      <Flex bg="yellow.primary" w="100%" h="7vh" alignItems="center">
        <Flex justifyContent="space-between" h="80%" px="3%" w="100%">
          <Link href={CONST.TOP}>
            <Image
              src="/icons/techpulse_transparent_big_name.png"
              alt="#"
              w="100%"
              h="100%"
              objectFit="cover"
              _hover={{ cursor: 'pointer', opacity: '0.5' }}
            />
          </Link>
          <Flex w="100%" h="100%" justifyContent="flex-end" gap="3%">
            <SearchIconComponent />
            <a href={`${CONST.BOOKMARK}`}>bookmark</a>
            <AuthButton />
          </Flex>
        </Flex>
      </Flex>
      <SearchInput />
    </>
  )
}
