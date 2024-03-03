import { Button, Link } from '@chakra-ui/react'

interface LinkButtonProps {
  title: string
  url: string
  bgColor?: string
  color?: string
}

export const LinkButton = (props: LinkButtonProps) => {
  const { title, url, bgColor = 'black.primary', color = 'white.primary' } = props
  return (
    <Button bg={bgColor} color={color}>
      <Link href={url}>{title}</Link>
    </Button>
  )
}
