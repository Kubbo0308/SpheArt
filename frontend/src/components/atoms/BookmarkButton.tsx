import { Image } from '@chakra-ui/react'

interface BookmarkButtonProps {
  onClick: () => void
  isBookmark: boolean
}

export const BookmarkButton = (props: BookmarkButtonProps) => {
  const { onClick, isBookmark } = props

  return (
    <button onClick={onClick}>
      {isBookmark ? (
        <Image src="/bookmark_fill_256.svg" alt="" w="16px" h="16px" />
      ) : (
        <Image src="/bookmark_256.svg" alt="" w="16px" h="16px" />
      )}
    </button>
  )
}
