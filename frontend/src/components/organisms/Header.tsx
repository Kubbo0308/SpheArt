import { Flex, Image } from '@chakra-ui/react'
import { SearchInput } from '../atoms/SearchInput'
import { SearchIconComponent } from '../atoms/SearchIconComponent'
import Link from 'next/link'
import { CONST } from '@/const'
import { LinkButton } from '../atoms/LinkButton'
import { getCookies } from 'next-client-cookies/server'
import { RemoveCookieButton } from '../atoms/RemoveCookieButton'

export const Header = () => {
  const cookies = getCookies()

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
            {cookies.get('token') !== undefined && <a href={`${CONST.BOOKMARK}`}>bookmark</a>}
            {cookies.get('token') !== undefined ? (
              <RemoveCookieButton title="ログアウト" url={`${CONST.TOP}`} />
            ) : (
              <LinkButton title="ログイン" url={`${CONST.AUTH}${CONST.SIGN_IN}`} />
            )}
          </Flex>
        </Flex>
      </Flex>
      <SearchInput />
    </>
  )
}
