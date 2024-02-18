'use client'

import { IsSearchAtom } from '@/states/IsSearchAtom'
import { Image } from '@chakra-ui/react'
import { useRecoilState } from 'recoil'

export const SearchIconComponent = () => {
  const [isSearch, setIsSearch] = useRecoilState(IsSearchAtom)
  return (
    <Image
      src="/icons/magnifier.svg"
      alt=""
      h="60%"
      my="auto"
      onClick={() => setIsSearch(!isSearch)}
      _hover={{ cursor: 'pointer' }}
    />
  )
}
