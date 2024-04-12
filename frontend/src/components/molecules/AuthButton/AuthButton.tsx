'use client'

import { useAuthButton } from './AuthButton.hooks'
import { LinkButton } from '@/components/atoms/LinkButton'
import { CONST } from '@/const'
import { Button, Flex, Image, Menu, MenuButton, MenuDivider, MenuItem, MenuList } from '@chakra-ui/react'
import Link from 'next/link'

export const AuthButton = () => {
  const { onSignOut, token } = useAuthButton()

  return (
    <>
      {token === undefined ? (
        <LinkButton title="Sign In" url={`${CONST.AUTH}${CONST.SIGN_IN}`} />
      ) : (
        <Menu>
          <MenuButton as={Button}>メニュー</MenuButton>
          <MenuList px="10px" w="200px">
            <MenuItem as={Link} href={`${CONST.BOOKMARK}`} fontSize="16px" fontWeight={600} lineHeight={1.8}>
              <Flex gap="3px">
                <Image src="/header/bookmark_256.svg" alt="" w="16px" />
                ブックマーク
              </Flex>
            </MenuItem>
            <MenuDivider />
            <MenuItem
              as={Button}
              onClick={onSignOut}
              fontSize="16px"
              fontWeight={600}
              lineHeight={1.8}
              justifyContent="left"
            >
              <Flex gap="3px">
                <Image src="/header/sign_out.svg" alt="" w="16px" />
                サインアウト
              </Flex>
            </MenuItem>
          </MenuList>
        </Menu>
      )}
    </>
  )
}
