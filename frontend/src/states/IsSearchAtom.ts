import { atom } from 'recoil'

export const IsSearchAtom = atom<boolean>({
  key: 'IsSearch',
  default: false
})
