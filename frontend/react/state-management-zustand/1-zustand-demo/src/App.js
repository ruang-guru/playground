import './App.css';
import { Flex, HStack, Text } from '@chakra-ui/react';
import useStore from './store/Store';
import ProfileDetail from './components/ProfileDetail';
import ProfileEdit from './components/ProfileEdit';

function App() {
  const { name } = useStore()

  return (
    <div className="App">
      <Flex backgroundColor='whatsapp.400' width='full'>
        <HStack justify='space-between' width='full' paddingInline='1rem' paddingBlock='0.5rem'>
          <Text>
            Root Component
          </Text>
          <Text fontWeight='semibold'>
            {name}
          </Text>
        </HStack>
      </Flex>
      <HStack width='full' justify='space-around' padding='2rem'>
        <ProfileDetail />
        <ProfileEdit />
      </HStack>
    </div>
  );
}

export default App;
