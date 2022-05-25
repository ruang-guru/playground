import './App.css';
import { HStack, Text, VStack } from '@chakra-ui/react';
import ListTodos from './components/ListTodos';
import AddTodo from './components/AddTodo';
import SummaryTodos from './components/SummaryTodos';

function App() {
  return (
    <div className="App" aria-label='app'>
     <HStack justify='center' padding='0.5rem' width='full' backgroundColor='yellow.400'>
        <Text fontSize='1.5rem' fontWeight='semibold'>
          Simple Todos App
        </Text>
     </HStack>
     <HStack align='start' width='full' justify='space-around' padding='1rem'>
       <ListTodos />
       <VStack>
         <SummaryTodos />
         <AddTodo />
       </VStack>
     </HStack>
    </div>
  );
}

export default App;
