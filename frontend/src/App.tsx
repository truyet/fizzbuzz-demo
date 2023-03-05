import React, { useEffect, useState } from 'react';
import './App.css';
import { Button, Flex, MantineProvider, Text } from '@mantine/core';
import { fizzbuzz } from './fizzbuzz';

function App() {
  const [count, setCount] = useState(0)
  const [message, setMessage] = useState("")

  useEffect(() => {
    fizzbuzz({ count }).then(resp => {
      setMessage(resp.message)
    })
  }, [count])

  return (
    <MantineProvider
      withGlobalStyles
      withNormalizeCSS
      theme={{
        breakpoints: {
          xs: '30em',
          sm: '48em',
          md: '64em',
          lg: '74em',
          xl: '90em',
        },
      }}
    >

    <Flex
      h="100vh"
      justify="center"
      align="center"
      direction="column"
    >
      <Text fz="lg">Your count</Text>
      <Text fz="lg">{ count }</Text>
      <Button mt={24} onClick={() => {setCount(count + 1)}}>
      Push me!
    </Button>
    <Text mt={48} fz="xl" fw={700} >{message}</Text>
    </Flex>
    </MantineProvider>
  );
}

export default App;
