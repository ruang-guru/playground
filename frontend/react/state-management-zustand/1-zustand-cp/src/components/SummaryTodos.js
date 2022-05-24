import useTodoStore from "../store/todoStore";
import { Box, Text } from "@chakra-ui/react";

export default function SummaryTodos() {
    // TODO: answer here

    return (
        <Box width="md" rounded='xl' border='1px'>
            <Text aria-label='todo-count-total'>
                Total : {count}
            </Text>
            <Text aria-label='todo-count-done'>
                Done : {doneCount}
            </Text>
        </Box>
    )
}