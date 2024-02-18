import { IsSearchAtom } from '@/states/IsSearchAtom'
import { SearchIcon } from '@chakra-ui/icons'
import { useRecoilState } from 'recoil'

export const SearchIconComponent = () => {
  const [isSearch, setIsSearch] = useRecoilState(IsSearchAtom)
  return <SearchIcon onClick={() => setIsSearch(!isSearch)} />
}
