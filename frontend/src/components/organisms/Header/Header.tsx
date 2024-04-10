import { Flex, Image, Text } from '@chakra-ui/react'
import { SearchInput } from '../../atoms/SearchInput'
import { SearchIconComponent } from '../../atoms/SearchIconComponent'
import Link from 'next/link'
import { CONST } from '@/const'
import { AuthButton } from '@/components/molecules/AuthButton/AuthButton'

export const Header = () => {
  return (
    <>
      <Flex bg="white.primary" w="100%" h="7vh" alignItems="center">
        <Flex justifyContent="space-between" h="80%" px="3%" w="100%">
          <Link href={CONST.TOP}>
            <Flex gap="5px" h="100%" alignItems="center">
              <Image
                src="/icons/SpheArt.svg"
                alt="#"
                w="100%"
                h="100%"
                _hover={{ cursor: 'pointer', opacity: '0.5' }}
              />
              <Text fontSize="28px" fontWeight={700}>
                SpheArt
              </Text>
            </Flex>
          </Link>
          <Flex w="100%" h="100%" justifyContent="flex-end" gap="3%" alignItems="center">
            <SearchIconComponent />
            {/* <a href={`${CONST.BOOKMARK}`}>bookmark</a> */}
            <AuthButton />
          </Flex>
        </Flex>
      </Flex>
      <SearchInput />
    </>
  )
}
