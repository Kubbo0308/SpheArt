import { Button, Link } from '@chakra-ui/react'

interface PrimaryButtonProps {
  title: string
  url: string
  bgColor?: string
  color?: string
}

export const PrimaryButton = (props: PrimaryButtonProps) => {
  const { title, url, bgColor = 'black.primary', color = 'white.primary' } = props
  return (
    <Button bg={bgColor} color={color}>
      <Link href={url}>{title}</Link>
    </Button>
  )
}
