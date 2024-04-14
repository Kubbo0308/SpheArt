import { Button, Link } from '@chakra-ui/react'

interface LinkButtonProps {
  title: string
  url?: string
  bgColor?: string
  color?: string
  onClick?: () => void
}

export const LinkButton = (props: LinkButtonProps) => {
  const { title, url, bgColor = 'blue.accent', color = 'white.primary', onClick } = props
  return (
    <Button bg={bgColor} color={color} onClick={onClick} borderRadius="20px" h="35px" w="70px">
      <Link href={url} fontSize="15px" fontWeight={600}>
        {title}
      </Link>
    </Button>
  )
}
