import useTodoStore from "../store/todoStore";
import { Box, Tabs,Tab, TabList, TabPanels, TabPanel, VStack } from "@chakra-ui/react";
import TodoItem from "./TodoItem";


export default function ListTodos() {
    // TODO: answer here

    return (
      <Box width="xl" padding='1rem'>
        <Tabs width="full">
          <TabList>
            <Tab>All</Tab>
            <Tab>Done</Tab>
            <Tab>Undone</Tab>
          </TabList>
          <TabPanels>
            <TabPanel>
              <VStack>
                {todos.map((todo) => (
                  <TodoItem key={todo.id} {...todo} />
                ))}
              </VStack>
            </TabPanel>
            <TabPanel>
                <VStack>
              {doneTodos.map((todo) => (
                <TodoItem key={todo.id} {...todo} />
              ))}
                </VStack>
            </TabPanel>
            <TabPanel>
                <VStack>
              {undoneTodos.map((todo) => (
                <TodoItem key={todo.id} {...todo} />
              ))}
                </VStack>
            </TabPanel>
          </TabPanels>
        </Tabs>
      </Box>
    )
}