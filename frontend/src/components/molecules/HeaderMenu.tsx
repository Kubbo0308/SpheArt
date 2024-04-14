import { CONST } from '@/const'
import { Button, Flex, Image, Menu, MenuButton, MenuDivider, MenuItem, MenuList } from '@chakra-ui/react'
import Link from 'next/link'

interface HeaderMenuProps {
  onClick: () => void
}

export const HeaderMenu = (props: HeaderMenuProps) => {
  const { onClick } = props

  return (
    <Menu>
      <MenuButton as={Button} bg="blue.accent" color="white.primary" borderRadius="20px" h="35px" w="70px" px="8px">
        Menu
      </MenuButton>
      <MenuList px="10px" w="200px">
        <MenuItem as={Link} href={`${CONST.BOOKMARK}`} fontSize="16px" fontWeight={600} lineHeight={1.8}>
          <Flex gap="3px">
            <Image src="/header/bookmark_256.svg" alt="" w="16px" />
            ブックマーク
          </Flex>
        </MenuItem>
        <MenuDivider />
        <MenuItem as={Button} onClick={onClick} fontSize="16px" fontWeight={600} lineHeight={1.8} justifyContent="left">
          <Flex gap="3px">
            <Image src="/header/sign_out.svg" alt="" w="16px" />
            サインアウト
          </Flex>
        </MenuItem>
      </MenuList>
    </Menu>
  )
}
