import './App.css';
import { HStack, Text } from '@chakra-ui/react';
import CartContainer from './components/CartContainer';
import ItemsContainer from './components/ItemsContainer';
import NewItem from './components/NewItemContainer';

function App() {
  return (
    <div className="App" aria-label='app'>
      <HStack
        justify="space-between"
        padding="0.5rem"
        width="full"
        backgroundColor="green.400"
      >
        <Text textColor="white" fontSize="1.5rem" fontWeight="semibold">
          Oversimplified E-Commerce App
        </Text>
        <HStack>
          <NewItem />
          <CartContainer />
        </HStack>
      </HStack>
      <ItemsContainer />
    </div>
  )
}

export default App;
