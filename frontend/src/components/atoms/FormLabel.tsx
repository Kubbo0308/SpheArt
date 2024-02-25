import { Center, Flex } from '@chakra-ui/react'

interface FormLabelProps {
  label: string
  required: boolean
}

export const FormLabel = (props: FormLabelProps) => {
  const { label, required } = props
  return (
    <Flex alignItems="center">
      {label}
      <Center display="inline-block" bg={required ? 'black.primary' : 'yellow.primary'} borderRadius="18px">
        {required ? '必須' : '任意'}
      </Center>
    </Flex>
  )
}
