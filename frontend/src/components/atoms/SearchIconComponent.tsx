'use client'

import { IsSearchAtom } from '@/states/IsSearchAtom'
import { Search2Icon } from '@chakra-ui/icons'
import { useRecoilState } from 'recoil'

export const SearchIconComponent = () => {
  const [isSearch, setIsSearch] = useRecoilState(IsSearchAtom)
  return (
    <Search2Icon
      h="25px"
      w="25px"
      my="auto"
      color="gray.accent"
      onClick={() => setIsSearch(!isSearch)}
      _hover={{ cursor: 'pointer', opacity: '0.5' }}
    />
  )
}
