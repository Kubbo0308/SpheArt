import { Button, Link } from '@chakra-ui/react'

interface LinkButtonProps {
  title: string
  url: string
  bgColor?: string
  color?: string
  onClick?: () => void
}

export const LinkButton = (props: LinkButtonProps) => {
  const { title, url, bgColor = 'black.primary', color = 'white.primary', onClick } = props
  return (
    <Button bg={bgColor} color={color} onClick={onClick}>
      <Link href={url}>{title}</Link>
    </Button>
  )
}
