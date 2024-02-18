'use client'

import { IsSearchAtom } from '@/states/IsSearchAtom'
import { Search2Icon } from '@chakra-ui/icons'
import { useRecoilState } from 'recoil'

export const SearchIconComponent = () => {
  const [isSearch, setIsSearch] = useRecoilState(IsSearchAtom)
  return (
    <Search2Icon
      h="60%"
      w="5%"
      my="auto"
      color="black.primary"
      onClick={() => setIsSearch(!isSearch)}
      _hover={{ cursor: 'pointer', opacity: '0.5' }}
    />
  )
}
